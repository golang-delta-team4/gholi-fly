package outbox

import (
	"encoding/json"
	"log"
	"time"

	"notification-nats/models"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type OutboxProcesser struct {
	DB        *gorm.DB
	JSContext nats.JetStreamContext
	Subject   string
}

func (p *OutboxProcesser) HandleOutboxMessage() {
	messages := make([]models.OutBoxMessage, 0)
	err := p.DB.
		Where("is_processed = ?", false).
		Find(&messages).Error
	if err != nil {
		log.Println("query outbox messages error: ", err)
		return
	}

	// no waiting message
	if len(messages) == 0 {
		return
	}

	// Publish each message.
	// If success, add to processed slice
	processedID := make([]string, 0)
	for _, m := range messages {
		b, err := json.Marshal(m)
		if err != nil {
			continue
		}

		// publish message to a queue
		if err := p.publishMessage(b); err != nil {
			log.Println("publish outbox message error: ", err)
			continue
		}

		processedID = append(processedID, m.ID)

		record := models.NotificationHistory{
			ID:        m.ID,
			UserID:    m.UserID,
			Name:      m.Name,
			Email:     m.Email,
			Event:     m.EventName,
			Message:   m.Message,
			Is_read:   false,
			CreatedAt: time.Now(),
		}
		if err := p.DB.Create(&record).Error; err != nil {
			log.Println("Archiving error:", err)
			// Decide if you want to skip processing, or just log and continue
			continue
		}
	}

	// Update processed messages in database
	err = p.DB.Model(&models.OutBoxMessage{}).
		Where("id IN ?", processedID).
		UpdateColumn("is_processed", true).Error
	if err != nil {
		log.Println("update outbox error: ", err)
		return
	}

	log.Println("Published messages:", processedID)
}

func (p *OutboxProcesser) publishMessage(body []byte) error {
	_, err := p.JSContext.Publish(p.Subject, body)
	return err
}
