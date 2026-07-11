package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCError struct {
	Code    codes.Code
	Message string
	Details map[string]string
}

func (e *GRPCError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code.String(), e.Message)
}

func NewGRPCError(code codes.Code, message string) *GRPCError {
	return &GRPCError{
		Code:    code,
		Message: message,
	}
}

func NewGRPCErrorWithDetails(code codes.Code, message string, details map[string]string) *GRPCError {
	return &GRPCError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func (e *GRPCError) ToStatusError() error {
	return status.Error(e.Code, e.Message)
}

func FromError(err error) *GRPCError {
	if err == nil {
		return nil
	}

	if grpcErr, ok := err.(*GRPCError); ok {
		return grpcErr
	}

	st, ok := status.FromError(err)
	if !ok {
		return &GRPCError{
			Code:    codes.Unknown,
			Message: err.Error(),
		}
	}

	return &GRPCError{
		Code:    st.Code(),
		Message: st.Message(),
	}
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	return ok && st.Code() == codes.NotFound
}

func IsInvalidArgument(err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	return ok && st.Code() == codes.InvalidArgument
}

func IsInternal(err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	return ok && st.Code() == codes.Internal
}
