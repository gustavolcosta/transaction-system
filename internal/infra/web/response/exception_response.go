package response

type ExceptionResponse struct {
	Message string `json:"message"`
}

func NewExceptionResponse(message string) *ExceptionResponse {
	return &ExceptionResponse{Message: message}
}
