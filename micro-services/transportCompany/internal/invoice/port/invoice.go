package port

import (
	"context"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/internal/invoice/domain"
	"github.com/google/uuid"
)

type Repo interface {
	CreateInvoice(ctx context.Context, invoiceDomain domain.Invoice) (uuid.UUID, error)
}
