package storage

import (
	"context"
	"errors"
	"gholi-fly-bank/internal/credit/domain"
	"gholi-fly-bank/internal/credit/port"
	"gholi-fly-bank/pkg/adapters/storage/mapper"
	"gholi-fly-bank/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type creditCardRepo struct {
	db *gorm.DB
}

// NewCreditCardRepo creates a new instance of the credit card repository.
func NewCreditCardRepo(db *gorm.DB) port.Repo {
	return &creditCardRepo{db: db}
}

func (r *creditCardRepo) Create(ctx context.Context, creditCardDomain domain.CreditCard) (domain.CreditCardUUID, error) {
	creditCard := mapper.CreditCardDomain2Storage(creditCardDomain)
	err := r.db.WithContext(ctx).Table("credit_cards").Create(creditCard).Error
	if err != nil {
		return domain.CreditCardUUID{}, err
	}
	return domain.CreditCardUUID(creditCard.ID), nil
}

func (r *creditCardRepo) GetByID(ctx context.Context, creditCardID domain.CreditCardUUID) (*domain.CreditCard, error) {
	var creditCard types.CreditCard
	err := r.db.WithContext(ctx).Table("credit_cards").Where("id = ?", creditCardID).First(&creditCard).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.CreditCardStorage2Domain(creditCard), nil
}

func (r *creditCardRepo) Get(ctx context.Context, filters domain.CreditCardFilters) ([]domain.CreditCard, error) {
	var creditCards []types.CreditCard
	query := r.db.WithContext(ctx).Table("credit_cards")

	// Apply filters
	if filters.CardNumber != "" {
		query = query.Where("card_number = ?", filters.CardNumber)
	}
	if filters.HolderName != "" {
		query = query.Where("holder_name = ?", filters.HolderName)
	}

	err := query.Find(&creditCards).Error
	if err != nil {
		return nil, err
	}

	return mapper.BatchCreditCardStorage2Domain(creditCards), nil
}

func (r *creditCardRepo) Update(ctx context.Context, creditCardDomain domain.CreditCard) error {
	creditCard := mapper.CreditCardDomain2Storage(creditCardDomain)
	result := r.db.WithContext(ctx).Table("credit_cards").
		Where("id = ?", creditCard.ID).
		Updates(creditCard)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("credit card not found")
	}

	return nil
}

func (r *creditCardRepo) Delete(ctx context.Context, creditCardID domain.CreditCardUUID) error {
	result := r.db.WithContext(ctx).Table("credit_cards").
		Where("id = ?", creditCardID).
		Delete(&types.CreditCard{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("credit card not found")
	}

	return nil
}
