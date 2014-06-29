package main

import (
	"net/http"

	"github.com/docopt/docopt-go"
	"github.com/TeamTwilightSparkle/webserver/controller/database"
	"github.com/TeamTwilightSparkle/webserver/controller/handler"
)

func main() {

	usage :=
`WebServer:
  Get you one.

Usage:
  webserver --version | -v
  webserver --help | -h
  webserver [options] (--user=<user>) (--database=<db>)

Options:
  -h --help  Show this screen
  --version  Show version
  --port=<port>  Port value for server to listen on [default: 80].
  --root=<root>  Root directory for server to serve files [default: /].
  --api=<dev>  API path for developers [default: /api/].
  --user=<user>  Database username for access.
  --database=<db>  Database name to access.`

	args, _ := docopt.Parse(usage, nil, true, "Webserver 1.0", true)
	database.BootDatabase(args["--user"].(string), args["--database"].(string))
	defer database.CloseDB()

	http.HandleFunc(args["--root"].(string), handler.Root())
	http.HandleFunc(args["--api"].(string), handler.Api())

	http.ListenAndServe(":" + args["--port"].(string), nil)
}
