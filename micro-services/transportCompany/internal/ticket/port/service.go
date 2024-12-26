package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	"github.com/google/uuid"
)

type Service interface {
	BuyTicket(ctx context.Context, ticket domain.Ticket) (uuid.UUID, error)
	BuyAgencyTicket(ctx context.Context, ticketDomain domain.Ticket) (uuid.UUID, error)
	CancelTicket(ctx context.Context, ticketId uuid.UUID) error
}
