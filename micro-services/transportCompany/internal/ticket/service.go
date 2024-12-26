package ticket

import (
	"context"
	"errors"
	"fmt"
	"time"

	invoiceDomain "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/domain"
	invoicePort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	tripPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/trip/port"
	"github.com/google/uuid"
)

var (
	ErrBuyTicket    = errors.New("error on buy new ticket")
	ErrCancelTicket = errors.New("error on buy new ticket")
)

type service struct {
	repo           port.Repo
	tripService    tripPort.Service
	invoiceService invoicePort.Service
}

func NewService(repo port.Repo,
	tripService tripPort.Service,
	invoiceService invoicePort.Service) port.Service {
	return &service{
		repo:           repo,
		tripService:    tripService,
		invoiceService: invoiceService,
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
	if trip.UserReleaseDate.After(time.Now()) {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is not released yet")
	}
	if trip.IsCanceled {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is canceled")
	}
	if !trip.IsConfirmed {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is not confirmed")
	}
	if trip.StartDate.Before(time.Now()) {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is started")
	}
	// bank
	// user

	invoiceId, err := s.invoiceService.CreateInvoice(ctx, invoiceDomain.Invoice{
		IssuedDate: time.Now(),
		Status:     invoiceDomain.Paid,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	ticket.InvoiceId = invoiceId
	ticketId, err := s.repo.BuyTicket(ctx, ticket)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}

	return ticketId, nil
}

func (s *service) BuyAgencyTicket(ctx context.Context, ticket domain.Ticket) (uuid.UUID, error) {
	trip, err := s.tripService.GetTripById(ctx, ticket.TripID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	if trip.SoldTickets+1 > trip.MaxTickets {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "No more tickets available")
	}
	if trip.TourReleaseDate.After(time.Now()) {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is not released yet")
	}
	if trip.IsCanceled {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is canceled")
	}
	if !trip.IsConfirmed {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is not confirmed")
	}
	if trip.StartDate.Before(time.Now()) {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, "trip is started")
	}
	// bank
	// user

	invoiceId, err := s.invoiceService.CreateInvoice(ctx, invoiceDomain.Invoice{
		IssuedDate: time.Now(),
		Status:     invoiceDomain.Paid,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	ticket.InvoiceId = invoiceId
	ticketId, err := s.repo.BuyAgencyTicket(ctx, ticket)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}

	return ticketId, nil
}

func (s *service) CancelTicket(ctx context.Context, ticketId uuid.UUID) error {
	trip, err := s.tripService.GetTripById(ctx, ticketId)
	if err != nil {
		return fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	if trip.StartDate.Before(time.Now()) {
		return fmt.Errorf("%w %s", ErrBuyTicket, "trip is started")
	}

	// give  back bank
	err = s.repo.CancelTicket(ctx, ticketId)
	if err != nil {
		return fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	return nil
}
