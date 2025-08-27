package ungerr

import "net/http"

type badRequestError struct {
	details any
}

func (bre badRequestError) GrpcStatus() uint32 {
	return 3
}

func (bre badRequestError) HttpStatus() int {
	return http.StatusBadRequest
}

func (bre badRequestError) Error() string {
	return http.StatusText(bre.HttpStatus())
}

func (bre badRequestError) Details() any {
	return bre.details
}

func BadRequestError(details any) AppError {
	return badRequestError{details}
}
