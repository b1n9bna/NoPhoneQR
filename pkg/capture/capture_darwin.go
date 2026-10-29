//go:build darwin
// +build darwin

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

	// /usr/sbin/screencapture -i brings up native interactive selection
	cmd := exec.Command("/usr/sbin/screencapture", "-i", tmpFile)
	err := cmd.Run()

	// If screencapture exited or was aborted with ESC
	fi, statErr := os.Stat(tmpFile)
	if os.IsNotExist(statErr) || (statErr == nil && fi.Size() == 0) {
		return nil, ErrCanceled
	}
	if err != nil {
		return nil, fmt.Errorf("screencapture error: %w", err)
	}

	f, err := os.Open(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open captured screenshot: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode captured screenshot: %w", err)
	}

	return img, nil
}
