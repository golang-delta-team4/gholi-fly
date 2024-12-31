package service

import (
	"context"
	"fmt"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/api/pb"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	ticketPort "github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	"github.com/google/uuid"
)

type TicketService struct {
	svc ticketPort.Service
}

func NewTicketService(svc ticketPort.Service) *TicketService {
	return &TicketService{
		svc: svc,
	}
}

var (
	ErrBuyTicket = ticket.ErrBuyTicket
)

func (s *TicketService) BuyTicket(ctx context.Context, req *pb.BuyTicketRequest, userId uuid.UUID) (*pb.BuyTicketResponse, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrBuyTicket, err)
	}

	ticketId, err := s.svc.BuyTicket(ctx, domain.Ticket{
		UserID: &userId,
		TripID: tripId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.BuyTicketResponse{
		TicketId: ticketId.String(),
	}, nil
}

func (s *TicketService) BuyAgencyTicket(ctx context.Context, req *pb.BuyAgencyTicketRequest) (*pb.BuyTicketResponse, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrBuyTicket, err)
	}

	agencyID, err := uuid.Parse(req.AgencyId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrBuyTicket, err)
	}

	ownerOfAgencyId, err := uuid.Parse(req.OwnerOfAgencyId)
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrBuyTicket, err)
	}

	ticketId, totalPrice, err := s.svc.BuyAgencyTicket(ctx, domain.Ticket{
		AgencyID:        &agencyID,
		TripID:          tripId,
		Count:           uint(req.TicketCount),
		OwnerOfAgencyId: ownerOfAgencyId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.BuyTicketResponse{
		TicketId:   ticketId.String(),
		TotalPrice: uint64(totalPrice),
	}, nil
}

func (s *TicketService) CancelTicket(ctx context.Context, ticketId string) (*pb.CancelTicketResponse, error) {
	ticketUId, err := uuid.Parse(ticketId)
	if err != nil {
		return nil, fmt.Errorf("error on parse ticket id: %w", err)
	}

	err = s.svc.CancelTicket(ctx, ticketUId)
	if err != nil {
		return nil, err
	}

	return &pb.CancelTicketResponse{}, nil
}

func (s *TicketService) CancelAgencyTicket(ctx context.Context, ticketId string) (*pb.CancelTicketResponse, error) {
	ticketUId, err := uuid.Parse(ticketId)
	if err != nil {
		return nil, fmt.Errorf("error on parse ticket id: %w", err)
	}

	err = s.svc.CancelAgencyTicket(ctx, ticketUId)
	if err != nil {
		return nil, err
	}

	return &pb.CancelTicketResponse{}, nil
}
