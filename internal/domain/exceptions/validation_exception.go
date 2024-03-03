package exceptions

type ValidationException struct {
	Message string
}

func NewValidationException(message string) *ValidationException {
	return &ValidationException{Message: message}
}

func (validationException ValidationException) Error() string {
	return validationException.Message
}
