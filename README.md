# NoPhoneQR

<p align="center">
  <a href="https://github.com/b1n9bna/NoPhoneQR/releases"><img src="https://img.shields.io/github/v/release/b1n9bna/NoPhoneQR?style=flat-square" alt="Release"></a>
  <img src="https://img.shields.io/badge/RAM_Peak-9.1_MB_(实测)-success?style=flat-square&logo=speedtest&logoColor=white" alt="Peak RAM">
  <img src="https://img.shields.io/badge/Binary_Size-3.3_MB-blue?style=flat-square&logo=files" alt="Binary Size">
  <img src="https://img.shields.io/badge/Cold_Start-<50_ms-orange?style=flat-square&logo=lightning" alt="Cold Start">
  <img src="https://img.shields.io/badge/Pure_Go-100%25_(Zero_CGO)-blueviolet?style=flat-square&logo=go" alt="Pure Go">
  <img src="https://img.shields.io/badge/Platform-macOS_|_Windows_|_Linux-brightgreen?style=flat-square" alt="Platform">
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square" alt="License"></a>
</p>

桌面端屏幕二维码框选识别与链接处理工具（Go 实现）。

在电脑浏览网页、技术文档或阅读文章时，常会遇到屏幕上的二维码。NoPhoneQR 可以直接通过快捷键或命令框选屏幕区域识别二维码，自动过滤常见的追踪参数，并在浏览器中打开或复制到剪贴板，省去掏手机扫码再把链接转发回电脑的繁琐步骤。

[简体中文](README.md) • [English](README_en.md)

---

## 主要功能

- **屏幕框选识别**：直接在屏幕上划选二维码区域即可完成识别，无需外接设备或手机辅助。
- **链接参数清理**：自动剥离常见的跟踪与推广参数（如 `utm_*`、`spm_*`、`chksm`、`scene`、`fbclid` 等），保持链接整洁。
- **重定向直达**：对知乎、QQ 等常见的中转拦截跳转链接进行解析，直接打开真实目标网址。
- **多类型内容识别**：
  - **网页链接**：清理跟踪参数后存入剪贴板，并默认在浏览器中打开。
  - **Wi-Fi 配置**：解析网络 SSID，并将密码提取复制到剪贴板。
  - **双重认证（2FA）**：解析 `otpauth://` 格式二维码，提取 Secret 密钥。
  - **纯文本信息**：识别后直接复制到系统剪贴板。
- **轻量独立**：纯 Go 编写（基于 `zxing-cpp` 纯 Go 移植版），无 CGO 依赖，编译为单个独立二进制，无需安装 Python、OpenCV 或 Node.js 运行环境。

---

## 性能与资源实测

在 macOS 14 (Sonoma) 系统上通过 `/usr/bin/time -l` 测试完整识别流程：

| 指标维度 | NoPhoneQR (Go 原生) | 常见桌面截屏扫码方案 | 说明 |
| :--- | :--- | :--- | :--- |
| **二进制体积** | **3.3 MB** (Windows 4.7MB) | 150 MB ~ 300 MB | 单文件解压即用，体积小巧 |
| **峰值物理内存 (RSS)** | **9.1 MB** | 200 MB ~ 400 MB | 仅在框选识别瞬间分配必要内存 |
| **空闲常驻内存** | **0 MB** | 50 MB ~ 150 MB | 即用即走，非后台常驻进程 |
| **启动与识别耗时** | **< 50 ms** | 1.0 s ~ 3.0 s | 呼出后毫秒级就绪 |
| **运行依赖** | **零外部依赖** | Node.js / Python / OpenCV 等 | 纯 Go 编写，无 CGO 依赖 |

<details>
<summary>点击查看 macOS 实测控制台输出</summary>

```bash
$ /usr/bin/time -l nophoneqr -no-open test_qr.png

✅ QR Code Recognized successfully!
  • Type:   url
  • Result: https://github.com/b1n9bna/NoPhoneQR
  • Action: Copied to clipboard

        0.28 real         0.06 user         0.07 sys
             9613312  maximum resident set size  # 峰值物理内存约 9.1 MB
             2228224  peak memory footprint      # 核心 footprint 约 2.1 MB
```
</details>

---

## 快速上手

### 方式一：下载预编译可执行文件

从 [Releases 页面](https://github.com/b1n9bna/NoPhoneQR/releases) 下载适合当前系统的单文件即可使用：

- **macOS (Apple Silicon)**: `nophoneqr-darwin-arm64`
- **macOS (Intel)**: `nophoneqr-darwin-amd64`
- **Windows (x64)**: `nophoneqr-windows-amd64.exe`
- **Linux (x64)**: `nophoneqr-linux-amd64`

*(macOS / Linux 用户下载后赋予执行权限：`chmod +x nophoneqr*`，并建议移动到 PATH 路径如 `/usr/local/bin/nophoneqr`)*

### 方式二：通过 Go 安装

若本地已安装 Go（1.22+）：
```bash
go install github.com/b1n9bna/nophoneqr@latest
```

---

## 使用说明

### 1. 屏幕框选识别（默认）
直接运行命令（或通过全局快捷键呼出）：
```bash
nophoneqr
```
鼠标指针会变为选区十字光标，框选屏幕上的二维码即可自动识别；按 `ESC` 可取消框选。

### 2. 识别本地图片
```bash
nophoneqr /path/to/qrcode.png
```

### 3. 生成二维码（辅助功能）
在终端中以字符形式打印二维码，或保存为 PNG 图片：
```bash
# 终端控制台打印
nophoneqr -g "https://example.com"

# 保存为图片
nophoneqr -g "https://example.com" -o qr.png
```

### 常用参数说明

| 参数 | 说明 |
| :--- | :--- |
| `nophoneqr` | 默认交互式截屏识别 |
| `nophoneqr <image>` | 识别指定图片文件中的二维码 |
| `-no-open` | 识别出 URL 后仅复制到剪贴板，不自动在浏览器打开 |
| `-g <content>` | 生成字符二维码并输出到终端 |
| `-o <file.png>` | 与 `-g` 搭配使用，将二维码保存为图片 |
| `-v` | 查看版本信息 |

---

## 推荐：配置全局快捷键

将 `nophoneqr` 绑定到全局快捷键（如 `Cmd + Shift + Q` 或 `Alt + Q`）使用体验更加方便：

### macOS
- **快捷指令（Shortcuts）**：
  1. 打开系统自带的“快捷指令”应用，新建一个指令；
  2. 添加“运行 Shell 脚本”动作，内容填入 `/usr/local/bin/nophoneqr`；
  3. 在右侧边栏设置快捷键（例如 `Cmd + Shift + Q`）。
- **Raycast / Alfred**：
  新建一个 Script Command 直接调用 `nophoneqr`，并分配快捷键。

### Windows
1. 将 `nophoneqr.exe` 放置在固定目录（如 `C:\Program Files\NoPhoneQR\` 或用户目录下）；
2. 创建桌面快捷方式，右键点击快捷方式选择“属性”；
3. 在“快捷键”输入框中按下想要的组合键（如 `Ctrl + Alt + Q`），点击确定。

---

## 本地编译

```bash
git clone https://github.com/b1n9bna/NoPhoneQR.git
cd NoPhoneQR

# 本地编译
go build -o nophoneqr .

# 运行测试
go test -v ./...
```

---

## 许可证

本项目采用 [MIT License](LICENSE) 授权。
