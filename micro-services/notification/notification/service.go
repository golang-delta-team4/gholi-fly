package notification

import (
	"context"
	"log"
	"notification-nats/shared"

	// This import path matches whatever was generated by your .proto go_package
	pb "notification-nats/pb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	pb.UnimplementedNotificationServiceServer
	DB *gorm.DB
}

func (s *Service) AddNotification(ctx context.Context, req *pb.AddNotificationRequest) (*pb.AddNotificationResponse, error) {
	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.AddNotificationResponse{
			Success: false,
			Error:   "ID Error: " + err.Error(),
		}, nil
	}

	//TODO: should handle name and email from your logic
	msg := shared.OutBoxMessage{
		ID:          uuid.NewString(),
		EventName:   req.EventName,
		UserID:      userUUID,
		Name:        "name",
		Email:       "example@email.com",
		Message:     req.Message,
		IsProcessed: false,
	}
	if err := s.DB.Create(&msg).Error; err != nil {
		log.Println("DB error:", err)
		return &pb.AddNotificationResponse{
			Success: false,
			Error:   "DB Error: " + err.Error(),
		}, nil
	}

	return &pb.AddNotificationResponse{
		Success: true,
		Error:   "",
	}, nil
}
