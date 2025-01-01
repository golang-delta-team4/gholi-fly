package port

import (
	notificationPB "gholi-fly-hotel/pkg/adapters/clients/grpc/pb"
)

type GRPCNotificationClient interface {
	AddNotification(req *notificationPB.AddNotificationRequest) (*notificationPB.AddNotificationResponse, error)
}
