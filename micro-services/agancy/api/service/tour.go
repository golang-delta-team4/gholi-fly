package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "gholi-fly-agancy/api/pb"
	agencyDomain "gholi-fly-agancy/internal/agency/domain"
	agencyEventSvc "gholi-fly-agancy/internal/agency/port"
	"gholi-fly-agancy/internal/tour/domain"
	"gholi-fly-agancy/internal/tour/port"
	tourEventDomain "gholi-fly-agancy/internal/tour_event/domain"
	tourEventSvc "gholi-fly-agancy/internal/tour_event/port"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type BookingCreateRequest struct {
	CheckIn  string   `json:"checkIn"`
	CheckOut string   `json:"checkOut"`
	HotelId  string   `json:"hotelId"`
	RoomIds  []string `json:"roomIds"`
	UserId   string   `json:"userId,omitempty"`
	AgencyId string   `json:"agencyId,omitempty"`
}

type BookingCreateResponse struct {
	ReservationId string `json:"reservationId"`
	TotalPrice    int64  `json:"totalPrice"`
}

type TourService struct {
	staffSvc       port.TourService
	tourEventSvc   tourEventSvc.TourEventService
	agencyEventSvc agencyEventSvc.AgencyService
}

func NewTourService(staffSvc port.TourService, tourEventSvc tourEventSvc.TourEventService, agencyEventSvc agencyEventSvc.AgencyService) *TourService {
	return &TourService{
		staffSvc:       staffSvc,
		tourEventSvc:   tourEventSvc,
		agencyEventSvc: agencyEventSvc,
	}
}
func (ts *TourService) CreateTour(ctx context.Context, agencyID string, hotelHost string, hotelPort uint, transportCompanyHost string, transportCompanyPort uint, req *pb.CreateTourRequest) (*pb.CreateTourResponse, error) {
	tracId := uuid.MustParse(uuid.NewString())
	// Parse agency ID
	parsedAgencyID, err := uuid.Parse(agencyID)
	if err != nil {
		return nil, errors.New("invalid agency ID format")
	}

	// Parse trip ID
	tripID, err := uuid.Parse(req.GetTripId())
	if err != nil {
		return nil, errors.New("invalid trip ID format")
	}
	var events = make([]tourEventDomain.TourEvent, 0)
	agUUID, err := uuid.Parse(agencyID)
	if err != nil {
		return nil, errors.New("invalid agency ID format")
	}
	agency, err := ts.agencyEventSvc.GetAgencyByID(ctx, agencyDomain.AgencyID(agUUID))
	if err != nil {
		return nil, err
	}
	// Step 1: Call transport company to buy tickets
	compensationPayload := req.TripId
	transportURL := fmt.Sprintf("http://%s:%d/api/v1/transport-company/ticket/agency-buy", transportCompanyHost, transportCompanyPort)

	// Step 2: Call hotel service to book a hotel
	compensationPayload = req.HotelId
	hotelURL := fmt.Sprintf("http://%s:%d/api/v1/hotel/booking/%s", hotelHost, hotelPort, req.GetHotelId())
	hotelRequest := BookingCreateRequest{
		CheckIn:  req.GetCheckIn(),
		CheckOut: req.GetCheckOut(),
		HotelId:  req.GetHotelId(),
		RoomIds:  req.GetRoomIds(),
		UserId:   agency.OwnerID.String(),
	}
	var hotelResponse BookingCreateResponse
	err = makePostRequest(ctx, hotelURL, hotelRequest, &hotelResponse)
	if err != nil {

		return nil, fmt.Errorf("failed to book hotel: %w", err)
	}
	events = append(events, tourEventDomain.TourEvent{
		ReservationID:       tracId,
		EventType:           tourEventDomain.EventTypeHotelReservation,
		Status:              tourEventDomain.StatusSuccess,
		CompensationPayload: hotelResponse.ReservationId,
		Payload:             compensationPayload,
	})

	transportRequest := map[string]interface{}{
		"tripId":          tripID.String(),
		"agencyId":        parsedAgencyID.String(),
		"ticketCount":     req.GetTicketCount(),
		"ownerOfAgencyId": agency.OwnerID.String(), // ownerid
	}
	var transportResponse struct {
		TicketId   string `json:"ticketId"`
		TotalPrice int    `json:"totalPrice"`
	}
	err = makePostRequest(ctx, transportURL, transportRequest, &transportResponse)
	if err != nil {
		events = append(events, tourEventDomain.TourEvent{
			ReservationID:       tracId,
			EventType:           tourEventDomain.EventTypeTripReservation,
			Status:              tourEventDomain.StatusFailed,
			CompensationPayload: transportResponse.TicketId,
			Payload:             compensationPayload,
		})
		err2 := ts.tourEventSvc.CreateEvent(ctx, events)
		return nil, fmt.Errorf("failed to buy transport ticket: %w", err2)
	}

	startDate, err := time.Parse(time.RFC3339, req.GetStartDate())
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	endDate, err := time.Parse(time.RFC3339, req.GetEndDate())
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	// Step 3: Create the tour entity
	tour := domain.Tour{
		Name:                req.GetName(),
		Description:         req.GetDescription(),
		StartDate:           startDate,
		EndDate:             endDate,
		SourceLocation:      req.GetSourceLocation(),
		DestinationLocation: req.GetDestinationLocation(),
		TripID:              tripID,
		TripAgencyPrice:     transportResponse.TotalPrice,
		HotelID:             uuid.MustParse(req.GetHotelId()),
		HotelAgencyPrice:    int(hotelResponse.TotalPrice),
		IsPublished:         req.GetIsPublished(),
		Capacity:            int(req.GetCapacity()),
	}

	// Save to the database
	tourID, err := ts.staffSvc.CreateTour(ctx, tour)
	if err != nil {
		return nil, fmt.Errorf("failed to create tour: %w", err)
	}

	return &pb.CreateTourResponse{Id: tourID.String()}, nil
}

