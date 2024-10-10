package paths

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Paths(t *testing.T) {
	t.Run("FileKubeconfig", func(t *testing.T) {
		if d := cmp.Diff("sbctl.kubeconfig", FileKubeconfig); d != "" {
			t.Errorf("FileKubeconfig mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("UserHome", func(t *testing.T) {
		exp, _ := os.UserHomeDir()
		if d := cmp.Diff(exp, UserHome); d != "" {
			t.Errorf("UserHome mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("Sumboard", func(t *testing.T) {
		exp := filepath.Join(UserHome, ".sumboard")
		if d := cmp.Diff(exp, Sumboard); d != "" {
			t.Errorf("Sumboard mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("SbCtl", func(t *testing.T) {
		exp := filepath.Join(UserHome, ".sumboard", "sbctl")
		if d := cmp.Diff(exp, SbCtl); d != "" {
			t.Errorf("SbCtl mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("Data", func(t *testing.T) {
		exp := filepath.Join(UserHome, ".sumboard", "sbctl", "data")
		if d := cmp.Diff(exp, Data); d != "" {
			t.Errorf("Data mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("Kubeconfig", func(t *testing.T) {
		exp := filepath.Join(UserHome, ".sumboard", "sbctl", "sbctl.kubeconfig")
		if d := cmp.Diff(exp, Kubeconfig); d != "" {
			t.Errorf("Kubeconfig mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("HelmRepoConfig", func(t *testing.T) {
		exp := filepath.Join(UserHome, ".sumboard", "sbctl", ".helmrepo")
		if d := cmp.Diff(exp, HelmRepoConfig); d != "" {
			t.Errorf("HelmRepoConfig mismatch (-want +got):\n%s", d)
		}
	})

	t.Run("HelmRepoCache", func(t *testing.T) {
		exp := filepath.Join(UserHome, ".sumboard", "sbctl", ".helmcache")
		if d := cmp.Diff(exp, HelmRepoCache); d != "" {
			t.Errorf("HelmRepoCache mismatch (-want +got):\n%s", d)
		}
	})
}
