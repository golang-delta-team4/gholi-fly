package credit

import (
	"context"
	"errors"
	"fmt"
	"gholi-fly-bank/internal/credit/domain"
	"gholi-fly-bank/internal/credit/port"
	"log"
)

var (
	ErrCreditCardCreation      = errors.New("error on creating credit card")
	ErrCreditCardValidation    = errors.New("credit card validation failed")
	ErrCreditCardNotFound      = errors.New("credit card not found")
	ErrCreditCardUpdate        = errors.New("error updating credit card")
	ErrCreditCardDeletion      = errors.New("error deleting credit card")
	ErrInvalidCreditCardNumber = errors.New("invalid credit card number")
	ErrInvalidCreditCardExpiry = errors.New("invalid expiry date format")
)

type service struct {
	repo port.Repo
}

// NewService creates a new instance of the credit card service.
func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCreditCard(ctx context.Context, creditCard domain.CreditCard) (domain.CreditCardUUID, error) {
	// Validate credit card details.
	if len(creditCard.CardNumber) != 16 {
		return domain.CreditCardUUID{}, fmt.Errorf("%w: card number must be 16 digits", ErrInvalidCreditCardNumber)
	}
	if !isValidExpiryDate(creditCard.ExpiryDate) {
		return domain.CreditCardUUID{}, ErrInvalidCreditCardExpiry
	}

	creditCardID, err := s.repo.Create(ctx, creditCard)
	if err != nil {
		log.Println("error creating credit card:", err.Error())
		return domain.CreditCardUUID{}, ErrCreditCardCreation
	}

	return creditCardID, nil
}

func (s *service) GetCreditCardByID(ctx context.Context, creditCardID domain.CreditCardUUID) (*domain.CreditCard, error) {
	creditCard, err := s.repo.GetByID(ctx, creditCardID)
	if err != nil {
		log.Println("error fetching credit card by ID:", err.Error())
		return nil, err
	}

	if creditCard == nil {
		return nil, ErrCreditCardNotFound
	}

	return creditCard, nil
}

func (s *service) GetCreditCards(ctx context.Context, filters domain.CreditCardFilters) ([]domain.CreditCard, error) {
	creditCards, err := s.repo.Get(ctx, filters)
	if err != nil {
		log.Println("error fetching credit cards:", err.Error())
		return nil, err
	}

	return creditCards, nil
}

func (s *service) UpdateCreditCard(ctx context.Context, creditCard domain.CreditCard) error {
	// Validate updated details if needed.
	if len(creditCard.CardNumber) != 16 {
		return fmt.Errorf("%w: card number must be 16 digits", ErrInvalidCreditCardNumber)
	}
	if !isValidExpiryDate(creditCard.ExpiryDate) {
		return ErrInvalidCreditCardExpiry
	}

	err := s.repo.Update(ctx, creditCard)
	if err != nil {
		log.Println("error updating credit card:", err.Error())
		return ErrCreditCardUpdate
	}

	return nil
}

func (s *service) DeleteCreditCard(ctx context.Context, creditCardID domain.CreditCardUUID) error {
	err := s.repo.Delete(ctx, creditCardID)
	if err != nil {
		log.Println("error deleting credit card:", err.Error())
		return ErrCreditCardDeletion
	}

	return nil
}

// isValidExpiryDate validates the credit card expiry date format (MM/YY).
func isValidExpiryDate(expiryDate string) bool {
	if len(expiryDate) != 5 || expiryDate[2] != '/' {
		return false
	}
	// Additional checks (e.g., valid month/year) can be added here.
	return true
}
