package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	details := map[string]interface{}{
		"field_errors": map[string]string{
			"email":    "must be a valid email address",
			"password": "must be at least 8 characters long",
		},
	}
	err := ungerr.ValidationError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestValidationErrorWithNilDetails(t *testing.T) {
	err := ungerr.ValidationError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Nil(t, err.Details())
}

func TestValidationErrorWithStringDetails(t *testing.T) {
	details := "validation failed for multiple fields"
	err := ungerr.ValidationError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestValidationErrorWithSliceDetails(t *testing.T) {
	details := []map[string]string{
		{"field": "name", "error": "required"},
		{"field": "age", "error": "must be positive"},
	}
	err := ungerr.ValidationError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Equal(t, details, err.Details())
}
