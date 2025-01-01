package service

import (
	"context"
	"errors"
	pb "gholi-fly-agancy/api/pb"
	"gholi-fly-agancy/internal/agency/domain"
	"gholi-fly-agancy/internal/agency/port"
	staffDomain "gholi-fly-agancy/internal/staff/domain"
	staffPort "gholi-fly-agancy/internal/staff/port"
	"time"

	"github.com/google/uuid"
)

type AgencyService struct {
	agencySvc port.AgencyService
	staffSvc  staffPort.StaffService
}

func NewAgencyService(agencySvc port.AgencyService, staffSvc staffPort.StaffService) *AgencyService {
	return &AgencyService{
		agencySvc: agencySvc,
		staffSvc:  staffSvc,
	}
}

func (as *AgencyService) CreateAgency(ctx context.Context, userUUID string, req *pb.CreateAgencyRequest) (*pb.CreateAgencyResponse, error) {
	// Parse the UUID for the user
	ownerID, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, errors.New("invalid user UUID format")
	}

	// Parse the wallet ID
	walletID, err := uuid.Parse(req.GetWalletId())
	if err != nil {
		return nil, errors.New("invalid wallet ID format")
	}

	// Create the agency using the user UUID directly as the OwnerID
	agency := domain.Agency{
		Name:             req.GetName(),
		OwnerID:          ownerID, // OwnerID directly references user UUID
		WalletID:         walletID,
		ProfitPercentage: req.GetProfitPercentage(),
	}

	agencyID, err := as.agencySvc.CreateAgency(ctx, agency)
	if err != nil {
		return nil, err
	}

	// Add the owner as staff with role "Owner"
	staff := staffDomain.Staff{
		UserID:   ownerID,
		Role:     "Owner",
		WalletID: walletID,
		AgencyID: uuid.UUID(agencyID), // Link staff to the created agency
	}

	_, err = as.staffSvc.CreateStaff(ctx, staff)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAgencyResponse{Id: agencyID.String()}, nil
}

func (as *AgencyService) GetAgency(ctx context.Context, id string) (*pb.GetAgencyByIDResponse, error) {
	// Parse the UUID for the agency ID
	agencyID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid agency ID format")
	}

	// Call the service to retrieve the agency
	agency, err := as.agencySvc.GetAgencyByID(ctx, domain.AgencyID(agencyID))
	if err != nil {
		return nil, err
	}

	// Map the result to the response message
	return &pb.GetAgencyByIDResponse{
		Id:               agency.ID.String(),
		Name:             agency.Name,
		OwnerId:          agency.OwnerID.String(),
		WalletId:         agency.WalletID.String(),
		ProfitPercentage: agency.ProfitPercentage,
		CreatedAt:        agency.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        agency.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (as *AgencyService) UpdateAgency(ctx context.Context, id string, userUUID string, req *pb.UpdateAgencyRequest) (*pb.UpdateAgencyResponse, error) {
	// Parse the UUID for agency ID
	agencyID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid agency ID format")
	}

	// Retrieve the existing agency
	existingAgency, err := as.agencySvc.GetAgencyByID(ctx, domain.AgencyID(agencyID))
	if err != nil {
		return nil, err
	}

	// Update fields if provided in the request
	if req.GetName() != "" {
		existingAgency.Name = req.GetName()
	}

	if req.GetWalletId() != "" {
		walletID, err := uuid.Parse(req.GetWalletId())
		if err != nil {
			return nil, errors.New("invalid wallet ID format")
		}
		existingAgency.WalletID = walletID
	}

	if req.GetProfitPercentage() != 0 {
		existingAgency.ProfitPercentage = req.GetProfitPercentage()
	}

	// Check owner ID from userUUID
	ownerID, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, errors.New("invalid owner ID format")
	}
	existingAgency.OwnerID = ownerID

	// Update the agency details
	if err := as.agencySvc.UpdateAgency(ctx, *existingAgency); err != nil {
		return nil, err
	}

	return &pb.UpdateAgencyResponse{}, nil
}

func (as *AgencyService) DeleteAgency(ctx context.Context, id string) (*pb.DeleteAgencyResponse, error) {
	// Parse the UUID for agency ID
	agencyID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid agency ID format")
	}

	// Delete the agency
	if err := as.agencySvc.DeleteAgency(ctx, domain.AgencyID(agencyID)); err != nil {
		return nil, err
	}

	return &pb.DeleteAgencyResponse{}, nil
}
