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

	err := r.db.Exec("UPDATE trips SET sold_tickets = sold_tickets + 1 WHERE id = ?", ticketDomain.TripID).Error
	if err != nil {
		return uuid.Nil, err
	}

	return ticket.Id, r.db.Table("tickets").WithContext(ctx).Create(ticket).Error
}

func (r *ticketRepo) BuyAgencyTicket(ctx context.Context, ticketDomain domain.Ticket) (uuid.UUID, error) {
	ticket := mapper.TicketDomain2Storage(ticketDomain)

	err := r.db.Exec("UPDATE trips SET sold_tickets = sold_tickets + ? WHERE id = ?", ticketDomain.Count, ticketDomain.TripID).Error
	if err != nil {
		return uuid.Nil, err
	}

	return ticket.Id, r.db.Table("tickets").WithContext(ctx).Create(ticket).Error
}

func (r *ticketRepo) CancelTicket(ctx context.Context, ticketId uuid.UUID, tripId uuid.UUID) error {
	err := r.db.Exec("UPDATE trips SET sold_tickets = sold_tickets - 1 WHERE id = ?", tripId).Error
	if err != nil {
		return err
	}
	return r.db.Table("tickets").WithContext(ctx).Delete(&types.Ticket{}, "id = ?", tripId).Error
}

func (r *ticketRepo) CancelAgencyTicket(ctx context.Context, ticketId uuid.UUID, tripId uuid.UUID, count uint) error {
	err := r.db.Exec("UPDATE trips SET sold_tickets = sold_tickets - ? WHERE id = ?", count, tripId).Error
	if err != nil {
		return err
	}
	return r.db.Table("tickets").WithContext(ctx).Delete(&types.Ticket{}, "id = ?", ticketId).Error
}

func (r *ticketRepo) GetTicket(ctx context.Context, ticketId uuid.UUID) (*domain.Ticket, error) {
	var ticket types.Ticket
	err := r.db.Table("tickets").WithContext(ctx).Where("id = ?", ticketId).First(&ticket).Error
	if err != nil {
		return nil, err
	}

	domainTIcket := mapper.TicketStorage2Domain(ticket)
	return domainTIcket, nil
}
