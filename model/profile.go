package model

import (
	"fmt"
	"time"
	"errors"
	"net/url"

	"github.com/TeamTwilightSparkle/webserver/controller/database"
)

type profile_functions func(url.Values, string, string) ([]Profile, error)

var profile_table map[string] profile_functions

type Profile struct {
	Id             int			`json:"id"`
	Username       string		`json:"username"`
	Level          int			`json:"level"`
	Charges        []string		`json:"charges",db:"-"`
	Contents       []string		`json:"contents",db:"-"`
	Last_seen      time.Time	`json:"last_seen"`
	Date_joined    time.Time	`json:"date_joined"`
}

func init() {
	profile_table = make(map[string] profile_functions)
	profile_table["id"] = getProfileFromInt
	profile_table["username"] = getProfileFromString
	profile_table["level"] = getProfileFromInt
}

func (_ Profile) Get(queries url.Values, field, find string) (profile []Profile, err error) {
	if call := profile_table[field]; call != nil {
		return call(queries, field, find)
	}
	return nil, errors.New("Bad Request")
}

func getProfileFromInt(queries url.Values, field, value string) (profile []Profile, err error) {
	if err = database.Conn.Select(&profile, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = %s", field, value)); err != nil {
		return nil, err
	}

	for i, _ := range profile {
		profile[i].setCharges(queries)
		profile[i].setContent(queries)
	}
	return
}

func getProfileFromString(queries url.Values, field, value string) (profile []Profile, err error) {
	var sql_query string
	if queries.Get("omnisearch") == "true" {
		sql_query = fmt.Sprintf("SELECT * FROM PROFILE WHERE %s LIKE '%s%%'", field, value)
	} else {
		sql_query = fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = '%s'", field, value)
	}

	if err = database.Conn.Select(&profile, sql_query); err != nil {
		return nil, err
	}

	return
}

func (p *Profile) setCharges(queries url.Values) error {
	var content []Content
	var sql_query string = "SELECT CONTENT.* FROM CONTENT JOIN PROFILE_CHARGES USING (ID) WHERE username = '%s'";

	if limit := queries.Get("charge_limit"); limit != "" {
		sql_query += " LIMIT "+limit;
	}

	if err := database.Conn.Select(&content, fmt.Sprintf(sql_query, p.Username)); err != nil {
		return err
	}

	p.Charges = make([]string, len(content))
	for i, c := range content {
		p.Charges[i] = fmt.Sprintf("/api/content/id/%d/", c.Id)
	}
	return nil
}

func (p *Profile) setContent(queries url.Values) error {
	var content []Content
	var sql_query string = "SELECT CONTENT.* FROM CONTENT JOIN CONTENT_PROFILES USING (ID) WHERE username = '%s'";

	if limit := queries.Get("content_limit"); limit != "" {
		sql_query += " LIMIT "+limit;
	}

	if err := database.Conn.Select(&content, fmt.Sprintf(sql_query, p.Username)); err != nil {
		return err
	}

	p.Contents = make([]string, len(content))
	for i, c := range content {
		p.Contents[i] = fmt.Sprintf("/api/content/id/%d/", c.Id)
	}
	return nil
}

func (_ Profile) Format(profile []Profile) interface{} {
	var iface interface{}
	if len := len(profile); len == 0 {
		iface = nil
	} else if len == 1 {
		iface = profile[0]
	} else {
		iface = struct {Result []Profile `json:"result"`} {profile}
	}
	return iface
}
