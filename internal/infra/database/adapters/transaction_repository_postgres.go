package adapters

import (
	"database/sql"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/infra/log_application"
)

type TransactionRepositoryPostgres struct {
	db         *sql.DB
	contextLog string
}

func NewTransactionRepositoryPostgres(db *sql.DB) *TransactionRepositoryPostgres {
	return &TransactionRepositoryPostgres{db: db, contextLog: "TRANSACTION_REPOSITORY"}
}

func (transactionRepository TransactionRepositoryPostgres) Create(transaction *entities.Transaction) error {

	insertQuery := "INSERT INTO transactions (account_id, operation_type_id, amount, event_date) values ($1, $2, $3, $4) RETURNING id"

	id := 0
	err := transactionRepository.db.QueryRow(insertQuery, transaction.AccountId, transaction.OperationType.Id, transaction.Amount, transaction.EventDate).Scan(&id)

	if err != nil {
		log_application.Error("Save transaction in database", err, transactionRepository.contextLog)
		return err
	}

	transaction.Id = id

	log_application.Info("Transaction saved successful!", transactionRepository.contextLog, "transactionId", transaction.Id)

	return nil
}
