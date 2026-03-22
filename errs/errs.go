package errs

import (
	"fmt"
)

type CustomError struct {
	Code    ErrorCode
	Message string
	err     error
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}

func New(code ErrorCode, msg string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: msg,
		err:     err,
	}
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
