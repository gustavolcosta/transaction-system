package use_cases

import (
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/domain/interfaces"
	"transaction-system/internal/infra/log_application"
)

var contextLog = "CREATE_ACCOUNT_USE_CASE"

type CreateAccountUseCase struct {
	accountRepository interfaces.AccountRepository
}

func NewCreateAccountUseCase(accountRepository interfaces.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{accountRepository: accountRepository}
}

func (createAccount CreateAccountUseCase) Execute(inputDTO dtos.CreateAccountInputDTO) (*dtos.CreateAccountOutputDTO, error) {

	account, err := entities.NewAccount(inputDTO.DocumentNumber)

	if err != nil {
		log_application.Error("Instance a new Account", err, contextLog)
		return nil, err
	}

	err = createAccount.accountRepository.Create(account)

	if err != nil {
		log_application.Error("Save Account", err, contextLog)
		return nil, err
	}

	outputDTO := dtos.CreateAccountOutputDTO{
		AccountId:      account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return &outputDTO, nil
}
