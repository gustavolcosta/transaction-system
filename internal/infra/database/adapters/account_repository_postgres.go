package adapters

import (
	"database/sql"
	"errors"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/infra/log_application"
)

type AccountRepositoryPostgres struct {
	db         *sql.DB
	contextLog string
}

func NewAccountRepositoryPostgres(db *sql.DB) *AccountRepositoryPostgres {
	return &AccountRepositoryPostgres{db: db, contextLog: "ACCOUNT_REPOSITORY"}
}

func (accountRepository AccountRepositoryPostgres) Create(account *entities.Account) error {

	insertQuery := "INSERT INTO accounts (document_number) VALUES ($1) RETURNING id"

	id := 0
	err := accountRepository.db.QueryRow(insertQuery, account.DocumentNumber).Scan(&id)

	if err != nil {
		log_application.Error("Save account in database", err, accountRepository.contextLog)
		return err
	}

	account.Id = id

	log_application.Info("Account saved successful!", accountRepository.contextLog, "accountId", account.Id)

	return nil
}

func (accountRepository AccountRepositoryPostgres) GetById(accountId int) (*entities.Account, error) {

	query := "SELECT id, document_number from accounts where id = $1"

	account := entities.Account{}

	err := accountRepository.db.QueryRow(query, accountId).Scan(&account.Id, &account.DocumentNumber)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log_application.Error("Execute query to get account", err, accountRepository.contextLog)
		return nil, err
	}

	return &account, nil
}
