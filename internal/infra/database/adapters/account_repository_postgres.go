package adapters

import (
	"database/sql"
	"log/slog"
	"transaction-system/internal/domain/entities"
)

var contextLog = "ACCOUNT_REPOSITORY"

type AccountRepositoryPostgres struct {
	db *sql.DB
}

func NewAccountRepositoryPostgres(db *sql.DB) *AccountRepositoryPostgres {
	return &AccountRepositoryPostgres{db: db}
}

func (accountRepository AccountRepositoryPostgres) Create(account *entities.Account) error {

	insertQuery := "INSERT INTO accounts (document_number) VALUES ($1) RETURNING id"

	id := 0
	err := accountRepository.db.QueryRow(insertQuery, account.DocumentNumber).Scan(&id)

	if err != nil {
		slog.Error("Save account in database:", "error:", err, "context:", contextLog)
		return err
	}

	account.Id = id

	slog.Info("Account saved successful!", "account_id", account.Id, "context", contextLog)

	return nil
}

func (accountRepository AccountRepositoryPostgres) GetById(accountId int) (*entities.Account, error) {
	return &entities.Account{}, nil
}