func (ts *TourService) UpdateTour(ctx context.Context, id string, req *pb.UpdateTourRequest) (*pb.UpdateTourResponse, error) {
	// Parse the UUID for tour ID
	tourID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid tour ID format")
	}

	// Retrieve the existing tour
	existingTour, err := ts.staffSvc.GetTourByID(ctx, domain.TourID(tourID))
	if err != nil {
		return nil, err
	}

	if existingTour == nil {
		return nil, errors.New("tour not found")
	}

	// Update fields if provided in the request
	if req.GetName() != "" {
		existingTour.Name = req.GetName()
	}

	if req.GetDescription() != "" {
		existingTour.Description = req.GetDescription()
	}

	if req.GetStartDate() != "" {
		startDate, err := time.Parse(time.RFC3339, req.GetStartDate())
		if err != nil {
			return nil, errors.New("invalid start date format")
		}
		existingTour.StartDate = startDate
	}

	if req.GetEndDate() != "" {
		endDate, err := time.Parse(time.RFC3339, req.GetEndDate())
		if err != nil {
			return nil, errors.New("invalid end date format")
		}
		existingTour.EndDate = endDate
	}

	if req.GetSourceLocation() != "" {
		existingTour.SourceLocation = req.GetSourceLocation()
	}

	if req.GetDestinationLocation() != "" {
		existingTour.DestinationLocation = req.GetDestinationLocation()
	}

	if req.GetTripId() != "" {
		tripID, err := uuid.Parse(req.GetTripId())
		if err != nil {
			return nil, errors.New("invalid trip ID format")
		}
		existingTour.TripID = tripID
	}

	if req.GetTicketCount() != 0 {
		existingTour.Capacity = int(req.GetTicketCount())
	}

	if req.GetTripAgencyPrice() != 0 {
		existingTour.TripAgencyPrice = int(req.GetTripAgencyPrice())
	}

	if req.GetHotelId() != "" {
		hotelID, err := uuid.Parse(req.GetHotelId())
		if err != nil {
			return nil, errors.New("invalid hotel ID format")
		}
		existingTour.HotelID = hotelID
	}

	if req.GetHotelAgencyPrice() != 0 {
		existingTour.HotelAgencyPrice = int(req.GetHotelAgencyPrice())
	}

	if req.GetCapacity() != 0 {
		existingTour.Capacity = int(req.GetCapacity())
	}

	existingTour.IsPublished = req.GetIsPublished()

	// Update the tour details
	if err := ts.staffSvc.UpdateTour(ctx, *existingTour); err != nil {
		return nil, fmt.Errorf("failed to update tour: %w", err)
	}

	return &pb.UpdateTourResponse{}, nil
}

