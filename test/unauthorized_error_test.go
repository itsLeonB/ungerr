package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestUnauthorizedError(t *testing.T) {
	details := map[string]string{"reason": "invalid token"}
	err := ungerr.UnauthorizedError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(16), err.GrpcStatus())
	assert.Equal(t, http.StatusUnauthorized, err.HttpStatus())
	assert.Equal(t, "Unauthorized", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestUnauthorizedErrorWithNilDetails(t *testing.T) {
	err := ungerr.UnauthorizedError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(16), err.GrpcStatus())
	assert.Equal(t, http.StatusUnauthorized, err.HttpStatus())
	assert.Equal(t, "Unauthorized", err.Error())
	assert.Nil(t, err.Details())
}

func TestUnauthorizedErrorWithStringDetails(t *testing.T) {
	details := "missing authentication token"
	err := ungerr.UnauthorizedError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(16), err.GrpcStatus())
	assert.Equal(t, http.StatusUnauthorized, err.HttpStatus())
	assert.Equal(t, "Unauthorized", err.Error())
	assert.Equal(t, details, err.Details())
}
