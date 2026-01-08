package ungerr

import (
	"fmt"
	"runtime"
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

func Unknown(msg string) *UnknownError {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return &UnknownError{
			msg:  msg,
			file: "unknown",
			line: 0,
			fn:   "unknown",
		}
	}

	fn := runtime.FuncForPC(pc)
	fnName := "unknown"
	if fn != nil {
		fnName = fn.Name()
	}

	return &UnknownError{
		msg:  msg,
		file: file,
		line: line,
		fn:   fnName,
	}
}

func Unknownf(format string, args ...any) *UnknownError {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return &UnknownError{
		msg:  fmt.Sprintf(format, args...),
		file: file,
		line: line,
		fn:   fn.Name(),
	}
}

func Wrap(err error, msg string) *UnknownError {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return &UnknownError{
		msg:  msg,
		file: file,
		line: line,
		fn:   fn.Name(),
		err:  err,
	}
}

func Wrapf(err error, format string, args ...any) *UnknownError {
	return Wrap(err, fmt.Sprintf(format, args...))
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
