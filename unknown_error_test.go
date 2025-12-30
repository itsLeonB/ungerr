package ungerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknown(t *testing.T) {
	msg := "something went wrong"
	err := Unknown(msg)

	assert.Equal(t, msg, err.msg)
	assert.NotEmpty(t, err.file, "expected file to be set")
	assert.NotZero(t, err.line, "expected line to be set")
	assert.Contains(t, err.fn, "TestUnknown", "expected fn to contain 'TestUnknown'")
}

func TestUnknownErrorError(t *testing.T) {
	err := &UnknownError{
		msg:  "error message",
		file: "/path/to/file.go",
		line: 123,
		fn:   "pkg.Func",
	}

	expected := "error message\n\tat pkg.Func (/path/to/file.go:123)"
	assert.Equal(t, expected, err.Error())
}
