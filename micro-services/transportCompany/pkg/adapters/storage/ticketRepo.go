package storage

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/mapper"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ticketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	return &ticketRepo{db}
}

func (r *ticketRepo) BuyTicket(ctx context.Context, ticketDomain domain.Ticket) (uuid.UUID, error) {
	ticket := mapper.TicketDomain2Storage(ticketDomain)

	err := r.db.Raw("UPDATE trips SET sold_tickets = sold_tickets + 1 WHERE id = ?", ticketDomain.Id).Error
	if err != nil {
		return uuid.Nil, err
	}

	return ticket.Id, r.db.Table("tickets").WithContext(ctx).Create(ticket).Error
}

func (r *ticketRepo) BuyAgencyTicket(ctx context.Context, ticketDomain domain.Ticket) (uuid.UUID, error) {
	ticket := mapper.TicketDomain2Storage(ticketDomain)

	err := r.db.Raw("UPDATE trips SET sold_tickets = sold_tickets + 1 WHERE id = ?", ticketDomain.Id).Error
	if err != nil {
		return uuid.Nil, err
	}

	return ticket.Id, r.db.Table("tickets").WithContext(ctx).Create(ticket).Error
}

func (r *ticketRepo) CancelTicket(ctx context.Context, ticketId uuid.UUID) error {
	return r.db.Table("tickets").WithContext(ctx).Delete(&types.Ticket{}, "id = ?", ticketId).Error
}
