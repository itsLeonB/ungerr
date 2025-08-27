package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {
	err := ungerr.InternalServerError()

	assert.NotNil(t, err)
	assert.Equal(t, uint32(13), err.GrpcStatus())
	assert.Equal(t, http.StatusInternalServerError, err.HttpStatus())
	assert.Equal(t, "Internal Server Error", err.Error())
	assert.Nil(t, err.Details())
}

func TestInternalServerErrorAlwaysReturnsNilDetails(t *testing.T) {
	err := ungerr.InternalServerError()

	// Internal server error should always return nil details
	assert.Nil(t, err.Details())
}

func TestInternalServerErrorConsistentBehavior(t *testing.T) {
	err1 := ungerr.InternalServerError()
	err2 := ungerr.InternalServerError()

	// Both instances should behave identically
	assert.Equal(t, err1.GrpcStatus(), err2.GrpcStatus())
	assert.Equal(t, err1.HttpStatus(), err2.HttpStatus())
	assert.Equal(t, err1.Error(), err2.Error())
	assert.Equal(t, err1.Details(), err2.Details())
}
