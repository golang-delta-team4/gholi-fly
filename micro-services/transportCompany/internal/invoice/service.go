package invoice

import (
	"context"
	"errors"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/port"
	"github.com/google/uuid"
)

var (
	ErrCreateInvoice = errors.New("error on create new invoice")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateInvoice(ctx context.Context, invoiceDomain domain.Invoice) (uuid.UUID, error) {
	invoiceId, err := s.repo.CreateInvoice(ctx, invoiceDomain)

	if err != nil {
		return uuid.Nil, err
	}
	return invoiceId, nil
}
