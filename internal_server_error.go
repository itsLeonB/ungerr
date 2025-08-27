package ungerr

import "net/http"

type internalServerError struct{}

func (ise internalServerError) GrpcStatus() uint32 {
	return 13
}

func (ise internalServerError) HttpStatus() int {
	return http.StatusInternalServerError
}

func (ise internalServerError) Error() string {
	return http.StatusText(ise.HttpStatus())
}

func (ise internalServerError) Details() any {
	return nil
}

func InternalServerError() AppError {
	return internalServerError{}
}
