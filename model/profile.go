package model

import (
	"fmt"
	"time"
	"errors"
	"net/url"

	"github.com/TeamTwilightSparkle/webserver/database"
)

type profile_functions func (url.Values, string, string)([]Profile, error)

var profile_table map[string] profile_functions

type Profile struct {
	Id          int			`json:"id"`
	Username    string		`json:"username"`
	Level       int			`json:"level"`
	Charges		[]string	`json:"charges",db:"-"`
	Contents	[]string	`json:"contents",db:"-"`
	Last_seen   time.Time	`json:"last_seen"`
	Date_joined time.Time	`json:"date_joined"`
}

func init() {
	profile_table = make(map[string] profile_functions)
	profile_table["id"] = getFromInt
	profile_table["username"] = getFromString
	profile_table["level"] = getFromInt
}

func (_ Profile) Get(queries url.Values, field, find string) (profile []Profile, err error) {
	if call := profile_table[field]; call != nil {
		return call(queries, field, find)
	}
	return nil, errors.New("Bad Request")
}

func getFromInt(queries url.Values, field, value string) (profile []Profile, err error) {
	if err = database.Conn.Select(&profile, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = %s", field, value)); err != nil {
		return nil, err
	}

	if err = setCharges(queries, profile); err != nil {
		return nil, err
	}
	if err = setContent(queries, profile); err != nil {
		return nil, err
	}

	return
}

func getFromString(queries url.Values, field, value string) (profile []Profile, err error) {
	var sql_query string
	if queries.Get("omnisearch") == "true" {
		sql_query = fmt.Sprintf("SELECT * FROM PROFILE WHERE %s LIKE '%s%%'", field, value)
	} else {
		sql_query = fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = '%s'", field, value)
	}

	if err = database.Conn.Select(&profile, sql_query); err != nil {
		return nil, err
	}

	if err = setCharges(queries, profile); err != nil {
		return nil, err
	}
	if err = setContent(queries, profile); err != nil {
		return nil, err
	}

	return
}

func setCharges(queries url.Values, profiles []Profile) error {
	var content []Content
	init := func() string {
		var sql_query string = "SELECT CONTENT.* FROM CONTENT JOIN PROFILE_CHARGES USING (ID) WHERE username = '%s'";
		if limit := queries.Get("charge_limit"); limit != "" {
			sql_query += " LIMIT " + limit;
		}
		return sql_query
	}

	db_call := func(index int, sql_query string) error {
		if err := database.Conn.Select(&content, fmt.Sprintf(sql_query, profiles[index].Username)); err != nil {
			return err
		}
		return nil
	}

	setup := func(index int) {
		profiles[index].Charges = make([]string, len(content))
		for i, c := range content {
			profiles[index].Charges[i] = fmt.Sprintf("/api/content/id/%d/", c.Id)
		}
	}

	clean := func() {
		content = nil
	}

	return setDetail(len(profiles), init, db_call, setup, clean)
}

func setContent(queries url.Values, profiles []Profile) error {
	var content []Content
	init := func() string {
		var sql_query string = "SELECT CONTENT.* FROM CONTENT JOIN CONTENT_PROFILES USING (ID) WHERE username = '%s'";
		if limit := queries.Get("content_limit"); limit != "" {
			sql_query += " LIMIT " + limit;
		}
		return sql_query
	}

	db_call := func(index int, sql_query string) error {
		if err := database.Conn.Select(&content, fmt.Sprintf(sql_query, profiles[index].Username)); err != nil {
			return err
		}
		return nil
	}

	setup := func(index int) {
		profiles[index].Contents = make([]string, len(content))
		for i, c := range content {
			profiles[index].Contents[i] = fmt.Sprintf("/api/content/id/%d/", c.Id)
		}
	}

	clean := func() {
		content = nil
	}

	return setDetail(len(profiles), init, db_call, setup, clean)
}

func setDetail(length int, init func() string, db_call func(int, string) error, setup func(int), clean func()) error {
	sql_query := init()

	for i := 0; i < length; i++ {
		if err := db_call(i, sql_query); err != nil {
			return err
		}
		setup(i)
		clean()
	}
	return nil
}
func (_ Profile) Format(profile []Profile) interface {} {
	var iface interface {}
	if len := len(profile); len == 0 {
		iface = nil
	} else if len == 1 {
		iface = profile[0]
	} else {
		iface = struct {Result []Profile `json:"result"`} {profile}
	}
	return iface
}
