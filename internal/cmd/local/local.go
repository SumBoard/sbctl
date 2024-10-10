package local

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/SumBoard/sbctl/internal/cmd/local/k8s"
	"github.com/SumBoard/sbctl/internal/cmd/local/localerr"
	"github.com/SumBoard/sbctl/internal/cmd/local/paths"
	"github.com/pterm/pterm"
)

type Cmd struct {
	Credentials CredentialsCmd `cmd:"" help:"Get local Sumboard user credentials."`
	Install     InstallCmd     `cmd:"" help:"Install local Sumboard."`
	Deployments DeploymentsCmd `cmd:"" help:"View local Sumboard deployments."`
	Status      StatusCmd      `cmd:"" help:"Get local Sumboard status."`
	Uninstall   UninstallCmd   `cmd:"" help:"Uninstall local Sumboard."`
}

func (c *Cmd) BeforeApply() error {
	if _, envVarDNT := os.LookupEnv("DO_NOT_TRACK"); envVarDNT {
		pterm.Info.Println("Telemetry collection disabled (DO_NOT_TRACK)")
	}
	if err := checkSumboardDir(); err != nil {
		return fmt.Errorf("%w: %w", localerr.ErrSumboardDir, err)
	}

	return nil
}

func (c *Cmd) AfterApply(provider k8s.Provider) error {
	pterm.Info.Println(fmt.Sprintf(
		"Using Kubernetes provider:\n  Provider: %s\n  Kubeconfig: %s\n  Context: %s",
		provider.Name, provider.Kubeconfig, provider.Context,
	))

	return nil
}

// checkSumboardDir verifies that, if the paths.Sumboard directory exists, that it has proper permissions.
// If the directory does not have the proper permissions, this method will attempt to fix them.
// A nil response either indicates that either:
// - no paths.Sumboard directory exists
// - the permissions are already correct
// - this function was able to fix the incorrect permissions.
func checkSumboardDir() error {
	fileInfo, err := os.Stat(paths.Sumboard)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// nothing to do, directory will be created later on
			return nil
		}
		return fmt.Errorf("unable to determine status of '%s': %w", paths.Sumboard, err)
	}

	if !fileInfo.IsDir() {
		return errors.New(paths.Sumboard + " is not a directory")
	}

	if fileInfo.Mode().Perm() >= 0744 {
		// directory has minimal permissions
		return nil
	}

	if err := os.Chmod(paths.Sumboard, 0744); err != nil {
		return fmt.Errorf("unable to change permissions of '%s': %w", paths.Sumboard, err)
	}

	return nil
}
