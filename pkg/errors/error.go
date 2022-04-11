package errors

import (
	"errors"
	"fmt"
)

type ComplexError struct {
	Err   error
	Cause error
}

func (e ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.Err.Error(), e.Cause.Error())
}

func Error(msg string) error {
	return errors.New(msg)
}

func Errorf(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

func WrapError(cause error, msg string) error {
	return WrapComplexError(cause, Error(msg))
}

func WrapErrorf(cause error, msg string, args ...interface{}) error {
	return WrapComplexError(cause, Errorf(msg, args...))
}

func WrapComplexError(cause, err error) error {
	if cause == nil {
		cause = Error("<nil cause>")
	}

	return ComplexError{
		Err:   err,
		Cause: cause,
	}
}
