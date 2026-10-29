//go:build !darwin && !windows
// +build !darwin,!windows

package capture

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"os/exec"
	"path/filepath"
)

func snipScreen() (image.Image, error) {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, fmt.Sprintf("nophoneqr_%d.png", os.Getpid()))
	defer os.Remove(tmpFile)

	var cmd *exec.Cmd

	// Try common Linux screenshot tools: grim/slurp (Wayland), maim, scrot, gnome-screenshot
	if _, err := exec.LookPath("grim"); err == nil {
		if _, err := exec.LookPath("slurp"); err == nil {
			cmd = exec.Command("sh", "-c", fmt.Sprintf("grim -g \"$(slurp)\" %s", tmpFile))
		}
	} else if _, err := exec.LookPath("maim"); err == nil {
		cmd = exec.Command("maim", "-s", tmpFile)
	} else if _, err := exec.LookPath("scrot"); err == nil {
		cmd = exec.Command("scrot", "-s", tmpFile)
	} else if _, err := exec.LookPath("gnome-screenshot"); err == nil {
		cmd = exec.Command("gnome-screenshot", "-a", "-f", tmpFile)
	} else {
		return nil, fmt.Errorf("no supported screenshot utility found (please install grim/slurp, maim, or scrot)")
	}

	if err := cmd.Run(); err != nil {
		return nil, ErrCanceled
	}

	fi, statErr := os.Stat(tmpFile)
	if os.IsNotExist(statErr) || (statErr == nil && fi.Size() == 0) {
		return nil, ErrCanceled
	}

	f, err := os.Open(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open screenshot: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode screenshot: %w", err)
	}

	return img, nil
}
