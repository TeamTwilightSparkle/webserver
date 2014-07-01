package rest

import (
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"net/url"

	"github.com/TeamTwilightSparkle/webserver/model"
	"github.com/TeamTwilightSparkle/webserver/controller"
)

type RestHelper func([]string, http.ResponseWriter, *http.Request)

func RestProfile(rest []string, w http.ResponseWriter, r *http.Request) {
	if len(rest) < 4 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	queries, _ := url.ParseQuery(r.URL.RawQuery)
	profile := new(model.Profile)

	if result, err := profile.Get(queries, rest[controller.FIELD_INDEX], rest[controller.VALUE_INDEX]); err == nil {
		output(profile.Format(result), queries, w)
		return
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}



func RestContent(rest []string, w http.ResponseWriter, r *http.Request) {
	if len(rest) < 4 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	queries, _ := url.ParseQuery(r.URL.RawQuery)
	content := new(model.Content)

	if result, err := content.Get(queries, rest[controller.FIELD_INDEX], rest[controller.VALUE_INDEX]); err == nil {
		output(content.Format(result), queries, w)
		return
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func output(iface interface {}, queries url.Values, w http.ResponseWriter) {
	json, _ := json.MarshalIndent(iface, "", "\t")

	var output string = string(json)
	if encode := queries.Get("encode"); encode == "true" {
		output = base64.StdEncoding.EncodeToString(json)
	}
	fmt.Fprintf(w, "%v\n", output)
}
