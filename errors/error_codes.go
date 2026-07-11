package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorCode string

const (
	ErrorCodeNotFound      ErrorCode = "NOT_FOUND"
	ErrorCodeInvalidArg    ErrorCode = "INVALID_ARGUMENT"
	ErrorCodeInternal      ErrorCode = "INTERNAL"
	ErrorCodeAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrorCodeForbidden     ErrorCode = "FORBIDDEN"
	ErrorCodeUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrorCodeUnknown       ErrorCode = "UNKNOWN"
)

// CodeToGRPCCode маппит ErrorCode в gRPC код
func CodeToGRPCCode(code ErrorCode) codes.Code {
	switch code {
	case ErrorCodeNotFound:
		return codes.NotFound
	case ErrorCodeInvalidArg:
		return codes.InvalidArgument
	case ErrorCodeInternal:
		return codes.Internal
	case ErrorCodeAlreadyExists:
		return codes.AlreadyExists
	case ErrorCodeForbidden:
		return codes.PermissionDenied
	case ErrorCodeUnauthorized:
		return codes.Unauthenticated
	default:
		return codes.Unknown
	}
}

func GRPCCodeToErrorCode(code codes.Code) ErrorCode {
	switch code {
	case codes.NotFound:
		return ErrorCodeNotFound
	case codes.InvalidArgument:
		return ErrorCodeInvalidArg
	case codes.Internal:
		return ErrorCodeInternal
	case codes.AlreadyExists:
		return ErrorCodeAlreadyExists
	case codes.PermissionDenied:
		return ErrorCodeForbidden
	case codes.Unauthenticated:
		return ErrorCodeUnauthorized
	default:
		return ErrorCodeUnknown
	}
}

func StatusFromErrorCode(code ErrorCode, message string) *status.Status {
	return status.New(CodeToGRPCCode(code), message)
}
