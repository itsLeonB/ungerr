package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestForbiddenError(t *testing.T) {
	details := map[string]string{"reason": "insufficient permissions"}
	err := ungerr.ForbiddenError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(7), err.GrpcStatus())
	assert.Equal(t, http.StatusForbidden, err.HttpStatus())
	assert.Equal(t, "Forbidden", err.Error())
	assert.Equal(t, details, err.Details())
}

func TestForbiddenErrorWithNilDetails(t *testing.T) {
	err := ungerr.ForbiddenError(nil)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(7), err.GrpcStatus())
	assert.Equal(t, http.StatusForbidden, err.HttpStatus())
	assert.Equal(t, "Forbidden", err.Error())
	assert.Nil(t, err.Details())
}

func TestForbiddenErrorWithStringDetails(t *testing.T) {
	details := "user does not have admin privileges"
	err := ungerr.ForbiddenError(details)

	assert.NotNil(t, err)
	assert.Equal(t, uint32(7), err.GrpcStatus())
	assert.Equal(t, http.StatusForbidden, err.HttpStatus())
	assert.Equal(t, "Forbidden", err.Error())
	assert.Equal(t, details, err.Details())
}
