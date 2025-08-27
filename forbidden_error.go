package ungerr

import "net/http"

type forbiddenError struct {
	details any
}

func (fe forbiddenError) GrpcStatus() uint32 {
	return 7
}

func (fe forbiddenError) HttpStatus() int {
	return http.StatusForbidden
}

func (fe forbiddenError) Error() string {
	return http.StatusText(fe.HttpStatus())
}

func (fe forbiddenError) Details() any {
	return fe.details
}

func ForbiddenError(details any) AppError {
	return forbiddenError{details}
}
