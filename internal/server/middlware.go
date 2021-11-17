package server

import (
	"context"
	"strings"

	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ChangeLoggLevel() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			levels := md.Get("log-level")
			logger.InfoKV(ctx, "got log level", "levels", levels)
			if len(levels) > 0 {
				if parsedLevel, ok := parseLevel(levels[0]); ok {
					newLogger := logger.CloneWithLevel(ctx, parsedLevel)
					ctx = logger.AttachLogger(ctx, newLogger)
				}
			}
		}

		return handler(ctx, req)
	}
}

func displayDetailedRequestResponse() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		resp, err := handler(ctx, req)

		if ok && err == nil {
			display := md.Get("display-detailed-info")
			if len(display) > 0 && display[0] == "true" {
				logger.InfoKV(ctx, "detailed info",
					"info", info,
					"metadata", md,
					"request", req,
					"response", resp)
			}
		}

		return resp, err

	}
}

func parseLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.InfoLevel, true
	case "warn":
		return zapcore.WarnLevel, true
	case "error":
		return zapcore.ErrorLevel, true
	default:
		return zapcore.DebugLevel, false
	}
}
