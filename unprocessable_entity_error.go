package ungerr

import "net/http"

type unprocessableEntityError struct {
	details any
}

func (uee unprocessableEntityError) GrpcStatus() uint32 {
	return 3
}

func (uee unprocessableEntityError) HttpStatus() int {
	return http.StatusUnprocessableEntity
}

func (uee unprocessableEntityError) Error() string {
	return http.StatusText(uee.HttpStatus())
}

func (uee unprocessableEntityError) Details() any {
	return uee.details
}

func UnprocessableEntityError(details any) AppError {
	return unprocessableEntityError{details}
}
