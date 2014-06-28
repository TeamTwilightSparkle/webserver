package profile

import (
	"fmt"
	"time"

	"github.com/souleiman/seesaw/webserver/database"
	"errors"
)

type profile_functions func (string, string)([]Profile, error)

var table map[string] profile_functions

type Profile struct {
	Id          int			`json:"id"`
	Username    string		`json:"username"`
	Level       int			`json:"level"`
	charges		int			`json:"charges",db:"-"`
	Last_seen   time.Time	`json:"last_seen"`
	Date_joined time.Time	`json:"date_joined"`
}

func init() {
	table = make(map[string] profile_functions)
	table["id"] = getFromInt
	table["username"] = getFromString
	table["level"] = getFromInt
	table["charge"] = getCharges
}

func Get(field, find string) (profile []Profile, err error) {
	if call := table[field]; call != nil {
		return call(field, find)
	}
	return nil, errors.New("Bad Request")
}

func getFromInt(field, value string) (profile []Profile, err error) {
	if _, err = database.DBMap.Select(&profile, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = %s", field, value)); err != nil {
		return nil, err
	}
	return
}

func getFromString(field, value string) (profile []Profile, err error) {
	if _, err = database.DBMap.Select(&profile, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = '%s'", field, value)); err != nil {
		return nil, err
	}
	return
}
