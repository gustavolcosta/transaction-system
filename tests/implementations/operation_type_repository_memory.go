package implementations

import (
	"errors"
	"transaction-system/domain/entities"
)

type OperationTypeRepositoryMemory struct {
	operationTypes []*entities.OperationType
}

func NewOperationTypeRepositoryMemory() *OperationTypeRepositoryMemory {
	operationTypes := make([]*entities.OperationType, 0)

	operationTypes = append(operationTypes,
		&entities.OperationType{
			Id:             1,
			Description:    "COMPRA A VISTA",
			NegativeAmount: true,
		}, &entities.OperationType{
			Id:             2,
			Description:    "COMPRA PARCELADA",
			NegativeAmount: true,
		}, &entities.OperationType{
			Id:             3,
			Description:    "SAQUE",
			NegativeAmount: true,
		}, &entities.OperationType{
			Id:             4,
			Description:    "PAGAMENTO",
			NegativeAmount: false,
		},
	)

	return &OperationTypeRepositoryMemory{operationTypes: operationTypes}
}

func (operationTypeRepository OperationTypeRepositoryMemory) GetById(operationTypeId int) (*entities.OperationType, error) {

	for _, operationType := range operationTypeRepository.operationTypes {

		if operationType.Id == operationTypeId {
			return operationType, nil
		}
	}

	return nil, errors.New("operation type not found")
}
