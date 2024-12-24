package service

import (
	"context"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/staff/domain"
	staffPort "gholi-fly-hotel/internal/staff/port"

	"github.com/google/uuid"
)

type StaffService struct {
	svc staffPort.Service
	// notifSvc              notifPort.Service
}

func NewStaffService(svc staffPort.Service) *StaffService {
	return &StaffService{
		svc: svc,
	}
}

func (s *StaffService) CreateStaff(ctx context.Context, req *pb.StaffCreateRequest, hotelID string) (*pb.StaffCreateResponse, error) {
	hotelUUID, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, err
	}
	staffId, err := s.svc.CreateStaffByHotelID(ctx, domain.Staff{
		HotelID:   hotelUUID,
		Name:      req.Name,
		StaffType: uint8(req.StaffType),
	}, hotelUUID)

	if err != nil {
		return nil, err
	}

	return &pb.StaffCreateResponse{
		StaffId: staffId.String(),
	}, nil
}
