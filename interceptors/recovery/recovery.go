package recovery

import (
	"context"
	"runtime/debug"

	log "github.com/erdedan1/shared/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryServerInterceptor(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		defer func() {
			if r := recover(); r != nil {
				logger.Error(
					"Interceptors",
					"RecoveryServerInterceptor",
					"panic recovered in gRPC handler",
					err,
					"stacktrace", string(debug.Stack()),
				)

				err = status.Errorf(codes.Internal, "panic: %v", r)
			}
		}()

		return handler(ctx, req)
	}
}
