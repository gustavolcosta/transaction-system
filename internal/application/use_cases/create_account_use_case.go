package use_cases

import (
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/domain/interfaces"
	"transaction-system/internal/infra/log_application"
)

type CreateAccountUseCase struct {
	accountRepository interfaces.AccountRepository
	contextLog        string
}

func NewCreateAccountUseCase(accountRepository interfaces.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{accountRepository: accountRepository, contextLog: "CREATE_ACCOUNT_USE_CASE"}
}

func (createAccount CreateAccountUseCase) Execute(inputDTO dtos.CreateAccountInputDTO) (*dtos.CreateAccountOutputDTO, error) {

	account, err := entities.NewAccount(inputDTO.DocumentNumber)

	if err != nil {
		log_application.Error("Instance a new Account", err, createAccount.contextLog)
		return nil, err
	}

	err = createAccount.accountRepository.Create(account)

	if err != nil {
		log_application.Error("Save Account", err, createAccount.contextLog)
		return nil, err
	}

	outputDTO := dtos.CreateAccountOutputDTO{
		AccountId:      account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return &outputDTO, nil
}
