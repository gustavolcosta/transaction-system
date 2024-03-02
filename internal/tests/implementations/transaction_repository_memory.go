package implementations

import (
	"transaction-system/internal/domain/entities"
)

type TransactionRepositoryMemory struct {
	transactions []*entities.Transaction
	id           int
}

func NewTransactionRepositoryMemory() *TransactionRepositoryMemory {
	return &TransactionRepositoryMemory{
		transactions: make([]*entities.Transaction, 0),
		id:           1,
	}
}

func (transactionRepository TransactionRepositoryMemory) Create(transaction *entities.Transaction) error {

	transaction.Id = transactionRepository.id
	transactionRepository.transactions = append(transactionRepository.transactions, transaction)
	transactionRepository.id++

	return nil
}
