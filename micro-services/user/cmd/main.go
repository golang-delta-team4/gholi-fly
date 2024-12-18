package main

import (
	"log"
	"user-service/api/handlers/http"
	"user-service/app"
	"user-service/config"
)

func main() {
	config, err := config.ReadConfig(".")
	if err != nil {
		log.Println(err)
		return
	}
	app, err := app.NewApp(config)
	if err != nil {
		log.Println(err)
		return
	}
	err = http.Run(app, config)
	if err != nil {
		log.Fatal(err)
	}
	
}
