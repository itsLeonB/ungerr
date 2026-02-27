package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (ue unauthorizedError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "UnauthorizedError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", ue.details)},
	}
}

func UnauthorizedError(details any) AppError {
	return unauthorizedError{details}
}
