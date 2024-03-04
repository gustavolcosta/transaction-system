package use_cases

import (
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/exceptions"
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

	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, exceptions.NewNotFoundException("account not found")
	}

	outputDTO := dtos.GetAccountByIdOutputDTO{
		AccountId:      account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return &outputDTO, nil
}
