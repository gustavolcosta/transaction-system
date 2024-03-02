package interfaces

import (
	"transaction-system/internal/domain/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
}
