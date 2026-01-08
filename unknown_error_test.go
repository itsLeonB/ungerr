package ungerr

import (
	"errors"
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

func TestUnknownf(t *testing.T) {
	format := "error %d: %s"
	err := Unknownf(format, 404, "not found")

	assert.Equal(t, "error 404: not found", err.msg)
	assert.NotEmpty(t, err.file)
	assert.NotZero(t, err.line)
	assert.Contains(t, err.fn, "TestUnknownf")
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

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	err := Wrap(originalErr, "wrapped message")

	assert.Equal(t, "wrapped message", err.msg)
	assert.Equal(t, originalErr, err.err)
	assert.Contains(t, err.fn, "TestWrap")
}

func TestWrapf(t *testing.T) {
	originalErr := errors.New("original error")
	err := Wrapf(originalErr, "wrapped %s: %d", "message", 500)

	assert.Equal(t, "wrapped message: 500", err.msg)
	assert.Equal(t, originalErr, err.err)
	assert.Contains(t, err.fn, "TestWrapf")
}

func TestUnwrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "wrapped")

	assert.Equal(t, originalErr, Unwrap(wrappedErr))
	assert.Nil(t, Unwrap(nil))
	assert.Equal(t, originalErr, Unwrap(originalErr))
}

func TestUnknownErrorErrorWithWrappedError(t *testing.T) {
	originalErr := errors.New("original error")
	err := &UnknownError{
		msg:  "wrapped message",
		file: "/path/to/file.go",
		line: 123,
		fn:   "pkg.Func",
		err:  originalErr,
	}

	result := err.Error()
	assert.Contains(t, result, "wrapped error: wrapped message")
	assert.Contains(t, result, "at pkg.Func (/path/to/file.go:123)")
	assert.Contains(t, result, "original error")
}
