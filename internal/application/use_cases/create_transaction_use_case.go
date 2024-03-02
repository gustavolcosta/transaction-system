package use_cases

import (
	"errors"
	"math"
	"time"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/domain/interfaces"
	"transaction-system/internal/infra/log_application"
)

type CreateTransactionUseCase struct {
	transactionRepository   interfaces.TransactionRepository
	accountRepository       interfaces.AccountRepository
	operationTypeRepository interfaces.OperationTypeRepository
	contextLog              string
}

func NewCreateTransactionUseCase(transactionRepository interfaces.TransactionRepository, accountRepository interfaces.AccountRepository,
	opTypeRepository interfaces.OperationTypeRepository) *CreateTransactionUseCase {

	return &CreateTransactionUseCase{
		transactionRepository:   transactionRepository,
		accountRepository:       accountRepository,
		operationTypeRepository: opTypeRepository,
		contextLog:              "CREATE_TRANSACTION_USE_CASE",
	}
}

func (createTransaction *CreateTransactionUseCase) Execute(inputDTO dtos.CreateTransactionInputDTO) (*dtos.CreateTransactionOutputDTO, error) {

	err := verifyAccount(inputDTO.AccountId, createTransaction)

	if err != nil {
		return nil, err
	}

	operationType, err := getOperationType(inputDTO.OperationTypeId, createTransaction)

	if err != nil {
		return nil, err
	}

	transaction, err := entities.NewTransaction(inputDTO.AccountId, operationType, inputDTO.Amount, time.Now())

	if err != nil {
		log_application.Error("Create a instance of a new transaction", err, createTransaction.contextLog)
		return nil, err
	}

	err = createTransaction.transactionRepository.Create(transaction)

	if err != nil {
		log_application.Error("Save a new transaction", err, createTransaction.contextLog)
		return nil, err
	}

	outputDTO := dtos.CreateTransactionOutputDTO{
		TransactionId:   transaction.Id,
		AccountId:       transaction.AccountId,
		OperationTypeId: transaction.OperationType.Id,
		Amount:          math.Abs(transaction.Amount), //To not return a negative value
	}

	return &outputDTO, nil

}

func verifyAccount(accountId int, createTransaction *CreateTransactionUseCase) error {
	account, err := createTransaction.accountRepository.GetById(accountId)

	if err != nil {
		log_application.Error("To get account in create transaction", err, createTransaction.contextLog)
		return err
	}

	if account == nil {
		return errors.New("account not found")
	}

	return nil
}

func getOperationType(opTypeId int, createTransaction *CreateTransactionUseCase) (*entities.OperationType, error) {

	operationType, err := createTransaction.operationTypeRepository.GetById(opTypeId)

	if err != nil {
		log_application.Error("To get operation in create transaction", err, createTransaction.contextLog)
		return nil, err
	}

	if operationType == nil {
		return nil, errors.New("operation type not found")
	}

	return operationType, nil
}
