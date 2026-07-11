package stream

import (
	"shared/interceptors/unary"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func StreamLoggerInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()
		requestID := unary.GetRequestID(ss.Context())

		logger.Info("gRPC stream started",
			zap.String("x-request-id", requestID),
			zap.String("method", info.FullMethod),
			zap.Bool("is_client_stream", info.IsClientStream),
			zap.Bool("is_server_stream", info.IsServerStream),
		)

		err := handler(srv, ss)

		duration := time.Since(start)

		logLevel := zap.InfoLevel
		if err != nil {
			logLevel = zap.ErrorLevel
		}

		logger.Log(logLevel, "gRPC stream completed",
			zap.String("x-request-id", requestID),
			zap.String("method", info.FullMethod),
			zap.Duration("duration", duration),
			zap.Error(err),
		)

		return err
	}
}
