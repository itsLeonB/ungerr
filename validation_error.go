package ungerr

import "net/http"

type validationError struct {
	details any
}

func (ve validationError) GrpcStatus() uint32 {
	return 3
}

func (ve validationError) HttpStatus() int {
	return http.StatusUnprocessableEntity
}

func (ve validationError) Error() string {
	return http.StatusText(ve.HttpStatus())
}

func (ve validationError) Details() any {
	return ve.details
}

func ValidationError(details any) AppError {
	return validationError{details}
}
