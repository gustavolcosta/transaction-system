package use_cases

import (
	"log"
	"transaction-system/application/dtos"
	"transaction-system/domain/entities"
	"transaction-system/domain/interfaces"
)

type CreateAccountUseCase struct {
	accountRepository interfaces.AccountRepository
}

func NewCreateAccountUseCase(accountRepository interfaces.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{accountRepository: accountRepository}
}

func (createAccount CreateAccountUseCase) Execute(inputDTO dtos.CreateAccountInputDTO) (*dtos.CreateAccountOutputDTO, error) {

	account, err := entities.NewAccount(inputDTO.DocumentNumber)

	if err != nil {
		log.Printf("Error during instance a new Account: %v", err)
		return nil, err
	}

	err = createAccount.accountRepository.Create(account)

	if err != nil {
		log.Printf("Error during save Account: %v", err)
		return nil, err
	}

	outputDTO := dtos.CreateAccountOutputDTO{
		AccountId:      account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return &outputDTO, nil
}
