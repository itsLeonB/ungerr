package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type timeoutError struct {
	details any
}

func (ue timeoutError) GrpcStatus() uint32 {
	return 4
}

func (ue timeoutError) HttpStatus() int {
	return http.StatusRequestTimeout
}

func (ue timeoutError) Error() string {
	return http.StatusText(ue.HttpStatus())
}

func (ue timeoutError) Details() any {
	return ue.details
}

func (ue timeoutError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "TimeoutError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", ue.details)},
	}
}

func TimeoutError(details any) AppError {
	return timeoutError{details}
}
