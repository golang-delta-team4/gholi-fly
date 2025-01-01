package port

import (
	"context"
	"gholi-fly-bank/internal/credit/domain"
)

type Service interface {
	// Create a new credit card in the system.
	CreateCreditCard(ctx context.Context, creditCard domain.CreditCard) (domain.CreditCardUUID, error)

	// Retrieve a credit card by its ID.
	GetCreditCardByID(ctx context.Context, creditCardID domain.CreditCardUUID) (*domain.CreditCard, error)

	// Retrieve credit cards based on filters.
	GetCreditCards(ctx context.Context, filters domain.CreditCardFilters) ([]domain.CreditCard, error)

	// Update details of an existing credit card.
	UpdateCreditCard(ctx context.Context, creditCard domain.CreditCard) error

	// Delete a credit card by its ID.
	DeleteCreditCard(ctx context.Context, creditCardID domain.CreditCardUUID) error
}
