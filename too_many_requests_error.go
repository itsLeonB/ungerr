package ungerr

import (
	"fmt"
	"net/http"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type tooManyRequestsError struct {
	details any
}

func (ve tooManyRequestsError) GrpcStatus() uint32 {
	return 8
}

func (ve tooManyRequestsError) HttpStatus() int {
	return http.StatusTooManyRequests
}

func (ve tooManyRequestsError) Error() string {
	return http.StatusText(ve.HttpStatus())
}

func (ve tooManyRequestsError) Details() any {
	return ve.details
}

func (ve tooManyRequestsError) ToLogAttrs() []LogAttr {
	return []LogAttr{
		{Key: string(semconv.ErrorTypeKey), Value: "TooManyRequestsError"},
		{Key: string(semconv.ErrorMessageKey), Value: fmt.Sprintf("%v", ve.details)},
	}
}

func TooManyRequestsError(details any) AppError {
	return tooManyRequestsError{details}
}
