package main

import (
	"log"
	"sync"
	"user-service/api/handlers/grpc"
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting gRPC server on port %d", config.Server.GRPCPort)
		if err := grpc.Run(app, config); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = http.Run(app, config)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()

}
