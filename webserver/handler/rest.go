package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"net/url"

	"github.com/souleiman/seesaw/webserver/model/profile"
//	"github.com/souleiman/seesaw/webserver/model"
)

type RestHelper func([]string, http.ResponseWriter, *http.Request)

func RestProfile(rest []string, w http.ResponseWriter, r *http.Request) {
	if len(rest) < 4 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var profiles []profile.Profile
	var err error

	if profiles, err = profile.Get(rest[FIELD_INDEX], rest[VALUE_INDEX]); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var iface interface {}
	if len := len(profiles); len == 0 {
		iface = nil
	} else if len == 1 {
		iface = profiles[0]
	} else {
		iface = struct {Result []profile.Profile `json:"result"`} {profiles}
	}

	json, _ := json.MarshalIndent(iface, "", "\t")
	qMap, _ := url.ParseQuery(r.URL.RawQuery)

	var output string = string(json)
	if encode := qMap["encode"]; encode != nil && encode[0] == "true" {
		output = base64.StdEncoding.EncodeToString(json)
	}
	fmt.Fprintf(w, "%v\n", output)
}

func RestContent(rest []string, w http.ResponseWriter, r *http.Request) {
	if len(rest) < 4 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

//	var contents []model.Content
//	var sql_query string

	switch rest[FIELD_INDEX] {
	case "id":
//		sql_query = fmt.Sprintf("SELECT * FROM CONTENT WHERE %s = %s", rest[FIELD_INDEX], rest[VALUE_INDEX])
	case "author":
//		sql_query = fmt.Sprintf("SELECT CONTENT.* FROM CONTENT, PROFILE WHERE %s = '%s' AND content.author = profile.username", rest[FIELD_INDEX], rest[VALUE_INDEX])
	case "tag":
//		sql_query = fmt.Sprintf("SELECT CONTENT.* FROM CONTENT, TAG WHERE tag.content_id = content.id AND %s = '#%s';", rest[FIELD_INDEX], rest[VALUE_INDEX])
	default:
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

/*	if _, err := dbmap.Select(&contents, sql_query); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusConflict)
		return
	}

	tag_query := "SELECT DISTINCT TAG.* FROM TAG, CONTENT WHERE tag.content_id = $1"
	comment_query := "SELECT DISTINCT COMMENT.* FROM COMMENT, CONTENT WHERE comment.content_id = $1"
	charge_query := "SELECT DISTINCT CHARGE.* FROM CONTENT, CHARGE WHERE charge.charged = TRUE AND charge.content_id = $1"

	for i, c := range contents {
		var comments []model.Comment
		var tags []model.Tag
		var charges []model.Charge
		if _, err := dbmap.Select(&comments, comment_query, c.Id); err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusConflict)
			return
		}

		if _, err := dbmap.Select(&tags, tag_query, c.Id); err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusConflict)
			return
		}

		if _, err := dbmap.Select(&charges, charge_query, c.Id); err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusConflict)
			return
		}

		contents[i].Comments = comments
		contents[i].Tags = tags
		contents[i].Charges = len(charges)
	}

	var iface interface {}
	if len := len(contents); len == 0 {
		iface = nil
	} else if len == 1 {
		iface = contents[0]
	} else {
		iface = struct {Result []model.Content `json:"result"`} {contents}
	}

	json, _ := json.MarshalIndent(iface, "", "\t")
	qMap, _ := url.ParseQuery(r.URL.RawQuery)

	var output string = string(json)
	if encode := qMap["encode"]; encode != nil && encode[0] == "true" {
		output = base64.StdEncoding.EncodeToString(json)
	}
	fmt.Fprintf(w, "%v\n", output)*/
}
