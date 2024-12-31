package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	"github.com/google/uuid"
)

type Repo interface {
	BuyTicket(ctx context.Context, ticketDomain domain.Ticket) (uuid.UUID, error)
	BuyAgencyTicket(ctx context.Context, ticketDomain domain.Ticket) (uuid.UUID, error)
	GetTicket(ctx context.Context, ticketId uuid.UUID) (*domain.Ticket, error)
	CancelTicket(ctx context.Context, ticketId uuid.UUID, tripId uuid.UUID) error
	CancelAgencyTicket(ctx context.Context, ticketId uuid.UUID, tripId uuid.UUID, count uint) error
}
