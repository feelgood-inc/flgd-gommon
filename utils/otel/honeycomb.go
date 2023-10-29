package otel

import (
	"fmt"
	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"os"
)

type InitHoneycombConfig struct {
	Namespace   string
	ServiceName string
	SampleRate  float64
	OTELMetrics bool
}

func InitHoneycomb(cfg *InitHoneycombConfig) (func(), error) {
	// Set some env vars for honeycomb
	if cfg.OTELMetrics {
		err := os.Setenv("OTEL_METRICS_ENABLED", "true")
		if err != nil {
			return nil, err
		}
	}

	if cfg.SampleRate > 0 {
		err := os.Setenv("SAMPLE_RATE", fmt.Sprintf("%f", cfg.SampleRate))
		if err != nil {
			return nil, err
		}
	}

	// enable multi-span attributes
	bsp := honeycomb.NewBaggageSpanProcessor()

	// use honeycomb distro to setup OpenTelemetry SDK
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(bsp),
		otelconfig.WithServiceName(cfg.ServiceName),
		otelconfig.WithResourceAttributes(map[string]string{
			"service.name":      cfg.ServiceName,
			"service.namespace": cfg.Namespace,
		}),
	)
	if err != nil {
		return nil, err
	}

	return otelShutdown, nil
}
