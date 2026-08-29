package ungerr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// errorBody is the JSON representation an ungerr error marshals to when it is
// written directly as an HTTP response body. It mirrors the OpenAPI
// problem/ErrorModel shape ({title, status, detail}) used by Huma, so that
// errors returned from Huma handlers render identically to Huma's own built-in
// (e.g. request-validation) errors.
type errorBody struct {
	Title  string `json:"title,omitempty"`
	Status int    `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// detailString renders an error's details as a human-readable string for the
// ErrorModel "detail" field.
func detailString(details any) string {
	switch d := details.(type) {
	case nil:
		return ""
	case string:
		return d
	case error:
		return d.Error()
	case fmt.Stringer:
		return d.String()
	default:
		return fmt.Sprintf("%v", d)
	}
}

// marshalAppError serializes any AppError into the shared errorBody shape.
func marshalAppError(e AppError) ([]byte, error) {
	return json.Marshal(errorBody{
		Title:  e.Error(),
		Status: e.HttpStatus(),
		Detail: detailString(e.Details()),
	})
}

// The methods below add GetStatus() to every concrete error type so that they
// satisfy Huma's StatusError interface (GetStatus() int + Error() string),
// letting a handler return a raw ungerr error and have Huma respond with the
// correct HTTP status. MarshalJSON renders the shared errorBody shape.

func (bre badRequestError) GetStatus() int               { return bre.HttpStatus() }
func (bre badRequestError) MarshalJSON() ([]byte, error) { return marshalAppError(bre) }

func (ce conflictError) GetStatus() int               { return ce.HttpStatus() }
func (ce conflictError) MarshalJSON() ([]byte, error) { return marshalAppError(ce) }

func (fe forbiddenError) GetStatus() int               { return fe.HttpStatus() }
func (fe forbiddenError) MarshalJSON() ([]byte, error) { return marshalAppError(fe) }

func (ise internalServerError) GetStatus() int               { return ise.HttpStatus() }
func (ise internalServerError) MarshalJSON() ([]byte, error) { return marshalAppError(ise) }

func (bre methodNotAllowedError) GetStatus() int               { return bre.HttpStatus() }
func (bre methodNotAllowedError) MarshalJSON() ([]byte, error) { return marshalAppError(bre) }

func (nfe notFoundError) GetStatus() int               { return nfe.HttpStatus() }
func (nfe notFoundError) MarshalJSON() ([]byte, error) { return marshalAppError(nfe) }

func (ue timeoutError) GetStatus() int               { return ue.HttpStatus() }
func (ue timeoutError) MarshalJSON() ([]byte, error) { return marshalAppError(ue) }

func (ve tooManyRequestsError) GetStatus() int               { return ve.HttpStatus() }
func (ve tooManyRequestsError) MarshalJSON() ([]byte, error) { return marshalAppError(ve) }

func (ue unauthorizedError) GetStatus() int               { return ue.HttpStatus() }
func (ue unauthorizedError) MarshalJSON() ([]byte, error) { return marshalAppError(ue) }

func (uee unprocessableEntityError) GetStatus() int               { return uee.HttpStatus() }
func (uee unprocessableEntityError) MarshalJSON() ([]byte, error) { return marshalAppError(uee) }

func (ve validationError) GetStatus() int               { return ve.HttpStatus() }
func (ve validationError) MarshalJSON() ([]byte, error) { return marshalAppError(ve) }

// GetStatus makes *UnknownError a Huma StatusError (HTTP 500). Because Huma
// writes the returned error as the response body, this is required to avoid
// leaking UnknownError's internal message — which includes the message, wrapped
// cause, file, line, and function of the call site — to the client. MarshalJSON
// therefore renders only a generic 500 body and never exposes those internals.
func (e *UnknownError) GetStatus() int { return http.StatusInternalServerError }

func (e *UnknownError) MarshalJSON() ([]byte, error) {
	return json.Marshal(errorBody{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: http.StatusInternalServerError,
		Detail: http.StatusText(http.StatusInternalServerError),
	})
}
