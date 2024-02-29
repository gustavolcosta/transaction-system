package implementations

import (
	"errors"
	"transaction-system/domain/entities"
)

type AccountRepositoryMemory struct {
	accounts []*entities.Account
	id       int
}

func NewAccountRepositoryMemory() *AccountRepositoryMemory {
	return &AccountRepositoryMemory{
		accounts: make([]*entities.Account, 0),
		id:       1,
	}
}

func (accountRepository *AccountRepositoryMemory) Create(account *entities.Account) error {

	account.Id = accountRepository.id
	_ = append(accountRepository.accounts, account)
	accountRepository.id++

	return nil
}

func (accountRepository *AccountRepositoryMemory) GetById(accountId int) (*entities.Account, error) {

	for _, account := range accountRepository.accounts {

		if account.Id == accountId {
			return account, nil
		}
	}

	return nil, errors.New("account not found")
}
