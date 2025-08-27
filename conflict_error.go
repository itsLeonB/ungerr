package ungerr

import "net/http"

type conflictError struct {
	details any
}

func (ce conflictError) GrpcStatus() uint32 {
	return 6
}

func (ce conflictError) HttpStatus() int {
	return http.StatusConflict
}

func (ce conflictError) Error() string {
	return http.StatusText(ce.HttpStatus())
}

func (ce conflictError) Details() any {
	return ce.details
}

func ConflictError(details any) AppError {
	return conflictError{details}
}
