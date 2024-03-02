package dtos

type CreateTransactionInputDTO struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	TransactionId   int     `json:"id"`
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
