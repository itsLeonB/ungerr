package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type notFoundError struct {
	details any
}

func (nfe notFoundError) GrpcStatus() uint32 {
	return 5
}

func (nfe notFoundError) HttpStatus() int {
	return http.StatusNotFound
}

func (nfe notFoundError) Error() string {
	return http.StatusText(nfe.HttpStatus())
}

func (nfe notFoundError) Details() any {
	return nfe.details
}

func (nfe notFoundError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "NotFoundError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", nfe.details)},
	}
}

func NotFoundError(details any) AppError {
	return notFoundError{details}
}
