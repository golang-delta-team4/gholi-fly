package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	agencyDomain "gholi-fly-agancy/internal/agency/domain"
	agencyPort "gholi-fly-agancy/internal/agency/port"
	factorDomain "gholi-fly-agancy/internal/factor/domain"
	factorPort "gholi-fly-agancy/internal/factor/port"
	"gholi-fly-agancy/internal/reservation/domain"
	"gholi-fly-agancy/internal/reservation/port"
	tourDomain "gholi-fly-agancy/internal/tour/domain"
	tourPort "gholi-fly-agancy/internal/tour/port"
	bankPb "gholi-fly-agancy/pkg/adapters/clients/grpc/pb"
	bankClientPort "gholi-fly-agancy/pkg/adapters/clients/grpc/port"

	"github.com/google/uuid"
)

type ReservationService struct {
	tourSvc        tourPort.TourService
	reservationSvc port.ReservationService
	factorSvc      factorPort.FactorService
	agencySvc      agencyPort.AgencyService
	bankClient     bankClientPort.GRPCBankClient
}

func NewReservationService(
	tourSvc tourPort.TourService,
	reservationSvc port.ReservationService,
	factorSvc factorPort.FactorService,
	agencySvc agencyPort.AgencyService,
	bankClient bankClientPort.GRPCBankClient,
) *ReservationService {
	return &ReservationService{
		tourSvc:        tourSvc,
		reservationSvc: reservationSvc,
		factorSvc:      factorSvc,
		agencySvc:      agencySvc,
		bankClient:     bankClient,
	}
}

func (rs *ReservationService) ReserveTour(ctx context.Context, userID, tourID string, capacity uint, agencyID uuid.UUID) (string, error) {
	// Parse UUIDs
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return "", errors.New("invalid user ID format")
	}
	parsedTourID, err := uuid.Parse(tourID)
	if err != nil {
		return "", errors.New("invalid tour ID format")
	}

	// Step 1: Fetch the tour
	tour, err := rs.tourSvc.GetTourByID(ctx, tourDomain.TourID(parsedTourID))
	if err != nil {
		return "", fmt.Errorf("failed to fetch tour: %w", err)
	}
	if tour == nil {
		return "", errors.New("tour not found")
	}

	// Step 2: Fetch the agency for the tour
	agency, err := rs.agencySvc.GetAgencyByID(ctx, agencyDomain.AgencyID(agencyID))
	if err != nil {
		return "", fmt.Errorf("failed to fetch agency: %w", err)
	}

	// Step 3: Fetch the wallet of the agency owner
	ownerID := agency.OwnerID
	walletsResponse, err := rs.bankClient.GetUserWallets(&bankPb.GetWalletsRequest{
		OwnerId: ownerID.String(),
		Type:    bankPb.WalletType_PERSON,
	})
	if err != nil {
		return "", fmt.Errorf("failed to fetch agency owner wallets: %w", err)
	}
	if walletsResponse.Status != bankPb.ResponseStatus_SUCCESS || len(walletsResponse.Wallets) == 0 {
		return "", errors.New("no wallet found for the agency owner")
	}
	agencyWallet := walletsResponse.Wallets[0] // Use the first wallet found

	// Step 4: Validate tour availability
	if uint(tour.Capacity) < capacity {
		return "", errors.New("not enough capacity in the tour")
	}

	// Step 5: Create a reservation
	reservation := domain.Reservation{
		ID:         domain.ReservationID(uuid.New()),
		CustomerID: parsedUserID,
		TourID:     parsedTourID,
		Status:     "Pending", // Or any default status
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	reservationID, err := rs.reservationSvc.CreateReservation(ctx, reservation)
	if err != nil {
		return "", fmt.Errorf("failed to create reservation: %w", err)
	}

	// Step 6: Calculate pricing and profit
	totalAgencyPrice := uint64(tour.TripAgencyPrice+tour.HotelAgencyPrice) * uint64(capacity)
	profit := int(float64(totalAgencyPrice) * (agency.ProfitPercentage / 100))
	totalAmount := totalAgencyPrice + uint64(profit)

	// Step 7: Create a factor in the bank service
	factorRequest := &bankPb.CreateFactorRequest{
		FactorType: bankPb.FactorType_FACTOR_TYPE_SIMPLE,
		Factor: &bankPb.Factor{
			SourceService: "reservation-service",
			ExternalId:    reservationID.String(),
			BookingId:     reservationID.String(),
			TotalAmount:   totalAmount,
			Distributions: []*bankPb.Distribution{
				{
					WalletId: agencyWallet.Id,
					Amount:   totalAmount,
				},
			},
			Details:        fmt.Sprintf("Tour reservation for user %s", userID),
			InstantPayment: false,
			CustomerId:     userID,
		},
	}
	factorResponse, err := rs.bankClient.CreateFactor(factorRequest)
	if err != nil || factorResponse.Status != bankPb.ResponseStatus_SUCCESS {
		return "", fmt.Errorf("failed to create factor in the bank service: %w", err)
	}

	// Step 8: Create the factor entity after successful RPC call
	factor := factorDomain.Factor{
		ID:            factorDomain.FactorID(uuid.MustParse(factorResponse.Factor.Id)),
		ReservationID: uuid.UUID(reservationID),
		AgencyPrice:   totalAgencyPrice,
		Profit:        profit,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	_, err = rs.factorSvc.CreateFactor(ctx, factor)
	if err != nil {
		return "", fmt.Errorf("failed to create factor entity: %w", err)
	}

	// Step 9: Update tour capacity
	tour.Capacity -= int(capacity)
	if err := rs.tourSvc.UpdateTour(ctx, *tour); err != nil {
		return "", fmt.Errorf("failed to update tour capacity: %w", err)
	}

	return reservationID.String(), nil
}
