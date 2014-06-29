package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TeamTwilightSparkle/webserver/controller"
	"github.com/TeamTwilightSparkle/webserver/controller/rest"
)

type RestHelper func([]string, http.ResponseWriter, *http.Request)
type HandleHelper func(http.ResponseWriter, *http.Request)

func Root() HandleHelper {
	maps := make(map[string]string)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, maps["Test"])
	}
}

func Api() HandleHelper {
	api_map := make(map[string] RestHelper)
	api_map["profile"] = rest.RestProfile
	api_map["content"] = rest.RestContent

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

		if rest := api_map[split[controller.REST_INDEX]]; rest != nil {
			rest(split, w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}
