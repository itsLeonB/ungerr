package ungerr

import (
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

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

func (bre internalServerError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "InternalServerError"},
	}
}

func InternalServerError() AppError {
	return internalServerError{}
}
