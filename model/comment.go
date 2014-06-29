package model

import "time"

type Comment struct {
	Comment_id    int		`json:"comment_id"`
	Username	  string	`json:"username"`
	Id    int				`json:"id"`
	Post_number   int		`json:"post_number"`
	Comment       string	`json:"comment"`
	Last_modified time.Time	`json:"last_modified"`
	Date_created  time.Time	`json:"date_created"`
}
