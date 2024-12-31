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
	adaptersPb "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/pb"
	grpcPort "github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/grpc/port"
	"github.com/google/uuid"
)

var (
	ErrBuyTicket    = errors.New("error on buy new ticket")
	ErrCancelTicket = errors.New("error on cancel ticket")
)

type service struct {
	repo           port.Repo
	tripService    tripPort.Service
	invoiceService invoicePort.Service
	bankGrpc       grpcPort.GRPCBankClient
}

func NewService(repo port.Repo,
	tripService tripPort.Service,
	invoiceService invoicePort.Service,
	bankGrpc grpcPort.GRPCBankClient) port.Service {
	return &service{
		repo:           repo,
		tripService:    tripService,
		invoiceService: invoiceService,
		bankGrpc:       bankGrpc,
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
	walletResponse, err := s.bankGrpc.GetWallets(&adaptersPb.GetWalletsRequest{
		OwnerId: ticket.UserID.String(),
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	totalPrice := 1 * trip.UserPrice
	response, err := s.bankGrpc.CreateFactor(&adaptersPb.CreateFactorRequest{
		Factor: &adaptersPb.Factor{
			SourceService: "transportCompany",
			TotalAmount:   uint64(trip.UserPrice),
			CustomerId:    ticket.UserID.String(),
			BookingId:     ticket.UserID.String(),
			ExternalId:    ticket.UserID.String(),
			Distributions: []*adaptersPb.Distribution{&adaptersPb.Distribution{
				WalletId: walletResponse.Wallets[0].Id,
				Amount:   uint64(totalPrice),
			}},
		},
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %w", ErrBuyTicket, err)
	}

	invoiceId, err := s.invoiceService.CreateInvoice(ctx, invoiceDomain.Invoice{
		IssuedDate: time.Now(),
		Status:     invoiceDomain.Paid,
		Info:       response.Factor.Id,
		TotalPrice: totalPrice,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	ticket.InvoiceId = invoiceId
	ticketId, err := s.repo.BuyTicket(ctx, ticket)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	s.bankGrpc.ApplyFactor(&adaptersPb.ApplyFactorRequest{
		FactorId: response.Factor.Id,
	})
	return ticketId, nil
}

func (s *service) BuyAgencyTicket(ctx context.Context, ticket domain.Ticket) (uuid.UUID, float64, error) {
	trip, err := s.tripService.GetTripById(ctx, ticket.TripID)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	if trip.SoldTickets+ticket.Count > trip.MaxTickets {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, "No more tickets available")
	}
	if trip.TourReleaseDate.After(time.Now()) {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, "trip is not released yet")
	}
	if trip.IsCanceled {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, "trip is canceled")
	}
	if !trip.IsConfirmed {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, "trip is not confirmed")
	}
	if trip.StartDate.Before(time.Now()) {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, "trip is started")
	}
	// bank
	walletResponse, err := s.bankGrpc.GetWallets(&adaptersPb.GetWalletsRequest{
		OwnerId: ticket.OwnerOfAgencyId.String(),
	})
	totalPrice := float64(ticket.Count) * trip.AgencyPrice
	response, err := s.bankGrpc.CreateFactor(&adaptersPb.CreateFactorRequest{
		Factor: &adaptersPb.Factor{
			SourceService: "transportCompany",
			TotalAmount:   uint64(trip.AgencyPrice),
			CustomerId:    ticket.OwnerOfAgencyId.String(),
			BookingId:     ticket.OwnerOfAgencyId.String(),
			ExternalId:    ticket.OwnerOfAgencyId.String(),
			Distributions: []*adaptersPb.Distribution{&adaptersPb.Distribution{
				WalletId: walletResponse.Wallets[0].Id,
				Amount:   uint64(totalPrice),
			}},
		},
	})
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("%w %w %s", ErrBuyTicket, err, response.Message)
	}

	invoiceId, err := s.invoiceService.CreateInvoice(ctx, invoiceDomain.Invoice{
		IssuedDate: time.Now(),
		Status:     invoiceDomain.Paid,
		TotalPrice: totalPrice,
	})
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	ticket.InvoiceId = invoiceId
	ticketId, err := s.repo.BuyAgencyTicket(ctx, ticket)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("%w %s", ErrBuyTicket, err)
	}
	s.bankGrpc.ApplyFactor(&adaptersPb.ApplyFactorRequest{
		FactorId: response.Factor.Id,
	})
	return ticketId, totalPrice, nil
}

func (s *service) CancelTicket(ctx context.Context, ticketId uuid.UUID) error {
	ticket, err := s.repo.GetTicket(ctx, ticketId)
	if err != nil {
		return fmt.Errorf("%w %s", ErrCancelTicket, err)
	}
	trip, err := s.tripService.GetTripById(ctx, ticket.TripID)
	if err != nil {
		return fmt.Errorf("%w %s", ErrCancelTicket, err)
	}
	if trip.StartDate.Before(time.Now()) {
		return fmt.Errorf("%w %s", ErrCancelTicket, "trip is started")
	}

	// TODO: Cancel factor
	err = s.repo.CancelTicket(ctx, ticketId, ticket.TripID)
	if err != nil {
		return fmt.Errorf("%w %s", ErrCancelTicket, err)
	}
	return nil
}

func (s *service) CancelAgencyTicket(ctx context.Context, ticketId uuid.UUID) error {
	ticket, err := s.repo.GetTicket(ctx, ticketId)
	if err != nil {
		return fmt.Errorf("%w %s", ErrCancelTicket, err)
	}
	trip, err := s.tripService.GetTripById(ctx, ticket.TripID)
	if err != nil {
		return fmt.Errorf("%w %s", ErrCancelTicket, err)
	}
	if trip.StartDate.Before(time.Now()) {
		return fmt.Errorf("%w %s", ErrCancelTicket, "trip is started")
	}

	// TODO: Cancel factor
	err = s.repo.CancelAgencyTicket(ctx, ticketId, ticket.TripID, ticket.Count)
	if err != nil {
		return fmt.Errorf("%w %s", ErrCancelTicket, err)
	}
	return nil
}
