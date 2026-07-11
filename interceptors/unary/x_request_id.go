package unary

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	XRequestIDKey = "x-request-id"
)

func XRequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestID := extractOrGenerateRequestID(ctx)
		ctx = context.WithValue(ctx, XRequestIDKey, requestID)

		ctx = metadata.AppendToOutgoingContext(ctx, XRequestIDKey, requestID)

		return handler(ctx, req)
	}
}

func XRequestIDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		requestID := extractOrGenerateRequestID(ctx)
		ctx = metadata.AppendToOutgoingContext(ctx, XRequestIDKey, requestID)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(XRequestIDKey).(string); ok {
		return requestID
	}
	return ""
}

func extractOrGenerateRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ids := md.Get(XRequestIDKey)
		if len(ids) > 0 && ids[0] != "" {
			return ids[0]
		}
	}

	mdOut, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		ids := mdOut.Get(XRequestIDKey)
		if len(ids) > 0 && ids[0] != "" {
			return ids[0]
		}
	}

	return uuid.New().String()
}

func GenerateRequestID() string {
	return uuid.New().String()
}
