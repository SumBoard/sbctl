package paths

import (
	"os"
	"path/filepath"
)

const (
	FileKubeconfig = "sbctl.kubeconfig"
)

var (
	// UserHome is the user's home directory
	UserHome = func() string {
		h, _ := os.UserHomeDir()
		return h
	}()

	// Sumboard is the full path to the ~/.sumboard directory
	Sumboard = sumboard()

	// SbCtl is the full path to the ~/.sumboard/sbctl directory
	SbCtl = sbctl()

	// Data is the full path to the ~/.sumboard/sbctl/data directory
	Data = data()

	// Kubeconfig is the full path to the kubeconfig file
	Kubeconfig = kubeconfig()

	// HelmRepoConfig is the full path to where helm stores
	// its repository configurations.
	HelmRepoConfig = helmRepoConfig()

	// HelmRepoCache is the full path to where helm stores
	// its cached data.
	HelmRepoCache = helmRepoCache()
)

func sumboard() string {
	return filepath.Join(UserHome, ".sumboard")
}

func sbctl() string {
	return filepath.Join(sumboard(), "sbctl")
}

func data() string {
	return filepath.Join(sbctl(), "data")
}

func kubeconfig() string {
	return filepath.Join(sbctl(), FileKubeconfig)
}

func helmRepoConfig() string { return filepath.Join(sbctl(), ".helmrepo") }

func helmRepoCache() string { return filepath.Join(sbctl(), ".helmcache") }
