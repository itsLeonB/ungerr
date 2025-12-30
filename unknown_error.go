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
}

func (e *UnknownError) Error() string {
	return fmt.Sprintf("%s\n\tat %s (%s:%d)", e.msg, e.fn, e.file, e.line)
}

func Unknown(msg string) *UnknownError {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return &UnknownError{
		msg:  msg,
		file: file,
		line: line,
		fn:   fn.Name(),
	}
}
