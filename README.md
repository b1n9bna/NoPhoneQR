<p align="center">
  <h1 align="center">📱❌ NoPhoneQR</h1>
  <p align="center">
    <strong>Stop pointing your phone at your computer screen.</strong><br>
    Instant on-screen QR code scanner, URL sanitizer & automatic browser opener for desktop.
  </p>
  <p align="center">
    <a href="README.md"><b>English</b></a> •
    <a href="README_zh.md">简体中文</a>
  </p>
  <p align="center">
    <a href="https://github.com/b1n9bna/NoPhoneQR/releases"><img src="https://img.shields.io/github/v/release/b1n9bna/NoPhoneQR?style=flat-square" alt="Release"></a>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go" alt="Go"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square" alt="License"></a>
    <img src="https://img.shields.io/badge/Platform-macOS%20%7C%20Windows%20%7C%20Linux-brightgreen?style=flat-square" alt="Platform">
  </p>
  <p align="center">
    <img src="assets/demo.gif" alt="NoPhoneQR vs Traditional Comparison" width="880" style="border-radius: 8px; box-shadow: 0 4px 20px rgba(0,0,0,0.5);">
  </p>
</p>

---

## 💡 Why NoPhoneQR?

Whenever you encounter a QR code on your computer screen:
- Links in documentation, articles, or GitHub repositories
- Login / 2FA two-factor authentication setup (`otpauth://`)
- Wi-Fi credentials shared on screen
- Short links in social media or slide presentations

### ⚡ The Ultimate Showdown

| Dimension | ❌ Traditional Phone Method | ⚡ With NoPhoneQR |
| :--- | :--- | :--- |
| **Workflow** | Pull out phone ➜ Unlock ➜ Open camera ➜ Aim at screen ➜ Copy link ➜ Send to chat ➜ Open on PC (7 steps) | **Global Hotkey ➜ Drag select ➜ Browser opens instantly (1 step)** |
| **Time Spent** | **~35 seconds** (breaks your focus and flow) | **~0.8 seconds** (instant, seamless, friction-free) |
| **URL Cleanliness** | Carries ugly tracking params (`utm_*`, `spm_*`, `chksm`) | **Automatically strips tracking bloat for clean URLs** |
| **Redirects** | Annoying intermediate interstitial warning pages | **Unwraps and bypasses redirects automatically** |
| **Special QR Codes** | 2FA secrets or Wi-Fi passwords must be retyped | **Extracts secret / password straight to clipboard** |

---

## ✨ Features

- ⚡ **Ultra-Lightweight**: Single static binary (~3–4MB), zero Electron bloat, <15MB RAM usage, sub-millisecond cold start.
- 🛡️ **URL Sanitizer**: Automatically strips tracking parameters (`utm_*`, `spm_*`, `chksm`, `scene`, `fbclid`, `gclid`) and unwraps interstitial redirects.
- 🧠 **Smart Payload Dispatching**:
  - **Web URLs**: Cleans URL + Launches default browser in new tab + Copies to clipboard.
  - **Wi-Fi QRs**: Parses SSID and copies password directly to clipboard.
  - **2FA / OTP**: Extracts account and OTP secret token without typing.
  - **Plain Text / Numbers**: Copies to clipboard with a native system toast notification.
- 🖥️ **True Cross-Platform**:
  - **macOS**: Native crosshair snippet with Retina HiDPI support.
  - **Windows**: Seamless Windows 10/11 Snipping & clipboard extraction.
  - **Linux**: Wayland (grim/slurp) and X11 (maim/scrot) support.

---

## 🚀 Quick Start (Pick ONE of the options below)

> 💡 **Note**: You only need to choose **EITHER Option 1 OR Option 2**. For everyday users without programming tools, **Option 1** is strongly recommended (no Go installation required).

### Option 1: Direct Download (Recommended, Zero Setup)
Download the standalone binary for your system from the [Releases page](https://github.com/b1n9bna/NoPhoneQR/releases):
- **macOS Apple Silicon (M1/M2/M3/M4)**: Download `nophoneqr-darwin-arm64`
- **macOS Intel**: Download `nophoneqr-darwin-amd64`
- **Windows (x64)**: Download `nophoneqr-windows-amd64.exe`
- **Linux (x64)**: Download `nophoneqr-linux-amd64`

*(macOS / Linux: Make it executable via `chmod +x nophoneqr*` and move to `/usr/local/bin/nophoneqr`)*

### Option 2: Install via Go (For Developers)
If you have Go 1.22+ installed:
```bash
go install github.com/b1n9bna/nophoneqr@latest
```

---

## 💻 Usage

### Scenario 1: Interactive On-Screen Snip (Most Common)
In your terminal, or bound to a global shortcut:
```bash
nophoneqr
```
Your cursor immediately turns into a crosshair. Drag a box over any QR code on screen:
- The destination URL opens in your default browser and is copied to your clipboard.
- Press `ESC` at any time to cancel.

### Scenario 2: Decode an Existing Image File
```bash
nophoneqr screenshot.png
# or
nophoneqr -f /path/to/qrcode.jpg
```

### Scenario 3: Generate QR Codes (Bonus Utility)
Render ASCII QR codes right in your terminal, or save directly as PNG:
```bash
# Render ASCII QR in terminal
nophoneqr -g "https://github.com/b1n9bna/NoPhoneQR"

# Save as high-res PNG image
nophoneqr -g "https://example.com" -o qr.png
```

### CLI Flags
| Flag | Description |
| :--- | :--- |
| `nophoneqr` | Default interactive on-screen snip & decode |
| `nophoneqr <image>` | Decode QR code from specified image file |
| `-no-open` | Copy decoded link to clipboard only (do not auto-open browser) |
| `-g <content>` | Generate ASCII QR code in terminal |
| `-o <file.png>` | Save generated QR code as a PNG file |
| `-v` | Display current version |

---

## ⌨️ Recommended: Set up a Global Hotkey

For the ultimate experience, bind `nophoneqr` to a global hotkey (e.g. `Cmd + Shift + Q` or `Alt + Q`):

### 🍏 macOS Setup:
1. **Raycast / Alfred**:
   - Create a `Script Command` running `nophoneqr`, bound to `Cmd + Shift + Q`.
2. **Native macOS Shortcuts**:
   - Open the **Shortcuts** app, create a new shortcut.
   - Add a "Run Shell Script" action with `/usr/local/bin/nophoneqr`.
   - Set keyboard shortcut in the right panel (e.g. `Cmd + Shift + Q`).

### 🪟 Windows Setup:
1. Download `nophoneqr-windows-amd64.exe`, rename to `nophoneqr.exe`, and place in a folder in your PATH (e.g. `C:\Windows\`).
2. Right-click the executable to create a desktop shortcut, open Properties, and assign a shortcut key like `Ctrl + Alt + Q`!

---

## 🛠️ Build from Source

```bash
git clone https://github.com/b1n9bna/NoPhoneQR.git
cd NoPhoneQR

# Build local binary
go build -o nophoneqr .

# Run test suite
go test -v ./...

# Cross-compile for all platforms (macOS, Windows, Linux)
make cross-build
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE). Contributions and pull requests are welcome!
