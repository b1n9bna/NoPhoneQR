# NoPhoneQR

<p align="center">
  <a href="https://github.com/b1n9bna/NoPhoneQR/releases"><img src="https://img.shields.io/github/v/release/b1n9bna/NoPhoneQR?style=flat-square" alt="Release"></a>
  <img src="https://img.shields.io/badge/RAM_Peak-9.1_MB_(Benchmarked)-success?style=flat-square&logo=speedtest&logoColor=white" alt="Peak RAM">
  <img src="https://img.shields.io/badge/Binary_Size-3.3_MB-blue?style=flat-square&logo=files" alt="Binary Size">
  <img src="https://img.shields.io/badge/Cold_Start-<50_ms-orange?style=flat-square&logo=lightning" alt="Cold Start">
  <img src="https://img.shields.io/badge/Pure_Go-100%25_(Zero_CGO)-blueviolet?style=flat-square&logo=go" alt="Pure Go">
  <img src="https://img.shields.io/badge/Platform-macOS_|_Windows_|_Linux-brightgreen?style=flat-square" alt="Platform">
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square" alt="License"></a>
</p>

A desktop screen QR code scanner and URL sanitizer written in Go.

When reading technical docs, articles, or browsing the web on a computer, you often encounter on-screen QR codes. NoPhoneQR allows you to select an on-screen QR code using a hotkey or command line, strips tracking parameters, and opens the clean URL in your browser or copies it to your clipboard—saving you the hassle of using a physical phone to scan your computer screen.

[简体中文](README.md) • [English](README_en.md)

---

## Quick Start

### Option 1: Download Pre-built Binary

Download the executable for your platform from the [Releases page](https://github.com/b1n9bna/NoPhoneQR/releases):

- **macOS (Apple Silicon)**: `nophoneqr-darwin-arm64`
- **macOS (Intel)**: `nophoneqr-darwin-amd64`
- **Windows (x64)**: `nophoneqr-windows-amd64.exe`
- **Linux (x64)**: `nophoneqr-linux-amd64`

*(macOS / Linux users grant execution permissions: `chmod +x nophoneqr*`, and move it to your PATH such as `/usr/local/bin/nophoneqr`)*

### Option 2: Install via Go

If Go (1.22+) is installed locally:
```bash
go install github.com/b1n9bna/nophoneqr@latest
```

---

## Recommended: Set Up Global Hotkeys

Binding `nophoneqr` to a global hotkey (e.g. `Cmd + Shift + Q` or `Alt + Q`) provides the smoothest workflow:

### macOS
- **Apple Shortcuts**:
  1. Open the built-in "Shortcuts" app and create a new shortcut;
  2. Add a "Run Shell Script" action with `/usr/local/bin/nophoneqr`;
  3. Assign a keyboard shortcut in the right sidebar (e.g., `Cmd + Shift + Q`).
- **Raycast / Alfred**:
  Create a Script Command calling `nophoneqr` directly and assign a hotkey.

### Windows
1. Place `nophoneqr.exe` in a permanent directory (e.g. `C:\Program Files\NoPhoneQR\` or your user folder);
2. Create a desktop shortcut, right-click and choose "Properties";
3. In the "Shortcut key" field, press your desired combination (e.g., `Ctrl + Alt + Q`), then click OK.

---

## Usage

### 1. Interactive Screen Scan (Default)
Run directly from terminal or via your hotkey:
```bash
nophoneqr
```
The mouse pointer turns into a crosshair. Drag to select the on-screen QR code; press `ESC` to cancel.

### 2. Scan Local Image
```bash
nophoneqr /path/to/qrcode.png
```

### 3. Generate QR Codes (Auxiliary Utility)
Print a QR code as ANSI characters in the terminal, or save to a PNG file:
```bash
# Print to terminal console
nophoneqr -g "https://example.com"

# Save as image file
nophoneqr -g "https://example.com" -o qr.png
```

### Command Flags

| Flag | Description |
| :--- | :--- |
| `nophoneqr` | Default interactive screen selection |
| `nophoneqr <image>` | Scan QR code from a specified image file |
| `-no-open` | Copy URL to clipboard without auto-opening in browser |
| `-g <content>` | Generate ANSI character QR code in terminal |
| `-o <file.png>` | Used with `-g` to export QR code to a PNG image |
| `-v` | Display version information |

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

## Performance & Resource Benchmarks

Benchmarked on macOS 14 (Sonoma) using `/usr/bin/time -l` through a complete screen scan and decode cycle:

| Metric | NoPhoneQR (Native Go) | Typical Desktop Scanner | Notes |
| :--- | :--- | :--- | :--- |
| **Binary Size** | **3.3 MB** (Windows 4.7MB) | 150 MB ~ 300 MB | Standalone portable binary |
| **Peak RAM (RSS)** | **9.1 MB** | 200 MB ~ 400 MB | Memory allocated only during decoding |
| **Idle RAM Usage** | **0 MB** | 50 MB ~ 150 MB | Runs on demand, zero background residue |
| **Cold Start + Decode** | **< 50 ms** | 1.0 s ~ 3.0 s | Millisecond response when triggered |
| **Runtime Dependencies** | **None** | Node.js / Python / OpenCV | Pure Go with zero CGO bindings |

<details>
<summary>Click to view raw macOS benchmark console output</summary>

```bash
$ /usr/bin/time -l nophoneqr -no-open test_qr.png

✅ QR Code Recognized successfully!
  • Type:   url
  • Result: https://github.com/b1n9bna/NoPhoneQR
  • Action: Copied to clipboard

        0.28 real         0.06 user         0.07 sys
             9613312  maximum resident set size  # Peak RSS: ~9.1 MB
             2228224  peak memory footprint      # Core footprint: ~2.1 MB
```
</details>

---

## Build & Test

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
