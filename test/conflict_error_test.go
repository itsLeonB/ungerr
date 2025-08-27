package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestConflictError(t *testing.T) {
	details := map[string]string{"resource": "user already exists"}
	err := ungerr.ConflictError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(6), err.GrpcStatus())
	assert.Equal(t, http.StatusConflict, err.HttpStatus())
	assert.Equal(t, "Conflict", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestConflictErrorWithNilDetails(t *testing.T) {
	err := ungerr.ConflictError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(6), err.GrpcStatus())
	assert.Equal(t, http.StatusConflict, err.HttpStatus())
	assert.Equal(t, "Conflict", err.Error())
	assert.Nil(t, err.Details())
}

func TestConflictErrorWithStringDetails(t *testing.T) {
	details := "duplicate email address"
	err := ungerr.ConflictError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(6), err.GrpcStatus())
	assert.Equal(t, http.StatusConflict, err.HttpStatus())
	assert.Equal(t, "Conflict", err.Error())
	assert.Equal(t, details, err.Details())
}
