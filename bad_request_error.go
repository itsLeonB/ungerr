package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (bre badRequestError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "BadRequestError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", bre.details)},
	}
}

func BadRequestError(details any) AppError {
	return badRequestError{details}
}
