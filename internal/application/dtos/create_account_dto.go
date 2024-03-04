package dtos

type CreateAccountInputDTO struct {
	DocumentNumber string `json:"document_number" example:"1234567890" validate:"required"`
}

type CreateAccountOutputDTO struct {
	AccountId      int    `json:"account_id" example:"1"`
	DocumentNumber string `json:"document_number" example:"1234567890"`
}
