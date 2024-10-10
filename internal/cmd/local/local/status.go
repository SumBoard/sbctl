package local

import (
	"context"
	"fmt"

	"github.com/SumBoard/sbctl/internal/common"
	"github.com/pterm/pterm"
	"go.opencensus.io/trace"
)

// Status handles the status of local.
func (c *Command) Status(ctx context.Context) error {
	_, span := trace.StartSpan(ctx, "command.Status")
	defer span.End()

	charts := []string{common.SbChartRelease, common.NginxChartRelease}
	for _, name := range charts {
		c.spinner.UpdateText(fmt.Sprintf("Verifying %s Helm Chart installation status", name))

		rel, err := c.helm.GetRelease(name)
		if err != nil {
			pterm.Warning.Println("Unable to fetch sumboard release")
			pterm.Debug.Printfln("unable to fetch sumboard release: %s", err)
			continue
		}

		pterm.Info.Println(fmt.Sprintf(
			"Found helm chart '%s'\n  Status: %s\n  Chart Version: %s\n  App Version: %s",
			name, rel.Info.Status.String(), rel.Chart.Metadata.Version, rel.Chart.Metadata.AppVersion,
		))
	}

	pterm.Info.Println(fmt.Sprintf("Sumboard should be accessible via http://localhost:%d", c.portHTTP))

	return nil
}
