package model

import (
	"time"
	"fmt"

	"github.com/coopernurse/gorp"
)

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

func getId(id string, dbmap *gorp.DbMap) (profile []Content, err error) {
	if _, err = dbmap.Select(&profile, fmt.Sprintf("SELECT * FROM PROFILE WHERE id = %s", id)); err != nil {
		return nil, err
	}
	return
}
