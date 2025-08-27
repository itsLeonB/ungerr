package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestBadRequestError(t *testing.T) {
	details := map[string]string{"field": "email", "error": "invalid format"}
	err := ungerr.BadRequestError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusBadRequest, err.HttpStatus())
	assert.Equal(t, "Bad Request", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestBadRequestErrorWithNilDetails(t *testing.T) {
	err := ungerr.BadRequestError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusBadRequest, err.HttpStatus())
	assert.Equal(t, "Bad Request", err.Error())
	assert.Nil(t, err.Details())
}

func TestBadRequestErrorWithStringDetails(t *testing.T) {
	details := "malformed JSON in request body"
	err := ungerr.BadRequestError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusBadRequest, err.HttpStatus())
	assert.Equal(t, "Bad Request", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestBadRequestErrorWithSliceDetails(t *testing.T) {
	details := []string{"missing required field: name", "invalid email format"}
	err := ungerr.BadRequestError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(3), err.GrpcStatus())
	assert.Equal(t, http.StatusBadRequest, err.HttpStatus())
	assert.Equal(t, "Bad Request", err.Error())
	assert.Equal(t, details, err.Details())
}
