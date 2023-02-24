package adapters

import (
	"context"
	"errors"
	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"os"
)

type LauncherConfig struct {
	ServiceName string
	AccessToken string
}

type ExporterConfig struct {
	Endpoint    string
	AccessToken string
}

type TraceProviderConfig struct {
	ServiceName    string
	ServiceVersion string
	Env            string
}

type Config struct {
	ExporterConfig      ExporterConfig
	TraceProviderConfig TraceProviderConfig
}

func NewLauncher(launcherConfig LauncherConfig) *launcher.Launcher {
	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(launcherConfig.ServiceName),
		launcher.WithAccessToken(launcherConfig.AccessToken),
	)

	return &ls
}

func newExporter(ctx context.Context, config ExporterConfig) (*otlptrace.Exporter, error) {
	// Check variables are set
	var token string
	if config.AccessToken == "" {
		token = os.Getenv("LS_ACCESS_TOKEN")
		if token == "" {
			return nil, errors.New("access token not set")
		}
	} else {
		token = config.AccessToken
	}

	var endpoint string
	if config.Endpoint == "" {
		endpoint = os.Getenv("LS_ENDPOINT")
		if endpoint == "" {
			return nil, errors.New("endpoint not set")
		}
	} else {
		endpoint = config.Endpoint
	}

	// Create a new OTLP trace exporter.
	var headers = map[string]string{
		"lightstep-access-token": token,
	}

	client := otlptracehttp.NewClient(
		otlptracehttp.WithHeaders(headers),
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithURLPath("traces/otlp/v0.9"),
	)
	return otlptrace.New(ctx, client)
}

func newTraceProvider(exp *otlptrace.Exporter, cfg TraceProviderConfig) *sdktrace.TracerProvider {
	rsrc, rErr :=
		resource.Merge(
			resource.Default(),
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(cfg.ServiceName),
				semconv.ServiceVersionKey.String(cfg.ServiceVersion),
				attribute.String("environment", cfg.Env),
			),
		)

	if rErr != nil {
		panic(rErr)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(rsrc),
	)
}

func NewTraceProvider(ctx context.Context, cfg Config) *sdktrace.TracerProvider {
	exporter, err := newExporter(ctx, cfg.ExporterConfig)
	if err != nil {
		panic(err)
	}

	return newTraceProvider(exporter, cfg.TraceProviderConfig)
}
