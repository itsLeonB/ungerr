package test

import (
	"net/http"
	"testing"

	"github.com/itsLeonB/ungerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AppErrorSuite struct {
	suite.Suite
}

func (suite *AppErrorSuite) TestAllErrorsImplementAppErrorInterface() {
	errors := []ungerr.AppError{
		ungerr.InternalServerError(),
		ungerr.ConflictError("test"),
		ungerr.NotFoundError("test"),
		ungerr.UnauthorizedError("test"),
		ungerr.ForbiddenError("test"),
		ungerr.BadRequestError("test"),
		ungerr.UnprocessableEntityError("test"),
		ungerr.ValidationError("test"),
	}

	for _, err := range errors {
		assert.NotNil(suite.T(), err)
		assert.NotNil(suite.T(), err.Error())
		assert.NotZero(suite.T(), err.HttpStatus())
		assert.NotZero(suite.T(), err.GrpcStatus())
		// Note: Details can be nil for some errors like InternalServerError
	}
}

func (suite *AppErrorSuite) TestAllErrorsImplementStandardErrorInterface() {
	errors := []error{
		ungerr.InternalServerError(),
		ungerr.ConflictError("test"),
		ungerr.NotFoundError("test"),
		ungerr.UnauthorizedError("test"),
		ungerr.ForbiddenError("test"),
		ungerr.BadRequestError("test"),
		ungerr.UnprocessableEntityError("test"),
		ungerr.ValidationError("test"),
	}

	for _, err := range errors {
		assert.NotNil(suite.T(), err)
		assert.NotEmpty(suite.T(), err.Error())
	}
}

func (suite *AppErrorSuite) TestHttpStatusCodes() {
	testCases := []struct {
		name           string
		error          ungerr.AppError
		expectedStatus int
	}{
		{"InternalServerError", ungerr.InternalServerError(), http.StatusInternalServerError},
		{"ConflictError", ungerr.ConflictError("test"), http.StatusConflict},
		{"NotFoundError", ungerr.NotFoundError("test"), http.StatusNotFound},
		{"UnauthorizedError", ungerr.UnauthorizedError("test"), http.StatusUnauthorized},
		{"ForbiddenError", ungerr.ForbiddenError("test"), http.StatusForbidden},
		{"BadRequestError", ungerr.BadRequestError("test"), http.StatusBadRequest},
		{"UnprocessableEntityError", ungerr.UnprocessableEntityError("test"), http.StatusUnprocessableEntity},
		{"ValidationError", ungerr.ValidationError("test"), http.StatusUnprocessableEntity},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStatus, tc.error.HttpStatus())
		})
	}
}

func (suite *AppErrorSuite) TestGrpcStatusCodes() {
	testCases := []struct {
		name           string
		error          ungerr.AppError
		expectedStatus uint32
	}{
		{"InternalServerError", ungerr.InternalServerError(), uint32(13)},                // INTERNAL
		{"ConflictError", ungerr.ConflictError("test"), uint32(6)},                       // ALREADY_EXISTS
		{"NotFoundError", ungerr.NotFoundError("test"), uint32(5)},                       // NOT_FOUND
		{"UnauthorizedError", ungerr.UnauthorizedError("test"), uint32(16)},              // UNAUTHENTICATED
		{"ForbiddenError", ungerr.ForbiddenError("test"), uint32(7)},                     // PERMISSION_DENIED
		{"BadRequestError", ungerr.BadRequestError("test"), uint32(3)},                   // INVALID_ARGUMENT
		{"UnprocessableEntityError", ungerr.UnprocessableEntityError("test"), uint32(3)}, // INVALID_ARGUMENT
		{"ValidationError", ungerr.ValidationError("test"), uint32(3)},                   // INVALID_ARGUMENT
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStatus, tc.error.GrpcStatus())
		})
	}
}

func (suite *AppErrorSuite) TestErrorMessages() {
	testCases := []struct {
		name            string
		error           ungerr.AppError
		expectedMessage string
	}{
		{"InternalServerError", ungerr.InternalServerError(), "Internal Server Error"},
		{"ConflictError", ungerr.ConflictError("test"), "Conflict"},
		{"NotFoundError", ungerr.NotFoundError("test"), "Not Found"},
		{"UnauthorizedError", ungerr.UnauthorizedError("test"), "Unauthorized"},
		{"ForbiddenError", ungerr.ForbiddenError("test"), "Forbidden"},
		{"BadRequestError", ungerr.BadRequestError("test"), "Bad Request"},
		{"UnprocessableEntityError", ungerr.UnprocessableEntityError("test"), "Unprocessable Entity"},
		{"ValidationError", ungerr.ValidationError("test"), "Unprocessable Entity"},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedMessage, tc.error.Error())
		})
	}
}

func (suite *AppErrorSuite) TestDetailsPreservation() {
	testDetails := map[string]interface{}{
		"field":   "email",
		"value":   "invalid-email",
		"message": "must be a valid email address",
	}

	// Test errors that support details
	errors := []ungerr.AppError{
		ungerr.ConflictError(testDetails),
		ungerr.NotFoundError(testDetails),
		ungerr.UnauthorizedError(testDetails),
		ungerr.ForbiddenError(testDetails),
		ungerr.BadRequestError(testDetails),
		ungerr.UnprocessableEntityError(testDetails),
		ungerr.ValidationError(testDetails),
	}

	for _, err := range errors {
		assert.Equal(suite.T(), testDetails, err.Details())
	}
}

func (suite *AppErrorSuite) TestInternalServerErrorHasNilDetails() {
	err := ungerr.InternalServerError()
	assert.Nil(suite.T(), err.Details())
}

func (suite *AppErrorSuite) TestInternalServerErrorWithoutDetails() {
	err := ungerr.InternalServerError()

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), uint32(13), err.GrpcStatus())
	assert.Equal(suite.T(), http.StatusInternalServerError, err.HttpStatus())
	assert.Equal(suite.T(), "Internal Server Error", err.Error())
	assert.Nil(suite.T(), err.Details())
}

func TestAppErrorSuite(t *testing.T) {
	suite.Run(t, new(AppErrorSuite))
}
