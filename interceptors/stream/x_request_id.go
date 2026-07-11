package stream

import (
	"context"

	"shared/interceptors/unary"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func StreamXRequestIDInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()
		requestID := extractOrGenerateRequestID(ctx)
		ctx = context.WithValue(ctx, unary.XRequestIDKey, requestID)
		ctx = metadata.AppendToOutgoingContext(ctx, unary.XRequestIDKey, requestID)

		wrappedStream := &wrappedServerStream{
			ServerStream: ss,
			ctx:          ctx,
		}

		return handler(srv, wrappedStream)
	}
}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

func extractOrGenerateRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ids := md.Get(unary.XRequestIDKey)
		if len(ids) > 0 && ids[0] != "" {
			return ids[0]
		}
	}

	mdOut, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		ids := mdOut.Get(unary.XRequestIDKey)
		if len(ids) > 0 && ids[0] != "" {
			return ids[0]
		}
	}

	return unary.GenerateRequestID()
}
