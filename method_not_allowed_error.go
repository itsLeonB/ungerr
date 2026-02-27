package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type methodNotAllowedError struct {
	details any
}

func (bre methodNotAllowedError) GrpcStatus() uint32 {
	return 12
}

func (bre methodNotAllowedError) HttpStatus() int {
	return http.StatusMethodNotAllowed
}

func (bre methodNotAllowedError) Error() string {
	return http.StatusText(bre.HttpStatus())
}

func (bre methodNotAllowedError) Details() any {
	return bre.details
}

func (bre methodNotAllowedError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "MethodNotAllowedError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", bre.details)},
	}
}

func MethodNotAllowedError(details any) AppError {
	return methodNotAllowedError{details}
}
