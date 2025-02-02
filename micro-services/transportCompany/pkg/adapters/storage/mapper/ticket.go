package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/ticket/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
)

func TicketDomain2Storage(ticketDomain domain.Ticket) *types.Ticket {
	return &types.Ticket{
		Id:        ticketDomain.Id,
		TripID:    ticketDomain.TripID,
		UserID:    ticketDomain.UserID,
		AgencyID:  ticketDomain.AgencyID,
		Price:     ticketDomain.Price,
		Status:    ticketDomain.Status,
		InvoiceId: ticketDomain.InvoiceId,
		Count:     ticketDomain.Count,
	}
}

func TicketStorage2Domain(ticketStorage types.Ticket) *domain.Ticket {
	return &domain.Ticket{
		Id:        ticketStorage.Id,
		TripID:    ticketStorage.TripID,
		UserID:    ticketStorage.UserID,
		AgencyID:  ticketStorage.AgencyID,
		Price:     ticketStorage.Price,
		Status:    ticketStorage.Status,
		InvoiceId: ticketStorage.InvoiceId,
		Count:     ticketStorage.Count,
		FactorId:  ticketStorage.Invoice.Info,
	}
}
