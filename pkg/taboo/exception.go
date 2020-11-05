package taboo

import (
	"errors"
	"fmt"
	"strings"
)

type Exception struct {
	message error
	caller  *caller
	cause   *Exception
}

func fromThrow(err error) *Exception {
	return &Exception{
		message: err,
		caller:  call(3),
		cause:   nil,
	}
}

func fromError(err error) *Exception {
	return &Exception{
		message: err,
		caller:  call(5),
		cause:   nil,
	}
}

func fromInterface(err interface{}) *Exception {
	return &Exception{
		message: fmt.Errorf("%v", err),
		caller:  call(4),
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
		message: errors.New(message),
		caller:  call(2),
		cause:   e,
	}

	panic(ex)
}

func (e *Exception) ThrowErr(err error) {
	ex := &Exception{
		message: err,
		caller:  call(2),
		cause:   e,
	}

	panic(ex)
}

func (e *Exception) Has(err error) bool {
	if e == nil || err == nil || e.message == nil {
		return false
	}

	if errors.Is(e.message, err) {
		return true
	}

	return e.cause.Has(err)
}