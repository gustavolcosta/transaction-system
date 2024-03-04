package dtos

type CreateTransactionInputDTO struct {
	AccountId       int     `json:"account_id" example:"1" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" example:"4" validate:"required"`
	Amount          float64 `json:"amount" example:"123.40" validate:"required"`
}

type CreateTransactionOutputDTO struct {
	TransactionId   int     `json:"id" example:"1"`
	AccountId       int     `json:"account_id" example:"1"`
	OperationTypeId int     `json:"operation_type_id" example:"4"`
	Amount          float64 `json:"amount" example:"123.4"`
}
