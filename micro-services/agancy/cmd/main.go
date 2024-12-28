package main

import (
	"flag"
	"gholi-fly-agancy/api/handlers/http"
	"gholi-fly-agancy/app"
	"gholi-fly-agancy/config"
	"log"
	"os"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(c)

	log.Fatal(http.Run(appContainer, c))
}
