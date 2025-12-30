package ungerr

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppErrors(t *testing.T) {
	tests := []struct {
		name        string
		err         AppError
		wantHTTP    int
		wantGRPC    uint32
		wantDetails any
	}{
		{
			name:        "BadRequestError",
			err:         BadRequestError("bad request details"),
			wantHTTP:    http.StatusBadRequest,
			wantGRPC:    3,
			wantDetails: "bad request details",
		},
		{
			name:        "ConflictError",
			err:         ConflictError("conflict details"),
			wantHTTP:    http.StatusConflict,
			wantGRPC:    6,
			wantDetails: "conflict details",
		},
		{
			name:        "ForbiddenError",
			err:         ForbiddenError("forbidden details"),
			wantHTTP:    http.StatusForbidden,
			wantGRPC:    7,
			wantDetails: "forbidden details",
		},
		{
			name:        "InternalServerError",
			err:         InternalServerError(),
			wantHTTP:    http.StatusInternalServerError,
			wantGRPC:    13,
			wantDetails: nil,
		},
		{
			name:        "NotFoundError",
			err:         NotFoundError("not found details"),
			wantHTTP:    http.StatusNotFound,
			wantGRPC:    5,
			wantDetails: "not found details",
		},
		{
			name:        "UnauthorizedError",
			err:         UnauthorizedError("unauthorized details"),
			wantHTTP:    http.StatusUnauthorized,
			wantGRPC:    16,
			wantDetails: "unauthorized details",
		},
		{
			name:        "UnprocessableEntityError",
			err:         UnprocessableEntityError("unprocessable details"),
			wantHTTP:    http.StatusUnprocessableEntity,
			wantGRPC:    3,
			wantDetails: "unprocessable details",
		},
		{
			name:        "ValidationError",
			err:         ValidationError("validation details"),
			wantHTTP:    http.StatusUnprocessableEntity,
			wantGRPC:    3,
			wantDetails: "validation details",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantHTTP, tt.err.HttpStatus())
			assert.Equal(t, tt.wantGRPC, tt.err.GrpcStatus())
			assert.Equal(t, http.StatusText(tt.wantHTTP), tt.err.Error())
			assert.Equal(t, tt.wantDetails, tt.err.Details())
		})
	}
}
