package dtos

type CreateAccountInputDTO struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

type CreateAccountOutputDTO struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
