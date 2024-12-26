package storage

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/mapper"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type invoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	return &invoiceRepo{db}
}

func (r *invoiceRepo) CreateInvoice(ctx context.Context, invoiceDomain domain.Invoice) (uuid.UUID, error) {
	invoice := mapper.InvoiceDomain2Storage(invoiceDomain)
	return invoice.Id, r.db.Table("invoices").WithContext(ctx).Create(invoice).Error
}
