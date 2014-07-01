package model

import (
	"time"
	"net/url"
	"fmt"
	"errors"

	"github.com/TeamTwilightSparkle/webserver/controller/database"
)

type content_functions func (url.Values, string, string)([]Content, error)

var content_table map[string] content_functions

type Content struct {
	Id            int		`json:"id"`
	Authors		  []string	`json:"authors",db:"-"`
	Title         string	`json:"title"`
	Summary       string	`json:"summary"`
	Description   string	`json:"description"`
	Comments	  []Comment	`json:"comments",db:"-"`
	Tags		  []string	`json:"tags",db:"-"`
	Charges		  int		`json:"charges"`
	Image         *string	`json:"image"`
	Last_modified time.Time	`json:"last_modfied"`
	Date_created  time.Time	`json:"date_created"`
}

func init() {
	content_table = make(map[string] content_functions)
	content_table["id"] = getContentFromInt
	content_table["author"] = getContentFromAuthor
	content_table["top"] = getContentFromRanking
	content_table["tag"] = getContentFromTag
}

func (_ Content) Get(queries url.Values, field, find string) ([]Content, error) {
	if call := content_table[field]; call != nil {
		return call(queries, field, find)
	}
	return nil, errors.New("Bad Request")
}

func getContentFromInt(queries url.Values, field, value string) (content []Content, err error) {
	if err = database.Conn.Select(&content, fmt.Sprintf("SELECT * FROM CONTENT WHERE %s = %s", field, value)); err != nil {
		return nil, err
	}

	finalizeContent(queries, content)
	return
}

func getContentFromAuthor(queries url.Values, _, value string) (content []Content, err error) {
	var sql_query string
	if queries.Get("omnisearch") == "true" {
		sql_query = fmt.Sprintf(
			"SELECT DISTINCT CONTENT.* FROM CONTENT JOIN CONTENT_PROFILES USING (id) WHERE username LIKE '%s%%'",
			value)
	} else {
		sql_query = fmt.Sprintf(
			"SELECT CONTENT.* FROM CONTENT JOIN CONTENT_PROFILES USING (id) WHERE username = '%s'",
			value)
	}
	if err = database.Conn.Select(&content, sql_query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	finalizeContent(queries, content)
	return
}

func getContentFromRanking(queries url.Values, _, value string) (content []Content, err error) {
	sql_query := "SELECT * FROM CONTENT ORDER BY CHARGES DESC LIMIT " + value
	if err = database.Conn.Select(&content, sql_query); err != nil {
		return nil, err
	}

	finalizeContent(queries, content)
	return
}

func getContentFromTag(queries url.Values, _, value string) (content []Content, err error) {
	sql_query := fmt.Sprintf("SELECT CONTENT.* FROM CONTENT JOIN CONTENT_TAG t USING (id) WHERE t.tag = '%s'", value)
	if err = database.Conn.Select(&content, sql_query); err != nil {
		return nil, err
	}

	finalizeContent(queries, content)
	return
}

func (c *Content) setAuthors(_ url.Values) error {
	var profiles []Profile
	var sql_query string = fmt.Sprintf(
		"SELECT CONTENT_PROFILES.username FROM CONTENT JOIN CONTENT_PROFILES USING (ID) WHERE CONTENT.ID = %d", c.Id)
	if err := database.Conn.Select(&profiles, sql_query); err != nil {
		return err
	}

	c.Authors = make([]string, len(profiles))
	for i, p := range profiles {
		c.Authors[i] = p.Username
	}
	return nil
}

func (c *Content) setComments(_ url.Values) error {
	var comment []Comment
	var sql_query string = fmt.Sprintf(
		"SELECT COMMENT.* FROM COMMENT WHERE COMMENT.ID = %d ORDER BY COMMENT.POST_NUMBER", c.Id)

	if err := database.Conn.Select(&comment, sql_query); err != nil {
		return err
	}

	c.Comments = comment
	return nil
}

func (c *Content) setTags(_ url.Values) error {
	var tags []Tag
	var sql_query string = fmt.Sprintf(
		"SELECT CONTENT_TAG.tag FROM CONTENT JOIN CONTENT_TAG USING (ID) WHERE CONTENT.ID = %d", c.Id)
	if err := database.Conn.Select(&tags, sql_query); err != nil {
		return err
	}

	c.Tags = make([]string, len(tags))
	for i, t := range tags {
		c.Tags[i] = t.Tag
	}
	return nil
}

func finalizeContent(queries url.Values, content []Content) {
	for i, _ := range content {
		content[i].setAuthors(queries)
		content[i].setComments(queries)
		content[i].setTags(queries)
	}
}

func (_ Content) Format(profile []Content) interface{} {
	var iface interface{}
	if len := len(profile); len == 0 {
		iface = nil
	} else if len == 1 {
		iface = profile[0]
	} else {
		iface = struct {Result []Content `json:"result"`} {profile}
	}
	return iface
}
