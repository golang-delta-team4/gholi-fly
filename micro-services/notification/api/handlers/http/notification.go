package http

import (
	"log"
	"notification-nats/api/pb"
	"notification-nats/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetNotificationsByUserID(c *fiber.Ctx) error {
	userUUID, ok := c.Locals("UserUUID").(uuid.UUID)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	var notifications []models.NotificationHistory
	if err := h.DB.
		Where("user_id = ?", userUUID).
		Order("created_at DESC").
		Find(&notifications).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("No notifications found for user_id:", userUUID)
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"error": "No notifications found"})
		}
		log.Println("DB error:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Database error"})
	}

	if err := h.DB.Model(&models.NotificationHistory{}).
		Where("user_id = ?", userUUID).
		Update("is_read", true).Error; err != nil {
		log.Println("Update error:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to update is_read"})
	}

	var notificationList []*pb.Notification
	for _, r := range notifications {
		notificationList = append(notificationList, &pb.Notification{
			Id:        r.ID,
			UserId:    r.UserID.String(),
			Name:      r.Name,
			Message:   r.Message,
			Email:     r.Email,
			EventName: r.Event,
			IsRead:    r.IsRead,
			CreatedAt: r.CreatedAt.String(),
		})
	}

	return c.JSON(&pb.GetNotificationResponse{
		Notifications: notificationList,
	})
}
