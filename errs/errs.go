package errs

import (
	"fmt"
)

type CustomError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (c *CustomError) Error() string {
	if c.Err != nil {
		return fmt.Sprintf("code: %d, message: %s, err: %v", c.Code, c.Message, c.Err)
	}
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}

func New(code ErrorCode, msg string, err ...error) *CustomError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}

	return &CustomError{
		Code:    code,
		Message: msg,
		Err:     e,
	}
}

func (c *CustomError) Unwrap() error {
	return c.Err
}

func (c *CustomError) Is(target error) bool {
	return c == target
}

type ErrorCode int

const (
	OK = iota
	CANCELLED
	UNKNOWN
	INVALID_ARGUMENT
	DEADLINE_EXCEEDED
	NOT_FOUND
	ALREADY_EXISTS
	PERMISSION_DENIED
	RESOURCE_EXHAUSTED
	FAILED_PRECONDITION
	ABORTED
	OUT_OF_RANGE
	UNIMPLEMENTED
	INTERNAL
	UNAVAILABLE
	DATA_LOSS
	UNAUTHENTICATED
)
