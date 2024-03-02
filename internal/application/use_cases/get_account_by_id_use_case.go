package use_cases

import (
	"errors"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/interfaces"
)

type GetAccountByIdUseCase struct {
	accountRepository interfaces.AccountRepository
}

func NewGetAccountByIdUseCase(accountRepository interfaces.AccountRepository) *GetAccountByIdUseCase {
	return &GetAccountByIdUseCase{accountRepository: accountRepository}
}

func (getAccount GetAccountByIdUseCase) Execute(accountId int) (*dtos.GetAccountByIdOutputDTO, error) {

	account, err := getAccount.accountRepository.GetById(accountId)

	if account == nil {
		return nil, errors.New("account not found")
	}

	if err != nil {
		return nil, err
	}

	outputDTO := dtos.GetAccountByIdOutputDTO{
		AccountId:      account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return &outputDTO, nil
}
