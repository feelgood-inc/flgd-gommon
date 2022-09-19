package adapters

import "github.com/lightstep/otel-launcher-go/launcher"

type LauncherConfig struct {
	ServiceName string
	AccessToken string
}

func NewLauncher(launcherConfig LauncherConfig) *launcher.Launcher {
	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName(launcherConfig.ServiceName),
		launcher.WithAccessToken(launcherConfig.AccessToken),
	)

	return &ls
}
