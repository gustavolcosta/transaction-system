package use_cases

import (
	"errors"
	"log"
	"math"
	"time"
	"transaction-system/internal/application/dtos"
	entities2 "transaction-system/internal/domain/entities"
	interfaces2 "transaction-system/internal/domain/interfaces"
)

type CreateTransactionUseCase struct {
	transactionRepository   interfaces2.TransactionRepository
	accountRepository       interfaces2.AccountRepository
	operationTypeRepository interfaces2.OperationTypeRepository
}

func NewCreateTransactionUseCase(transactionRepository interfaces2.TransactionRepository, accountRepository interfaces2.AccountRepository,
	opTypeRepository interfaces2.OperationTypeRepository) *CreateTransactionUseCase {

	return &CreateTransactionUseCase{
		transactionRepository:   transactionRepository,
		accountRepository:       accountRepository,
		operationTypeRepository: opTypeRepository,
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

	transaction, err := entities2.NewTransaction(inputDTO.AccountId, operationType, inputDTO.Amount, time.Now())

	if err != nil {
		log.Printf("Error to instanciate a new transaction: %v", err)
		return nil, err
	}

	err = createTransaction.transactionRepository.Create(transaction)

	if err != nil {
		log.Printf("Error to save a new transaction: %v", err)
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
		log.Printf("Error to get account in create transaction: %v", err)
		return err
	}

	if account == nil {
		return errors.New("account not found")
	}

	return nil
}

func getOperationType(opTypeId int, createTransaction *CreateTransactionUseCase) (*entities2.OperationType, error) {

	operationType, err := createTransaction.operationTypeRepository.GetById(opTypeId)

	if err != nil {
		log.Printf("Error to get operation type in create transaction: %v", err)
		return nil, err
	}

	if operationType == nil {
		return nil, errors.New("operation type not found")
	}

	return operationType, nil
}
