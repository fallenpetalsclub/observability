package main

import (
	"context"
	"log/slog"

	"github.com/zalgonoise/go-diagrams/diagram"
	"github.com/zalgonoise/go-diagrams/nodes/oci"
	"github.com/zalgonoise/x/cli/v2"

	"github.com/fallenpetalsclub/shared/log"
)

func main() {
	runner := cli.NewRunner("diagrams",
		cli.WithExecutors(map[string]cli.Executor{
			"generate": cli.Executable(ExecGenerate),
		}),
	)

	cli.Run(runner, log.New("debug", true, false))
}

func ExecGenerate(_ context.Context, _ *slog.Logger, _ []string) (int, error) {
	if err := generateInfra(); err != nil {
		return 1, err
	}

	return 0, nil
}

func generateInfra() error {
	d, err := diagram.New(
		diagram.BaseDir("diagrams"),
		diagram.Direction("LR"),
		diagram.Filename("observability_infra"),
		diagram.Label("Observability Infra"),
	)
	if err != nil {
		return err
	}

	service := oci.Database.Dis().Label("Service")

	logging := oci.Database.Science().Label("Logging")
	metrics := oci.Database.Science().Label("Metrics")
	tracing := oci.Database.Science().Label("Tracing")
	profiling := oci.Database.Science().Label("Profiling")

	server := diagram.NewGroup("server", diagram.WithBackground(diagram.BackgroundPurple)).
		Label("Server").Add(service)
	server.NewGroup("observability").Label("Observability").Add(logging, metrics, tracing, profiling)

	d.Group(server)

	d.Connect(service, logging)
	d.Connect(service, tracing)
	d.Connect(service, metrics)
	d.Connect(service, profiling)

	if err := d.Render(); err != nil {
		return err
	}

	return nil
}
