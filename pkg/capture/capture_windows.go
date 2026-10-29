//go:build windows
// +build windows

package capture

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/png"
	"os/exec"
	"time"

	"github.com/kbinani/screenshot"
)

func snipScreen() (image.Image, error) {
	// 1. Check if clipboard already has an image (e.g. from a recent Win+Shift+S)
	if img, err := getClipboardImage(); err == nil && img != nil {
		return img, nil
	}

	// 2. Trigger Windows built-in snipping overlay: ms-screenclip:
	cmd := exec.Command("cmd", "/c", "start", "ms-screenclip:")
	if err := cmd.Start(); err != nil {
		// Fallback: capture primary monitor directly if snipping protocol fails
		bounds := screenshot.GetDisplayBounds(0)
		return screenshot.CaptureRect(bounds)
	}

	// Wait up to 10 seconds for user to complete the snip onto the clipboard
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		time.Sleep(500 * time.Millisecond)
		if img, err := getClipboardImage(); err == nil && img != nil {
			return img, nil
		}
	}

	return nil, ErrCanceled
}

func getClipboardImage() (image.Image, error) {
	// Fetch clipboard image via PowerShell System.Windows.Forms
	psScript := `Add-Type -AssemblyName System.Windows.Forms; $img = [System.Windows.Forms.Clipboard]::GetImage(); if ($img) { $ms = New-Object System.IO.MemoryStream; $img.Save($ms, [System.Drawing.Imaging.ImageFormat]::Png); [System.Convert]::ToBase64String($ms.ToArray()) }`
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", psScript)
	out, err := cmd.Output()
	if err != nil || len(bytes.TrimSpace(out)) == 0 {
		return nil, fmt.Errorf("no image in clipboard")
	}

	data, err := base64.StdEncoding.DecodeString(string(bytes.TrimSpace(out)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode clipboard image base64: %w", err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode clipboard image: %w", err)
	}

	return img, nil
}
