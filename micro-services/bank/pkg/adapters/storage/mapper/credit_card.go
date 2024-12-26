package mapper

import (
	"gholi-fly-bank/internal/credit/domain"
	"gholi-fly-bank/pkg/adapters/storage/types"
	"gholi-fly-bank/pkg/fp"

	"github.com/google/uuid"
)

// CreditCardDomain2Storage converts a CreditCard from the domain layer to the storage layer.
func CreditCardDomain2Storage(creditCardDomain domain.CreditCard) *types.CreditCard {
	return &types.CreditCard{
		ID:         uuid.UUID(creditCardDomain.ID), // Cast domain ID to types ID.
		CardNumber: creditCardDomain.CardNumber,
		ExpiryDate: creditCardDomain.ExpiryDate,
		CVV:        creditCardDomain.CVV,
		HolderName: creditCardDomain.HolderName,
		CreatedAt:  creditCardDomain.CreatedAt,
		UpdatedAt:  creditCardDomain.UpdatedAt,
	}
}

func creditCardDomain2Storage(creditCardDomain domain.CreditCard) types.CreditCard {
	return types.CreditCard{
		ID:         uuid.UUID(creditCardDomain.ID), // Cast domain ID to types ID.
		CardNumber: creditCardDomain.CardNumber,
		ExpiryDate: creditCardDomain.ExpiryDate,
		CVV:        creditCardDomain.CVV,
		HolderName: creditCardDomain.HolderName,
		CreatedAt:  creditCardDomain.CreatedAt,
		UpdatedAt:  creditCardDomain.UpdatedAt,
	}
}

func BatchCreditCardDomain2Storage(domains []domain.CreditCard) []types.CreditCard {
	return fp.Map(domains, creditCardDomain2Storage)
}

// CreditCardStorage2Domain converts a CreditCard from the storage layer to the domain layer.
func CreditCardStorage2Domain(creditCard types.CreditCard) *domain.CreditCard {
	return &domain.CreditCard{
		ID:         domain.CreditCardUUID(creditCard.ID), // Cast types ID to domain ID.
		CardNumber: creditCard.CardNumber,
		ExpiryDate: creditCard.ExpiryDate,
		CVV:        creditCard.CVV,
		HolderName: creditCard.HolderName,
		CreatedAt:  creditCard.CreatedAt,
		UpdatedAt:  creditCard.UpdatedAt,
	}
}

func creditCardStorage2Domain(creditCard types.CreditCard) domain.CreditCard {
	return domain.CreditCard{
		ID:         domain.CreditCardUUID(creditCard.ID), // Cast types ID to domain ID.
		CardNumber: creditCard.CardNumber,
		ExpiryDate: creditCard.ExpiryDate,
		CVV:        creditCard.CVV,
		HolderName: creditCard.HolderName,
		CreatedAt:  creditCard.CreatedAt,
		UpdatedAt:  creditCard.UpdatedAt,
	}
}

func BatchCreditCardStorage2Domain(cards []types.CreditCard) []domain.CreditCard {
	return fp.Map(cards, creditCardStorage2Domain)
}
