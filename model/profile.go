package model

import (
	"fmt"
	"time"
	"errors"

	"github.com/TeamTwilightSparkle/webserver/database"
)

type profile_functions func (string, string, string)([]Profile, error)

var profile_table map[string] profile_functions

type Profile struct {
	Id          int			`json:"id"`
	Username    string		`json:"username"`
	Level       int			`json:"level"`
	Last_seen   time.Time	`json:"last_seen"`
	Date_joined time.Time	`json:"date_joined"`
}

func init() {
	profile_table = make(map[string] profile_functions)
	profile_table["id"] = getFromInt
	profile_table["username"] = getFromString
	profile_table["level"] = getFromInt
}

func (_ Profile) Get(omni, field, find string) (profile []Profile, err error) {
	if call := profile_table[field]; call != nil {
		return call(omni, field, find)
	}
	return nil, errors.New("Bad Request")
}

func getFromInt(_, field, value string) (profile []Profile, err error) {
	if err = database.Conn.Select(&profile, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = %s", field, value)); err != nil {
		return nil, err
	}
	return
}

func getFromString(omni, field, value string) (profile []Profile, err error) {
	var sql_query string
	if omni == "true" {
		sql_query = fmt.Sprintf("SELECT * FROM PROFILE WHERE %s LIKE '%s%%'", field, value)
	} else {
		sql_query = fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = '%s'", field, value)
	}

	if err = database.Conn.Select(&profile, sql_query); err != nil {
		return nil, err
	}
	return
}

func (_ Profile) Validate(profile []Profile) interface {} {
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
