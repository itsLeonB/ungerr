package ungerr

import "net/http"

type unauthorizedError struct {
	details any
}

func (ue unauthorizedError) GrpcStatus() uint32 {
	return 16
}

func (ue unauthorizedError) HttpStatus() int {
	return http.StatusUnauthorized
}

func (ue unauthorizedError) Error() string {
	return http.StatusText(ue.HttpStatus())
}

func (ue unauthorizedError) Details() any {
	return ue.details
}

func UnauthorizedError(details any) AppError {
	return unauthorizedError{details}
}
