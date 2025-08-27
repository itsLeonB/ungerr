package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundError(t *testing.T) {
	details := map[string]interface{}{"resource": "user", "id": 123}
	err := ungerr.NotFoundError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(5), err.GrpcStatus())
	assert.Equal(t, http.StatusNotFound, err.HttpStatus())
	assert.Equal(t, "Not Found", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestNotFoundErrorWithNilDetails(t *testing.T) {
	err := ungerr.NotFoundError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(5), err.GrpcStatus())
	assert.Equal(t, http.StatusNotFound, err.HttpStatus())
	assert.Equal(t, "Not Found", err.Error())
	assert.Nil(t, err.Details())
}

func TestNotFoundErrorWithStringDetails(t *testing.T) {
	details := "user with ID 123 not found"
	err := ungerr.NotFoundError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(5), err.GrpcStatus())
	assert.Equal(t, http.StatusNotFound, err.HttpStatus())
	assert.Equal(t, "Not Found", err.Error())
	assert.Equal(t, details, err.Details())
}
