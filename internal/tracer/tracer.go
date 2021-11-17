package tracer

import (
	"context"
	"fmt"
	"io"

	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/ozonmp/srv-verification-api/internal/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer - returns new tracer.
func NewTracer(cfg *config.Config) (io.Closer, error) {
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}
	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		logger.ErrorKV(context.Background(), fmt.Sprintf("failed init jaeger:"), "err", err)
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	logger.InfoKV(context.Background(), "Traces started")
	return closer, nil
}
