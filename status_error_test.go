package ungerr

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppErrorGetStatusMatchesHTTPStatus(t *testing.T) {
	errs := []AppError{
		BadRequestError("x"),
		ConflictError("x"),
		ForbiddenError("x"),
		InternalServerError(),
		MethodNotAllowedError("x"),
		NotFoundError("x"),
		TimeoutError("x"),
		TooManyRequestsError("x"),
		UnauthorizedError("x"),
		UnprocessableEntityError("x"),
		ValidationError("x"),
	}

	for _, err := range errs {
		se, ok := err.(interface{ GetStatus() int })
		assert.True(t, ok, "%T should implement GetStatus()", err)
		assert.Equal(t, err.HttpStatus(), se.GetStatus())
	}
}

func TestAppErrorMarshalJSON(t *testing.T) {
	b, err := json.Marshal(NotFoundError("transfer method is not found"))
	assert.NoError(t, err)

	var body map[string]any
	assert.NoError(t, json.Unmarshal(b, &body))
	assert.Equal(t, "Not Found", body["title"])
	assert.Equal(t, float64(http.StatusNotFound), body["status"])
	assert.Equal(t, "transfer method is not found", body["detail"])
}

func TestInternalServerErrorMarshalJSONOmitsNilDetail(t *testing.T) {
	b, err := json.Marshal(InternalServerError())
	assert.NoError(t, err)

	var body map[string]any
	assert.NoError(t, json.Unmarshal(b, &body))
	assert.Equal(t, "Internal Server Error", body["title"])
	assert.Equal(t, float64(http.StatusInternalServerError), body["status"])
	_, hasDetail := body["detail"]
	assert.False(t, hasDetail, "nil details should be omitted")
}

func TestUnknownErrorIsStatusErrorAndDoesNotLeak(t *testing.T) {
	ue := Wrap(assertSentinel{}, "loading profile from db")

	// Satisfies huma.StatusError (GetStatus + Error).
	assert.Equal(t, http.StatusInternalServerError, ue.GetStatus())

	b, err := json.Marshal(ue)
	assert.NoError(t, err)

	// The serialized body must not expose the internal message, wrapped cause,
	// or any call-site details.
	s := string(b)
	assert.NotContains(t, s, "loading profile from db")
	assert.NotContains(t, s, "status_error_test.go")
	assert.NotContains(t, s, "sentinel-cause")

	var body map[string]any
	assert.NoError(t, json.Unmarshal(b, &body))
	assert.Equal(t, "Internal Server Error", body["title"])
	assert.Equal(t, float64(http.StatusInternalServerError), body["status"])
	assert.Equal(t, "Internal Server Error", body["detail"])
}

type assertSentinel struct{}

func (assertSentinel) Error() string { return "sentinel-cause" }
