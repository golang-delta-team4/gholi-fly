package main

import (
	"flag"
	"log"
	"notification-nats/config"
	"notification-nats/database"
	shared "notification-nats/internal/outbox"
	"notification-nats/internal/queue"
	"notification-nats/models"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	config := config.MustReadConfig(*configPath)

	// 1. Connect to Postgres
	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatal("error connecting to db")
	}

	// 2. Connect to NATS
	conn, err := queue.CreateConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Drain() // Gracefully close on exit

	// 3. Create JetStream context
	js, err := queue.CreateJetStreamContext(conn)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Initialize your Outbox Processor with JetStream
	jobProcessor := shared.OutboxProcesser{
		DB:        db,
		JSContext: js,
		Subject:   "outbox.created", // or any subject pattern you want
	}

	// 5. Set up cron to run every 10s
	// a) handle outbox every 10s
	c := cron.New()
	_, err = c.AddFunc("@every 10s", jobProcessor.HandleOutboxMessage)
	if err != nil {
		log.Fatal("register handler error", err)
	}

	// b) New job: clean up processed outbox messages every 4 hours
	_, err = c.AddFunc("@every 4h", func() {
		cleanupProcessedOutboxMessages(db)
	})
	if err != nil {
		log.Fatal("register cleanup error", err)
	}

	log.Println("Start processing outbox messages")
	c.Start()
	defer c.Stop()

	// 6. Wait for terminate signal
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
	<-kill
}

func cleanupProcessedOutboxMessages(db *gorm.DB) {
	// Danger: this will permanently remove the rows!
	// Make sure you REALLY want them gone from outbox table.
	if err := db.Where("is_processed = ?", true).
		Delete(&models.OutBoxMessage{}).Error; err != nil {
		log.Println("cleanup processed error:", err)
	} else {
		log.Println("Cleaned up processed outbox messages...")
	}
}
