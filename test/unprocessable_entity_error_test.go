package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestUnprocessableEntityError(t *testing.T) {
	details := map[string]interface{}{
		"errors": []string{"age must be positive", "email already exists"},
	}
	err := ungerr.UnprocessableEntityError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestUnprocessableEntityErrorWithNilDetails(t *testing.T) {
	err := ungerr.UnprocessableEntityError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Nil(t, err.Details())
}

func TestUnprocessableEntityErrorWithStringDetails(t *testing.T) {
	details := "semantic validation failed"
	err := ungerr.UnprocessableEntityError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusUnprocessableEntity, err.HttpStatus())
	assert.Equal(t, "Unprocessable Entity", err.Error())
	assert.Equal(t, details, err.Details())
}
