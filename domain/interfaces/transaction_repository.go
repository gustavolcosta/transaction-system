package interfaces

import "transaction-system/domain/entities"

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
}
