package unary

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		requestID := GetRequestID(ctx)

		logger.Info("gRPC request started",
			zap.String("x-request-id", requestID),
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
		)

		resp, err := handler(ctx, req)

		duration := time.Since(start)
		st, _ := status.FromError(err)

		logLevel := zap.InfoLevel
		if err != nil {
			logLevel = zap.ErrorLevel
		}

		logger.Log(logLevel, "gRPC request completed",
			zap.String("x-request-id", requestID),
			zap.String("method", info.FullMethod),
			zap.String("status", st.Code().String()),
			zap.Duration("duration", duration),
			zap.Error(err),
		)

		return resp, err
	}
}

func ClientLoggerInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		requestID := GetRequestID(ctx)

		logger.Info("gRPC client request started",
			zap.String("x-request-id", requestID),
			zap.String("method", method),
		)

		err := invoker(ctx, method, req, reply, cc, opts...)

		duration := time.Since(start)
		st, _ := status.FromError(err)

		logger.Info("gRPC client request completed",
			zap.String("x-request-id", requestID),
			zap.String("method", method),
			zap.String("status", st.Code().String()),
			zap.Duration("duration", duration),
			zap.Error(err),
		)

		return err
	}
}
