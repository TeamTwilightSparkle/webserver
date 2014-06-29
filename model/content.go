package model

import (
	"time"
	"fmt"
	"errors"

	"github.com/TeamTwilightSparkle/webserver/controller/database"
	"net/url"
)

type content_functions func (string, string)([]Content, error)

var content_table map[string] content_functions

type Content struct {
	Id            int		`json:"id"`
	Author		  string	`json:"author"`
	Title         string	`json:"title"`
	Summary       string	`json:"summary"`
	Description   string	`json:"description"`
	Comments	  []Comment	`json:"comments",db:"-"`
	Tags		  []Tag		`json:"tags",db:"-"`
	Charges		  int		`json:"charges",db:"-"`
	Image         *string	`json:"image"`
	Last_modified time.Time	`json:"last_modfied"`
	Date_created  time.Time	`json:"date_created"`
}

func init() {
	content_table = make(map[string] content_functions)
	content_table["id"] = getContentFromInt
	content_table["author"] = getContentFromString
}

func (_ Content) Get(queries url.Values, field, find string) ([]Content, error) {
	if call := content_table[field]; call != nil {
		return call(field, find)
	}
	return nil, errors.New("Bad Request")
}

func getContentFromInt(field, value string) (content []Content, err error) {
	if err = database.Conn.Select(&content, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = %s", field, value)); err != nil {
		return nil, err
	}
	return
}

func getContentFromString(field, value string) (content []Content, err error) {
	if err = database.Conn.Select(&content, fmt.Sprintf("SELECT * FROM PROFILE WHERE %s = '%s'", field, value)); err != nil {
		return nil, err
	}
	return
}
