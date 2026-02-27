package ungerr

import (
	"fmt"
	"runtime"

	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type UnknownError struct {
	msg  string
	file string
	line int
	fn   string
	err  error
}

func (e *UnknownError) Error() string {
	if e.err == nil {
		return fmt.Sprintf("%s\n\tat %s (%s:%d)", e.msg, e.fn, e.file, e.line)
	}

	return fmt.Sprintf("wrapped error: %s\n\tat %s (%s:%d)\n%s", e.msg, e.fn, e.file, e.line, e.err.Error())
}

func (e *UnknownError) ToLogAttrs() []LogAttr {
	errType := "UnknownError"
	if e.err != nil {
		errType = fmt.Sprintf("%T", e.err)
	}

	attrs := []LogAttr{
		{Key: string(semconv.ErrorMessageKey), Value: e.msg},
		{Key: string(semconv.ErrorTypeKey), Value: errType},
		{Key: string(semconv.CodeFilePathKey), Value: e.file},
		{Key: string(semconv.CodeLineNumberKey), Value: e.line},
		{Key: string(semconv.CodeFunctionNameKey), Value: e.fn},
	}

	if e.err != nil {
		attrs = append(attrs, LogAttr{Key: "error.cause", Value: e.err.Error()})
	}

	return attrs
}

func extractCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", 0, "unknown"
	}

	fn := runtime.FuncForPC(pc)
	fnName := "unknown"
	if fn != nil {
		fnName = fn.Name()
	}

	return file, line, fnName
}

func Unknown(msg string) *UnknownError {
	file, line, fnName := extractCallerInfo()

	return &UnknownError{
		msg:  msg,
		file: file,
		line: line,
		fn:   fnName,
	}
}

func Unknownf(format string, args ...any) *UnknownError {
	file, line, fnName := extractCallerInfo()

	return &UnknownError{
		msg:  fmt.Sprintf(format, args...),
		file: file,
		line: line,
		fn:   fnName,
	}
}

func Wrap(err error, msg string) *UnknownError {
	file, line, fnName := extractCallerInfo()

	return &UnknownError{
		msg:  msg,
		file: file,
		line: line,
		fn:   fnName,
		err:  err,
	}
}

func Wrapf(err error, format string, args ...any) *UnknownError {
	file, line, fnName := extractCallerInfo()

	return &UnknownError{
		msg:  fmt.Sprintf(format, args...),
		file: file,
		line: line,
		fn:   fnName,
		err:  err,
	}
}

func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if unknownErr, ok := err.(*UnknownError); ok {
		return unknownErr.err
	}

	return err
}
