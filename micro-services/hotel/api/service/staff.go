package service

import (
	"context"

	"gholi-fly-hotel/api/pb"
	"gholi-fly-hotel/internal/staff"
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

var (
	ErrStaffCreationValidation = staff.ErrStaffCreationValidation
	ErrStaffCreationDuplicate  = staff.ErrStaffCreationDuplicate
	ErrStaffOnCreate           = staff.ErrStaffCreation
	ErrStaffNotFound           = staff.ErrStaffNotFound
)

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

func (s *StaffService) GetAllStaffs(ctx context.Context, hotelId string) (*pb.GetAllStaffsResponse, error) {
	hotelUUID, err := uuid.Parse(hotelId)
	if err != nil {
		return nil, err
	}
	staffs, err := s.svc.GetAllStaffsByHotelID(ctx, hotelUUID)
	if err != nil {
		return nil, err
	}

	var staffList []*pb.Staff
	for _, h := range staffs {
		staffList = append(staffList, &pb.Staff{
			Id:        h.UUID.String(),
			Name:      h.Name,
			StaffType: int32(h.StaffType),
		})
	}

	return &pb.GetAllStaffsResponse{
		Staffs: staffList,
	}, nil
}

func (s *StaffService) GetStaffByID(ctx context.Context, staffID string) (*pb.Staff, error) {
	staffUUID, err := uuid.Parse(staffID)
	if err != nil {
		return nil, err
	}
	staff, err := s.svc.GetStaffByID(ctx, staffUUID)
	if err != nil {
		return nil, err
	}

	return &pb.Staff{
		Id:        staff.UUID.String(),
		Name:      staff.Name,
		StaffType: int32(staff.StaffType),
	}, nil
}
