package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (fe forbiddenError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "ForbiddenError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", fe.details)},
	}
}

func ForbiddenError(details any) AppError {
	return forbiddenError{details}
}
