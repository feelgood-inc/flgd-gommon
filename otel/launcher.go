package otel

import (
	"github.com/lightstep/otel-launcher-go/launcher"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

var (
	endpoint = "ingest.lightstep.com:443"
)

type LauncherConfig struct {
	ServiceName    string
	ServiceVersion string
	AccessToken    string
	Env            string
}

func NewLauncher(config LauncherConfig) launcher.Launcher {
	otelLauncher := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(config.ServiceName),
		launcher.WithServiceVersion(config.ServiceVersion),
		launcher.WithAccessToken(config.AccessToken),
		launcher.WithSpanExporterEndpoint(endpoint),
		launcher.WithMetricExporterEndpoint(endpoint),
		launcher.WithResourceAttributes(map[string]string{
			string(semconv.ContainerNameKey):         config.ServiceName,
			string(semconv.DeploymentEnvironmentKey): config.Env,
		}),
	)

	return otelLauncher
}
