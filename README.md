<p align="center">
  <h1 align="center">📱❌ NoPhoneQR</h1>
  <p align="center">
    <strong>Stop pointing your phone at your computer screen.</strong><br>
    告别对着电脑屏幕举手机扫码的尴尬 —— 屏幕二维码免掏手机一键解析与跳转。
  </p>
  <p align="center">
    <a href="https://github.com/b1n9bna/NoPhoneQR/releases"><img src="https://img.shields.io/github/v/release/b1n9bna/NoPhoneQR?style=flat-square" alt="Release"></a>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go" alt="Go"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square" alt="License"></a>
    <img src="https://img.shields.io/badge/Platform-macOS%20%7C%20Windows%20%7C%20Linux-brightgreen?style=flat-square" alt="Platform">
  </p>
</p>

---

## 💡 为什么需要 NoPhoneQR？(Why NoPhoneQR?)

我们在电脑上办公或浏览网页时，经常遇到屏幕上的二维码：
- 网页、文档、GitHub 项目中的链接二维码
- 微信公众号、掘金、知乎文章中的参考资料或原文链接
- 2FA 两步验证绑定的 `otpauth://` 密钥二维码
- 咖啡厅、办公区 Wi-Fi 分享二维码

### 传统做法（繁琐的 7 步）：
1. 摸口袋掏出手机 📱
2. 面容/指纹解锁手机
3. 打开微信 / 相机
4. 身体后倾、手举着手机对准电脑屏幕扫码 📷
5. 在手机上打开页面并复制链接
6. 发送给电脑端微信“文件传输助手”
7. 在电脑上点开链接……

### 使用 NoPhoneQR（1 步搞定）：
> **按下一个快捷键（或敲一下命令） ➜ 鼠标框选屏幕上的二维码 ➜ 浏览器立即在新标签页打开干净链接，同时已自动复制到剪贴板！** 🎉

---

## ✨ 核心特性 (Features)

- ⚡ **极致轻量**：单文件绿色二进制包（约 4MB），内存占用 < 15MB，无 Electron 臃肿依赖，毫秒级冷启动。
- 🛡️ **智能链接净化 (URL Sanitizer)**：自动剥离 `utm_*`、`spm_*`、`chksm`、`scene`、`fbclid` 等恶心跟踪尾巴；自动展开知乎、QQ 等中转重定向，还你纯净短链接。
- 🧠 **多类型智能识别 (Smart Payloads)**：
  - **网页链接**：自动清理参数 + 唤起系统默认浏览器打开 + 存入剪贴板。
  - **Wi-Fi 配置**：自动解析 SSID，提取密码直接复制到剪贴板。
  - **2FA 密钥**：自动解析 `otpauth://` 账户与 Secret，免手动敲长串字符。
  - **普通文本 / 订单号**：直接复制到剪贴板，附带系统原生气泡提示。
- 🖥️ **全平台支持**：
  - **macOS**：原生接入 macOS 交互式准心选区，Retina 视网膜高清屏自适应。
  - **Windows**：原生无缝支持 Windows 10/11 Snipping 机制与剪贴板极速解析。
  - **Linux**：支持 Wayland (grim/slurp) 及 X11 (maim/scrot)。

---

## 🚀 快速上手 (Quick Start)

### 1. 下载预编译文件 (Releases)
直接在 [Releases 页面](https://github.com/b1n9bna/NoPhoneQR/releases) 下载对应系统的单文件执行档即可使用：
- macOS Apple Silicon (M1/M2/M3/M4): `nophoneqr-darwin-arm64`
- macOS Intel: `nophoneqr-darwin-amd64`
- Windows (x64): `nophoneqr-windows-amd64.exe`
- Linux (x64): `nophoneqr-linux-amd64`

### 2. 通过 Go 安装
如果你已安装 Go 环境：
```bash
go install github.com/b1n9bna/nophoneqr@latest
```

---

## 💻 使用方法 (Usage)

### 场景一：直接在屏幕上框选扫码（最常用）
在终端中输入，或绑定快捷键：
```bash
nophoneqr
```
鼠标指针会立即变成选区十字标，拖动鼠标框选屏幕上的二维码：
- 识别到链接：系统默认浏览器自动打开，并复制到剪贴板。
- 按 `ESC` 可随时取消框选。

### 场景二：解析本地已有的图片文件
```bash
nophoneqr screenshot.png
# 或
nophoneqr -f /path/to/qrcode.jpg
```

### 场景三：生成二维码（附赠工具）
在终端直接打印出 ASCII 二维码，或保存为 PNG：
```bash
# 终端中字符渲染
nophoneqr -g "https://github.com/b1n9bna/NoPhoneQR"

# 保存为高清 PNG 图片
nophoneqr -g "https://example.com" -o qr.png
```

### 常用选项
| 参数 | 说明 |
| :--- | :--- |
| `nophoneqr` | 默认交互式截屏识别 |
| `nophoneqr <image>` | 识别指定图片路径中的二维码 |
| `-no-open` | 识别链接后仅复制到剪贴板，不自动弹出浏览器 |
| `-g <content>` | 生成终端 ASCII 字符二维码 |
| `-o <file.png>` | 与 `-g` 搭配，将生成的二维码保存为 PNG |
| `-v` | 查看当前版本号 |

---

## ⌨️ 推荐：配置全局快捷键（打造极致体验）

建议将 `nophoneqr` 绑定一个全局快捷键（如 `Cmd + Shift + Q` 或 `Alt + Q`）：

### 🍏 macOS 用户推荐配置：
1. **Raycast / Alfred 用户**：
   - 在 Raycast 中添加一条 `Script Command`，直接调用 `nophoneqr`，并绑定快捷键 `Cmd + Shift + Q`。
2. **系统原生「快捷指令 (Shortcuts)」**：
   - 打开系统自带的 **快捷指令** App，新建快捷指令。
   - 添加「运行 Shell 脚本」动作，填入 `/usr/local/bin/nophoneqr`。
   - 在右侧边栏给该指令指定键盘快捷键（例如 `Cmd + Shift + Q`）。

### 🪟 Windows 用户推荐配置：
1. 下载 `nophoneqr-windows-amd64.exe`，重命名为 `nophoneqr.exe` 并放入如 `C:\Windows\` 或任意加入 PATH 的目录。
2. 右键创建桌面快捷方式，在快捷方式属性的「快捷键」一栏按下 `Ctrl + Alt + Q`，即可随时一键呼出！

---

## 🛠️ 本地编译 (Build from Source)

```bash
git clone https://github.com/b1n9bna/NoPhoneQR.git
cd NoPhoneQR

# 本地直接编译
go build -o nophoneqr .

# 运行单元测试
go test -v ./...

# 一键跨平台交叉编译 (macOS, Windows, Linux)
make cross-build
```

---

## 📄 开源许可 (License)

本项目遵循 [MIT License](LICENSE)。欢迎提交 Issue 与 Pull Request！
