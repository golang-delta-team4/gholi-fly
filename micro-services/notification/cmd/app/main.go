package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"notification-nats/config"
	"notification-nats/database"
	"notification-nats/notification"
	"notification-nats/pb"
	"notification-nats/shared"

	"google.golang.org/grpc"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	config := config.MustReadConfig(*configPath)

	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatal("error connecting to db")
	}

	if err := db.AutoMigrate(&shared.OutBoxMessage{}); err != nil {
		log.Fatal("migrate error - ", err)
	}

	grpcServer := grpc.NewServer()

	svc := &notification.Service{DB: db}
	pb.RegisterNotificationServiceServer(grpcServer, svc)

	address := fmt.Sprintf(":%d", config.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", config.Server.Port, err)
	}
	log.Printf("Starting gRPC server on :%d...\n", config.Server.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
