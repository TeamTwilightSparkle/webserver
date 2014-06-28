package handler

import (
//	"database/sql"
	"fmt"
	"net/http"
	"strings"

//	"github.com/coopernurse/gorp"
)

const (
	API_SIZE = 4
	API_INDEX = 0
	REST_INDEX = 1
	FIELD_INDEX = 2
	VALUE_INDEX = 3
)

type HandleHelper func(http.ResponseWriter, *http.Request)
var api_map map[string] RestHelper

func init() {
	api_map = make(map[string] RestHelper)
	api_map["profile"] = RestProfile
	api_map["content"] = RestContent
}

func Root() HandleHelper {
	maps := make(map[string]string)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, maps["Test"])
	}
}

func Api() HandleHelper {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path[len(path) - 1:] != "/" {
			path += "/"
		}

		split := strings.Split(path[1: len(path) - 1], "/")
		if len(split) < 2 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if rest := api_map[split[REST_INDEX]]; rest != nil {
			rest(split, w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}
