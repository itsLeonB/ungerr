package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (uee unprocessableEntityError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "UnprocessableEntityError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", uee.details)},
	}
}

func UnprocessableEntityError(details any) AppError {
	return unprocessableEntityError{details}
}
