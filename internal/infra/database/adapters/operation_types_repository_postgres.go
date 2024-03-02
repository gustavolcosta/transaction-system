package adapters

import (
	"database/sql"
	"errors"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/infra/log_application"
)

type OperationTypesRepositoryPostgres struct {
	db         *sql.DB
	contextLog string
}

func NewOperationTypesRepositoryPostgres(db *sql.DB) *OperationTypesRepositoryPostgres {
	return &OperationTypesRepositoryPostgres{db: db, contextLog: "OPERATION_TYPE_REPOSITORY"}
}

func (opTypesRepository OperationTypesRepositoryPostgres) GetById(operationTypeId int) (*entities.OperationType, error) {

	query := "SELECT id, description, negative_amount from operations_types where id = $1"

	operationType := entities.OperationType{}

	err := opTypesRepository.db.QueryRow(query, operationTypeId).Scan(&operationType.Id, &operationType.Description, &operationType.NegativeAmount)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log_application.Error("Execute query to get operation type", err, opTypesRepository.contextLog)

		return nil, err
	}

	return &operationType, nil
}
