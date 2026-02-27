package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (ve validationError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "ValidationError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", ve.details)},
	}
}

func ValidationError(details any) AppError {
	return validationError{details}
}
