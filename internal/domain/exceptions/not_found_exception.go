package exceptions

type NotFoundException struct {
	Message string
}

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{Message: message}
}

func (notFoundException NotFoundException) Error() string {
	return notFoundException.Message
}
