package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (ce conflictError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "ConflictError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", ce.details)},
	}
}

func ConflictError(details any) AppError {
	return conflictError{details}
}
