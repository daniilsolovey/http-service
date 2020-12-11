package main

import (
	"net/http"

	"github.com/daniilsolovey/http-service/internal/config"
	"github.com/daniilsolovey/http-service/internal/handler"
	"github.com/docopt/docopt-go"
	"github.com/reconquest/karma-go"
	"github.com/reconquest/pkg/log"
)

var version = "[manual build]"

var usage = `http-service

Create users, receive users, update users

Usage:
	http-service [options]

Options:
  -c --config <path>  Read specified config file. [default: config.toml]
  --debug             Enable debug messages.
  -v --version        Print version.
  -h --help           Show this help.
`

func main() {
	args, err := docopt.ParseArgs(
		usage,
		nil,
		"http-service version: "+version,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof(
		karma.Describe("version", version),
		"http-service started",
	)

	if args["--debug"].(bool) {
		log.SetLevel(log.LevelDebug)
	}

	log.Infof(nil, "loading configuration file: %q", args["--config"].(string))

	config, err := config.Load(args["--config"].(string))
	if err != nil {
		log.Fatal(err)
	}

	users := handler.CreateUsers()
	handler := handler.NewHandler(config, users)
	router := handler.CreateRouter()
	err = http.ListenAndServe(config.HTTPPort, router)
	if err != nil {
		log.Fatalf(err, "unable to listen and serve on port: %s", config.HTTPPort)
	}
}
