package logger

import (
	"context"
	"os"

	"github.com/ozonmp/srv-verification-api/internal/config"
	"go.uber.org/zap/zapcore"

	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
)

type ctxKey struct{}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger

// InitLogger : create logger
func InitLogger(ctx context.Context, cfg *config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		FatalKV(ctx, "Fatal error init graylog", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	SetLogger(sugaredLogger.With(
		"service", cfg.Project.Name,
	))

	return func() {
		notSugaredLogger.Sync()
	}
}

func CloneWithLevel(ctx context.Context, newLevel zapcore.Level) *zap.SugaredLogger {
	return fromContext(ctx).
		Desugar().
		WithOptions(WithLevel(newLevel)).
		Sugar()
}

func fromContext(ctx context.Context) *zap.SugaredLogger {
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		return attachedLogger
	}

	return globalLogger
}

func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Errorw(message, kvs...)
}

func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warnw(message, kvs...)
}

func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Infow(message, kvs...)
}

func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debugw(message, kvs...)
}

func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatalw(message, kvs...)
}

func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

func SetLogger(newLogger *zap.SugaredLogger) {
	globalLogger = newLogger
}

func init() {
	notSugaredLogger, err := zap.NewProduction()
	if err != nil {
		FatalKV(context.Background(), "Fatal error init logger", err)
	}

	globalLogger = notSugaredLogger.Sugar()
}
