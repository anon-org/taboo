package taboo

import (
	"fmt"
	"strings"
)

type Exception struct {
	message string
	caller  *caller
	cause   *Exception
}

func fromThrow(err error) *Exception {
	return &Exception{
		message: err.Error(),
		caller:  call(3),
		cause:   nil,
	}
}

func fromPanic(err error) *Exception {
	return &Exception{
		message: err.Error(),
		caller:  call(5),
		cause:   nil,
	}
}

func (e *Exception) extract() string {
	return fmt.Sprintf("%v:%v %v", e.caller.function, e.caller.Line(), e.message)
}

func (e *Exception) Error() string {
	if e.caller == nil {
		return "<nil>"
	}

	var b strings.Builder

	b.WriteString(e.extract())

	cause := e.cause

	depth := "  "

	for cause != nil {
		b.WriteString(" caused by:\n" + depth)
		b.WriteString(cause.extract())

		cause = cause.cause

		depth += "  "
	}

	return b.String()
}

func (e *Exception) Throw(message string) {
	ex := &Exception{
		message: message,
		caller:  call(2),
		cause:   e,
	}

	panic(ex)
}
