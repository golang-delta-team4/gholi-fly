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
		UserID: userId,
		TripID: tripId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.BuyTicketResponse{
		TicketId: ticketId.String(),
	}, nil
}
