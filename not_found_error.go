package ungerr

import "net/http"

type notFoundError struct {
	details any
}

func (nfe notFoundError) GrpcStatus() uint32 {
	return 5
}

func (nfe notFoundError) HttpStatus() int {
	return http.StatusNotFound
}

func (nfe notFoundError) Error() string {
	return http.StatusText(nfe.HttpStatus())
}

func (nfe notFoundError) Details() any {
	return nfe.details
}

func NotFoundError(details any) AppError {
	return notFoundError{details}
}
