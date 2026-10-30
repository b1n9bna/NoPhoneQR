# NoPhoneQR

A desktop screen QR code scanner and URL sanitizer written in Go.

When reading technical docs, articles, or browsing the web on a computer, you often encounter on-screen QR codes. NoPhoneQR allows you to select an on-screen QR code using a hotkey or command line, strips tracking parameters, and opens the clean URL in your browser or copies it to your clipboard—saving you the hassle of using a physical phone to scan your computer screen.

[简体中文](README.md) • [English](README_en.md)

---

## Features

- **On-Screen Selection**: Drag and select any QR code directly on your screen without needing a phone or secondary device.
- **URL Parameter Sanitization**: Automatically strips common tracking and marketing bloat (`utm_*`, `spm_*`, `chksm`, `scene`, `fbclid`, etc.) to keep URLs clean.
- **Redirect Unwrapping**: Unwraps intermediate interstitial redirect pages (such as Zhihu and QQ redirection links) directly to the target destination.
- **Multi-Format Support**:
  - **Web URLs**: Sanitizes parameters, copies to clipboard, and opens in your default browser.
  - **Wi-Fi Config**: Parses network SSID and extracts password directly to clipboard.
  - **Two-Factor Auth (2FA)**: Parses `otpauth://` QR codes and extracts the Secret token.
  - **Plain Text**: Copies decoded text straight to system clipboard.
- **Lightweight & Portable**: Written in pure Go with zero CGO dependencies. Compiles to a single standalone binary without requiring Python, OpenCV, or Node.js runtime environments.

---

## Getting Started

### Method 1: Download Pre-compiled Binary

Download the executable for your operating system from the [Releases page](https://github.com/b1n9bna/NoPhoneQR/releases):

- **macOS (Apple Silicon)**: `nophoneqr-darwin-arm64`
- **macOS (Intel)**: `nophoneqr-darwin-amd64`
- **Windows (x64)**: `nophoneqr-windows-amd64.exe`
- **Linux (x64)**: `nophoneqr-linux-amd64`

*(For macOS / Linux: make the binary executable via `chmod +x nophoneqr*`, and move it to a directory in your PATH such as `/usr/local/bin/nophoneqr`)*

### Method 2: Install via Go

If you have Go 1.22+ installed:
```bash
go install github.com/b1n9bna/nophoneqr@latest
```

---

## Usage

### 1. Screen Selection (Default)
Run the command directly (or trigger via global hotkey):
```bash
nophoneqr
```
Your cursor will turn into a crosshair. Drag to select the QR code on your screen; press `ESC` anytime to cancel.

### 2. Decode a Local Image File
```bash
nophoneqr /path/to/qrcode.png
```

### 3. Generate QR Code (Helper Utility)
Print an ASCII QR code in your terminal, or save it as a PNG image:
```bash
# Print to terminal
nophoneqr -g "https://example.com"

# Save as PNG image
nophoneqr -g "https://example.com" -o qr.png
```

### Command Options

| Flag | Description |
| :--- | :--- |
| `nophoneqr` | Default interactive screen capture mode |
| `nophoneqr <image>` | Decode a QR code from a specified image file |
| `-no-open` | Copy URL to clipboard only without opening browser |
| `-g <content>` | Generate an ASCII QR code in the terminal |
| `-o <file.png>` | Save the generated QR code to a PNG file |
| `-v` | Display version information |

---

## Recommended: Global Hotkey Setup

Binding `nophoneqr` to a global hotkey (e.g. `Cmd + Shift + Q` or `Alt + Q`) provides the smoothest workflow:

### macOS
- **System Shortcuts**:
  1. Open the built-in **Shortcuts** app and create a new shortcut;
  2. Add a **Run Shell Script** action with `/usr/local/bin/nophoneqr`;
  3. Assign a keyboard shortcut in the right sidebar (e.g. `Cmd + Shift + Q`).
- **Raycast / Alfred**:
  Create a Script Command that executes `nophoneqr` and bind it to a hotkey.

### Windows
1. Place `nophoneqr.exe` in a permanent directory (such as `C:\Program Files\NoPhoneQR\` or your user directory);
2. Create a desktop shortcut, right-click and choose **Properties**;
3. In the **Shortcut key** field, press your desired combination (e.g. `Ctrl + Alt + Q`), then click OK.

---

## Local Build

```bash
git clone https://github.com/b1n9bna/NoPhoneQR.git
cd NoPhoneQR

# Build binary
go build -o nophoneqr .

# Run tests
go test -v ./...
```

---

## License

This project is licensed under the [MIT License](LICENSE).
