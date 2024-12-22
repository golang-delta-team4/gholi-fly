package port

import (
	"context"
	"gholi-fly-bank/internal/credit/domain"
)

type Repo interface {
	// Create a new credit card in the system.
	Create(ctx context.Context, creditCard domain.CreditCard) (domain.CreditCardUUID, error)

	// Retrieve a credit card by its ID.
	GetByID(ctx context.Context, creditCardID domain.CreditCardUUID) (*domain.CreditCard, error)

	// Retrieve credit cards based on filters.
	Get(ctx context.Context, filters domain.CreditCardFilters) ([]domain.CreditCard, error)

	// Update details of an existing credit card.
	Update(ctx context.Context, creditCard domain.CreditCard) error

	// Delete a credit card by its ID.
	Delete(ctx context.Context, creditCardID domain.CreditCardUUID) error
}