func (ts *TourService) DeleteTour(ctx context.Context, id string) (*pb.DeleteTourResponse, error) {
	// Parse the UUID for tour ID
	tourID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid tour ID format")
	}

	// Delete the tour
	if err := ts.staffSvc.DeleteTour(ctx, domain.TourID(tourID)); err != nil {
		return nil, fmt.Errorf("failed to delete tour: %w", err)
	}

	return &pb.DeleteTourResponse{}, nil
}

func (ts *TourService) GetTourByID(ctx context.Context, id string) (*pb.GetTourByIDResponse, error) {
	// Parse the UUID for tour ID
	tourID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid tour ID format")
	}

	// Retrieve the tour
	tour, err := ts.staffSvc.GetTourByID(ctx, domain.TourID(tourID))
	if err != nil {
		return nil, err
	}

	if tour == nil {
		return nil, errors.New("tour not found")
	}

	// Map the result to the response message
	return &pb.GetTourByIDResponse{
		Id:                  tour.ID.String(),
		Name:                tour.Name,
		Description:         tour.Description,
		StartDate:           tour.StartDate.Format(time.RFC3339),
		EndDate:             tour.EndDate.Format(time.RFC3339),
		SourceLocation:      tour.SourceLocation,
		DestinationLocation: tour.DestinationLocation,
		TripId:              tour.TripID.String(),
		TripAgencyPrice:     uint64(tour.TripAgencyPrice),
		HotelId:             tour.HotelID.String(),
		HotelAgencyPrice:    uint64(tour.HotelAgencyPrice),
		Capacity:            uint32(tour.Capacity),
		IsPublished:         tour.IsPublished,
		CreatedAt:           tour.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           tour.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (ts *TourService) ListToursByAgency(ctx context.Context, agencyID string) ([]*pb.GetTourByIDResponse, error) {
	// Parse the UUID for agency ID
	parsedAgencyID, err := uuid.Parse(agencyID)
	if err != nil {
		return nil, errors.New("invalid agency ID format")
	}

	// Retrieve the list of tours
	tours, err := ts.staffSvc.ListToursByAgency(ctx, parsedAgencyID)
	if err != nil {
		return nil, fmt.Errorf("failed to list tours by agency: %w", err)
	}

	// Map each tour to the response message
	response := make([]*pb.GetTourByIDResponse, 0, len(tours))
	for _, tour := range tours {
		response = append(response, &pb.GetTourByIDResponse{
			Id:                  tour.ID.String(),
			Name:                tour.Name,
			Description:         tour.Description,
			StartDate:           tour.StartDate.Format(time.RFC3339),
			EndDate:             tour.EndDate.Format(time.RFC3339),
			SourceLocation:      tour.SourceLocation,
			DestinationLocation: tour.DestinationLocation,
			TripId:              tour.TripID.String(),
			TripAgencyPrice:     uint64(tour.TripAgencyPrice),
			HotelId:             tour.HotelID.String(),
			HotelAgencyPrice:    uint64(tour.HotelAgencyPrice),
			Capacity:            uint32(tour.Capacity),
			IsPublished:         tour.IsPublished,
			CreatedAt:           tour.CreatedAt.Format(time.RFC3339),
			UpdatedAt:           tour.UpdatedAt.Format(time.RFC3339),
		})
	}

	return response, nil
}
func makePostRequest(ctx context.Context, url string, requestBody interface{}, responseBody interface{}) error {
	reqBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP request failed with status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(responseBody)
}
