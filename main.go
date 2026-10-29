package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/b1n9bna/nophoneqr/pkg/capture"
	"github.com/b1n9bna/nophoneqr/pkg/decoder"
	"github.com/b1n9bna/nophoneqr/pkg/dispatcher"
	"github.com/b1n9bna/nophoneqr/pkg/notify"
	qrcode "github.com/skip2/go-qrcode"
)

const (
	Version = "1.0.0"
	Banner  = `
 _   _       ____  _                      ___  ____  
| \ | | ___ |  _ \| |__   ___  _ __   ___ / _ \|  _ \ 
|  \| |/ _ \| |_) | '_ \ / _ \| '_ \ / _ \ | | | |_) |
| |\  | (_) |  __/| | | | (_) | | | |  __/ |_| |  _ < 
|_| \_|\___/|_|   |_| |_|\___/|_| |_|\___|\__\_\|_| \_\
  Stop pointing your phone at your computer screen!
`
)

func main() {
	filePath := flag.String("f", "", "Path to local image file to decode")
	genText := flag.String("g", "", "Generate a QR code and print to terminal")
	outFile := flag.String("o", "", "Save generated QR code to specified PNG file")
	noOpen := flag.Bool("no-open", false, "Do not auto-open links in browser (only copy to clipboard)")
	showVersion := flag.Bool("v", false, "Show version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\nNoPhoneQR v%s - 免掏手机屏幕二维码一键解析与跳转器\n\n", Banner, Version)
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  nophoneqr                     Trigger screen capture & decode instantly\n")
		fmt.Fprintf(os.Stderr, "  nophoneqr [image_path]        Decode QR code from image file\n")
		fmt.Fprintf(os.Stderr, "  nophoneqr -g <text/url>       Generate ASCII QR code in terminal\n")
		fmt.Fprintf(os.Stderr, "  nophoneqr -g <text> -o qr.png Generate QR code and save to PNG\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("NoPhoneQR v%s (https://github.com/b1n9bna/NoPhoneQR)\n", Version)
		return
	}

	// Mode 1: Generate QR Code
	if *genText != "" {
		handleGenerate(*genText, *outFile)
		return
	}

	// Mode 2: Decode from file argument (positional or -f flag)
	targetFile := *filePath
	if targetFile == "" && flag.NArg() > 0 {
		targetFile = flag.Arg(0)
	}

	var rawQR string
	var err error

	if targetFile != "" {
		fmt.Printf("🔍 Scanning image file: %s ...\n", targetFile)
		rawQR, err = decoder.DecodeFile(targetFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to decode QR code: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Mode 3: Interactive Screen Capture
		fmt.Println("📸 Select a region with a QR code on your screen (press ESC to cancel)...")
		var img image.Image
		img, err = capture.Snip()
		if err != nil {
			if err == capture.ErrCanceled {
				fmt.Println("Selection canceled.")
				return
			}
			fmt.Fprintf(os.Stderr, "❌ Screen capture error: %v\n", err)
			os.Exit(1)
		}

		rawQR, err = decoder.DecodeImage(img)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ 未识别到二维码: %v\n", err)
			_ = notify.Send("NoPhoneQR", "未在选取区域识别到有效二维码，请重新框选")
			os.Exit(1)
		}
	}

	// Dispatch decoded content
	autoOpen := !*noOpen
	res, err := dispatcher.Dispatch(rawQR, autoOpen)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to process QR payload: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ QR Code Recognized successfully!")
	fmt.Printf("  • Type:   %s\n", res.Type)
	fmt.Printf("  • Result: %s\n", res.Processed)
	fmt.Printf("  • Action: %s\n", res.ActionTaken)
}

func handleGenerate(content, outFile string) {
	if outFile != "" {
		err := qrcode.WriteFile(content, qrcode.Medium, 256, outFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to write QR image: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ QR code saved to %s\n", outFile)
		return
	}

	// Terminal ASCII QR code
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to generate QR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(qr.ToSmallString(false))
	fmt.Printf("Content: %s\n", content)
}
