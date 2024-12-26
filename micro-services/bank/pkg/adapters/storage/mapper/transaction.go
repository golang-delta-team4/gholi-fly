package mapper

import (
	"gholi-fly-bank/internal/transaction/domain"
	"gholi-fly-bank/pkg/adapters/storage/types"
	"gholi-fly-bank/pkg/fp"

	"github.com/google/uuid"
)

// TransactionDomain2Storage converts a Transaction from the domain layer to the storage layer.
func TransactionDomain2Storage(transactionDomain domain.Transaction) *types.Transaction {
	return &types.Transaction{
		ID:          uuid.UUID(transactionDomain.ID), // Cast domain ID to storage ID.
		WalletID:    transactionDomain.WalletID,
		FactorID:    &transactionDomain.FactorID, // Use pointer for nullable field.
		Amount:      transactionDomain.Amount,
		Type:        uint8(transactionDomain.Type),   // Cast domain.TransactionType to uint8.
		Status:      uint8(transactionDomain.Status), // Cast domain.TransactionStatus to uint8.
		Description: transactionDomain.Description,
		CreatedAt:   transactionDomain.CreatedAt,
		UpdatedAt:   transactionDomain.UpdatedAt,
	}
}

func transactionDomain2Storage(transactionDomain domain.Transaction) types.Transaction {
	return types.Transaction{
		ID:          uuid.UUID(transactionDomain.ID), // Cast domain ID to storage ID.
		WalletID:    transactionDomain.WalletID,
		FactorID:    &transactionDomain.FactorID, // Use pointer for nullable field.
		Amount:      transactionDomain.Amount,
		Type:        uint8(transactionDomain.Type),   // Cast domain.TransactionType to uint8.
		Status:      uint8(transactionDomain.Status), // Cast domain.TransactionStatus to uint8.
		Description: transactionDomain.Description,
		CreatedAt:   transactionDomain.CreatedAt,
		UpdatedAt:   transactionDomain.UpdatedAt,
	}
}

func BatchTransactionDomain2Storage(domains []domain.Transaction) []types.Transaction {
	return fp.Map(domains, transactionDomain2Storage)
}

// TransactionStorage2Domain converts a Transaction from the storage layer to the domain layer.
func TransactionStorage2Domain(transaction types.Transaction) *domain.Transaction {
	return &domain.Transaction{
		ID:          domain.TransactionUUID(transaction.ID), // Cast storage ID to domain ID.
		WalletID:    transaction.WalletID,
		FactorID:    *transaction.FactorID, // Dereference pointer.
		Amount:      transaction.Amount,
		Type:        domain.TransactionType(transaction.Type),     // Cast uint8 to domain.TransactionType.
		Status:      domain.TransactionStatus(transaction.Status), // Cast uint8 to domain.TransactionStatus.
		Description: transaction.Description,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func transactionStorage2Domain(transaction types.Transaction) domain.Transaction {
	return domain.Transaction{
		ID:          domain.TransactionUUID(transaction.ID), // Cast storage ID to domain ID.
		WalletID:    transaction.WalletID,
		FactorID:    *transaction.FactorID, // Dereference pointer.
		Amount:      transaction.Amount,
		Type:        domain.TransactionType(transaction.Type),     // Cast uint8 to domain.TransactionType.
		Status:      domain.TransactionStatus(transaction.Status), // Cast uint8 to domain.TransactionStatus.
		Description: transaction.Description,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func BatchTransactionStorage2Domain(transactions []types.Transaction) []domain.Transaction {
	return fp.Map(transactions, transactionStorage2Domain)
}
