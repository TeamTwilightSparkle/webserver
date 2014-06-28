package model

import "time"

type Comment struct {
	Id            int		`json:"id"`
	Username	  string	`json:"username"`
	Content_id    int		`json:"content_id"`
	Post_number   int		`json:"post_number"`
	Comment       string	`json:"comment"`
	Last_modified time.Time	`json:"last_modified"`
	Date_created  time.Time	`json:"date_created"`
}
