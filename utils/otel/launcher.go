package otel

import (
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/lightstep/otel-launcher-go/launcher"
)

type LauncherConfig struct {
	ServiceName string
	AccessToken string
}

func NewLauncher(cfg *config.Config) *launcher.Launcher {
	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(config.ServiceName),
		launcher.WithAccessToken(config.AccessToken),
	)

	return &ls
}
