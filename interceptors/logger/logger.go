package logger

import (
	"context"
	"time"

	log "github.com/erdedan1/shared/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LoggerServerInterceptor(log log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		md, _ := metadata.FromIncomingContext(ctx)
		rid := md.Get("x-request-id")

		log.Info(
			"Interceptors",
			"LoggerServerInterceptor",
			"grpc request handled",
			"method", info.FullMethod,
			"request_id", firstOrEmpty(rid),
			"duration", time.Since(start),
			"error", err,
		)

		return resp, err
	}
}

func firstOrEmpty(v []string) string {
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
