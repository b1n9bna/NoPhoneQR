package dispatcher

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/b1n9bna/nophoneqr/pkg/cleaner"
	"github.com/b1n9bna/nophoneqr/pkg/notify"
)

// Result represents the processed QR payload.
type Result struct {
	Type        string // "url", "wifi", "otp", "text"
	Raw         string
	Processed   string
	CopiedText  string
	ActionTaken string
}

// Dispatch inspects the raw text decoded from a QR code, categorizes it,
// copies relevant information to the clipboard, opens browsers if applicable,
// and sends a desktop notification.
func Dispatch(raw string, autoOpen bool) (*Result, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, fmt.Errorf("empty QR code content")
	}

	res := &Result{
		Raw: raw,
	}

	switch {
	case strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://"):
		res.Type = "url"
		cleanURL := cleaner.CleanURL(raw)
		res.Processed = cleanURL
		res.CopiedText = cleanURL

		// Copy clean URL to system clipboard
		_ = clipboard.WriteAll(cleanURL)

		if autoOpen {
			if err := OpenBrowser(cleanURL); err != nil {
				res.ActionTaken = "Copied to clipboard (browser launch failed)"
			} else {
				res.ActionTaken = "Opened in browser & copied to clipboard"
			}
		} else {
			res.ActionTaken = "Copied to clipboard"
		}

		notifyMsg := truncate(cleanURL, 80)
		_ = notify.Send("NoPhoneQR 🔗 网页已打开", notifyMsg)

	case strings.HasPrefix(raw, "WIFI:"):
		res.Type = "wifi"
		ssid, pass := parseWiFi(raw)
		res.Processed = fmt.Sprintf("SSID: %s, Password: %s", ssid, pass)
		res.CopiedText = pass

		if pass != "" {
			_ = clipboard.WriteAll(pass)
			res.ActionTaken = "Wi-Fi password copied to clipboard"
			_ = notify.Send("NoPhoneQR 📶 Wi-Fi 信息", fmt.Sprintf("SSID: %s\n密码已复制到剪贴板", ssid))
		} else {
			_ = clipboard.WriteAll(raw)
			res.ActionTaken = "Wi-Fi config copied to clipboard"
			_ = notify.Send("NoPhoneQR 📶 Wi-Fi 信息", fmt.Sprintf("SSID: %s (开放网络)", ssid))
		}

	case strings.HasPrefix(raw, "otpauth://"):
		res.Type = "otp"
		account, secret := parseOTP(raw)
		res.Processed = fmt.Sprintf("Account: %s, Secret: %s", account, secret)
		res.CopiedText = secret

		if secret != "" {
			_ = clipboard.WriteAll(secret)
			res.ActionTaken = "2FA secret copied to clipboard"
		} else {
			_ = clipboard.WriteAll(raw)
			res.ActionTaken = "OTP URI copied to clipboard"
		}
		_ = notify.Send("NoPhoneQR 🔑 2FA 密钥", fmt.Sprintf("账号: %s\n密钥已复制到剪贴板", account))

	default:
		res.Type = "text"
		res.Processed = raw
		res.CopiedText = raw

		_ = clipboard.WriteAll(raw)
		res.ActionTaken = "Copied to clipboard"

		snippet := truncate(raw, 80)
		_ = notify.Send("NoPhoneQR 📋 内容已复制", snippet)
	}

	return res, nil
}

// OpenBrowser launches the default web browser with the target URL.
func OpenBrowser(targetURL string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", targetURL)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", targetURL)
	default: // linux, freebsd, etc.
		cmd = exec.Command("xdg-open", targetURL)
	}

	return cmd.Start()
}

func parseWiFi(raw string) (ssid, pass string) {
	// Format: WIFI:S:MySSID;T:WPA;P:MyPassword;;
	content := strings.TrimPrefix(raw, "WIFI:")
	parts := strings.Split(content, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "S:") {
			ssid = strings.TrimPrefix(part, "S:")
		} else if strings.HasPrefix(part, "P:") {
			pass = strings.TrimPrefix(part, "P:")
		}
	}
	return
}

func parseOTP(raw string) (account, secret string) {
	u, err := url.Parse(raw)
	if err != nil {
		return "Unknown", raw
	}
	account = strings.TrimPrefix(u.Path, "/")
	secret = u.Query().Get("secret")
	return
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
