package otel

import (
	"github.com/lightstep/otel-launcher-go/launcher"
)

type LauncherConfig struct {
	ServiceName string
	AccessToken string
}

func NewLauncher(config *LauncherConfig) *launcher.Launcher {
	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(config.ServiceName),
		launcher.WithAccessToken(config.AccessToken),
	)

	return &ls
}
