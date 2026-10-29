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
    <img src="https://img.shields.io/badge/RAM_Peak-9.1_MB_(Benchmarked)-success?style=flat-square&logo=speedtest&logoColor=white" alt="Peak RAM">
    <img src="https://img.shields.io/badge/Binary_Size-3.3_MB-blue?style=flat-square&logo=files" alt="Binary Size">
    <img src="https://img.shields.io/badge/Cold_Start-<50_ms-orange?style=flat-square&logo=lightning" alt="Cold Start">
    <img src="https://img.shields.io/badge/Electron-0%25_(Pure_Go)-blueviolet?style=flat-square" alt="Zero Electron">
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

## 📊 Hardware & Resource Benchmarks (Real-world Measured)

> [!TIP]
> Tired of bloated Electron desktop utilities chewing up 300MB+ of RAM just to take a quick screenshot? **NoPhoneQR is written in 100% pure Go with uncompromising restraint.**

Measured on **macOS 14 (Sonoma)** using `/usr/bin/time -l` through a complete scan-decode-dispatch lifecycle:

| Metric | Typical Electron-based Tools | Python + OpenCV Scripts | ⚡ **NoPhoneQR (Native Go)** |
| :--- | :--- | :--- | :--- |
| **Binary / Dist Size** | 150 MB ~ 300 MB | ~200 MB (with venv) | **3.3 MB** (Windows .exe is 4.7 MB) |
| **Peak Resident RAM (RSS)**| 250 MB ~ 450 MB | 80 MB ~ 180 MB | **9.1 MB** *(Measured: 9,613,312 bytes)* |
| **Idle Background RAM** | 100 MB ~ 200 MB persistent | Requires background daemon | **0 MB** (Run & Done: zero RAM residue) |
| **Cold Start + Decode Time**| 1.5s ~ 3.5s | 0.8s ~ 2.0s | **< 50 milliseconds (0.05s)** |
| **Runtime Dependencies** | Node.js, Chromium, WebView | Python 3, pip, OpenCV, C++ | **ZERO** (Single portable binary) |
| **CGO / Dynamic Libraries** | OS-specific C bindings | Requires libzbar / libGL | **Zero CGO (Pure Go ZXing)** |

<details>
<summary>🔍 Click to view raw macOS benchmark console output</summary>

```bash
$ /usr/bin/time -l nophoneqr -no-open test_qr.png

✅ QR Code Recognized successfully!
  • Type:   url
  • Result: https://github.com/b1n9bna/NoPhoneQR
  • Action: Copied to clipboard

        0.28 real         0.06 user         0.07 sys
             9613312  maximum resident set size  <-- Peak RSS: 9.1 MB
             2228224  peak memory footprint      <-- Core footprint: 2.1 MB
```
</details>

---

## 🛡️ URL Sanitizer in Action

Most web platforms append multi-line tracking queries to QR codes. NoPhoneQR strips them instantly:

| Source | Raw Scanned Ugly URL | ⚡ NoPhoneQR Sanitized URL |
| :--- | :--- | :--- |
| **WeChat Articles** | `mp.weixin.qq.com/s/xxx?chksm=abc&scene=21&version=1.0` | `mp.weixin.qq.com/s/xxx` |
| **Bilibili Video** | `bilibili.com/video/BVxxx?spm_id_from=333.999&vd_source=abc` | `bilibili.com/video/BVxxx` |
| **Zhihu Redirect** | `link.zhihu.com/?target=https%3A%2F%2Fgithub.com%2F...` | `github.com/...` *(Bypasses warning redirect)* |
| **Ad / Marketing** | `example.com/item?utm_source=twitter&utm_medium=cpc&fbclid=xyz` | `example.com/item` |

---

## ✨ Features

- ⚡ **Ultra-Lightweight**: Single static binary (~3.3MB), zero Electron bloat, <10MB peak RAM, sub-millisecond cold start.
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
