package response

type ExceptionResponse struct {
	Message string
}

func NewExceptionResponse(message string) *ExceptionResponse {
	return &ExceptionResponse{Message: message}
}
