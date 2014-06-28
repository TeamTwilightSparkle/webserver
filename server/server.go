package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/TeamTwilightSparkle/webserver/database"
	"github.com/TeamTwilightSparkle/webserver/handler"
)

func main() {
	port := flag.Int("p", 8080, "Port to listen to")
	root := flag.String("r", "/", "Root of the fileserver")
	api := flag.String("R", "/api/", "RESTful API Services")
	dbuser := flag.String("U", "", "User for database")
	dbname := flag.String("D", "", "Database to use")
	flag.Parse()

	database.BootDatabase(*dbuser, *dbname)
	defer database.CloseDB()
	http.HandleFunc(*root, handler.Root())
	http.HandleFunc(*api, handler.Api())

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
