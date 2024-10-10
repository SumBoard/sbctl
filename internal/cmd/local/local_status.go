package local

import (
	"context"
	"fmt"

	"github.com/SumBoard/sbctl/internal/cmd/local/k8s"
	"github.com/SumBoard/sbctl/internal/cmd/local/local"
	"github.com/SumBoard/sbctl/internal/telemetry"
	"github.com/SumBoard/sbctl/internal/trace"
	"github.com/pterm/pterm"
)

type StatusCmd struct{}

func (s *StatusCmd) Run(ctx context.Context, provider k8s.Provider, telClient telemetry.Client) error {
	ctx, span := trace.NewSpan(ctx, "local status")
	defer span.End()

	spinner := &pterm.DefaultSpinner
	if err := checkDocker(ctx, telClient, spinner); err != nil {
		return err
	}

	return telClient.Wrap(ctx, telemetry.Status, func() error {
		return status(ctx, provider, telClient, spinner)
	})
}

func checkDocker(ctx context.Context, telClient telemetry.Client, spinner *pterm.SpinnerPrinter) error {
	spinner, _ = spinner.Start("Starting status check")
	spinner.UpdateText("Checking for Docker installation")

	_, err := dockerInstalled(ctx, telClient)
	if err != nil {
		pterm.Error.Println("Unable to determine if Docker is installed")
		return fmt.Errorf("unable to determine docker installation status: %w", err)
	}

	return nil
}

func status(ctx context.Context, provider k8s.Provider, telClient telemetry.Client, spinner *pterm.SpinnerPrinter) error {
	spinner.UpdateText(fmt.Sprintf("Checking for existing Kubernetes cluster '%s'", provider.ClusterName))

	cluster, err := provider.Cluster(ctx)
	if err != nil {
		pterm.Error.Printfln("Unable to determine status of any existing '%s' cluster", provider.ClusterName)
		return err
	}

	if !cluster.Exists(ctx) {
		pterm.Warning.Println("Sumboard does not appear to be installed locally")
		return nil
	}

	pterm.Success.Printfln("Existing cluster '%s' found", provider.ClusterName)
	spinner.UpdateText(fmt.Sprintf("Validating existing cluster '%s'", provider.ClusterName))

	port, err := getPort(ctx, provider.ClusterName)
	if err != nil {
		return err
	}

	lc, err := local.New(provider,
		local.WithPortHTTP(port),
		local.WithTelemetryClient(telClient),
		local.WithSpinner(spinner),
	)
	if err != nil {
		pterm.Error.Printfln("Failed to initialize 'local' command")
		return fmt.Errorf("unable to initialize local command: %w", err)
	}

	if err := lc.Status(ctx); err != nil {
		spinner.Fail("Unable to install Sumboard locally")
		return err
	}

	_ = spinner.Stop()

	return nil
}
