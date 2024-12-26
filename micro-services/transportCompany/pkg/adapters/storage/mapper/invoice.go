package mapper

import (
	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/domain"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/storage/types"
)

func InvoiceDomain2Storage(invoiceDomain domain.Invoice) *types.Invoice {
	return &types.Invoice{
		Id:         invoiceDomain.Id,
		IssuedDate: invoiceDomain.IssuedDate,
		TotalPrice: invoiceDomain.TotalPrice,
		Info:       invoiceDomain.Info,
		Status:     invoiceDomain.Status,
	}
}

func InvoiceStorage2Domain(invoiceStorage types.Invoice) *domain.Invoice {
	return &domain.Invoice{
		Id:         invoiceStorage.Id,
		IssuedDate: invoiceStorage.IssuedDate,
		Info:       invoiceStorage.Info,
		TotalPrice: invoiceStorage.TotalPrice,
		Status:     invoiceStorage.Status,
	}
}
