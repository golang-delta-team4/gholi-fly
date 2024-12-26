package ticket

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	tripPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/google/uuid"
)

var (
	ErrBuyTicket = errors.New("error on buy new ticket")
)

type service struct {
	repo        port.Repo
	tripService tripPort.Service
}

func NewService(repo port.Repo, tripService tripPort.Service) port.Service {
	return &service{
		repo:        repo,
		tripService: tripService,
	}
}

func (s *service) BuyTicket(ctx context.Context, ticket domain.Ticket) (uuid.UUID, error) {
	trip, err := s.tripService.GetTripById(ctx, ticket.TripID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	if trip.SoldTickets+1 > trip.MaxTickets {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "No more tickets available")
	}
	if trip.TourReleaseDate.After(time.Now()) {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "Tour is not released yet")
	}
	// bank
	// user
	ticketId, err := s.repo.BuyTicket(ctx, ticket)
	if err != nil {
		return uuid.Nil, err
	}
	//make invoice
	return ticketId, nil
}
