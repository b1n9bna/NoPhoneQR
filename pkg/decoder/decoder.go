package decoder

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// DecodeImage parses a QR code from an in-memory image.Image object.
func DecodeImage(img image.Image) (string, error) {
	if img == nil {
		return "", fmt.Errorf("nil image provided")
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", fmt.Errorf("failed to create binary bitmap: %w", err)
	}

	qrReader := qrcode.NewQRCodeReader()

	// 1st attempt: standard fast decode
	result, err := qrReader.Decode(bmp, nil)
	if err == nil && result != nil {
		return result.GetText(), nil
	}

	// 2nd attempt: TRY_HARDER hint
	hints := map[gozxing.DecodeHintType]interface{}{
		gozxing.DecodeHintType_TRY_HARDER: true,
	}
	result, err = qrReader.Decode(bmp, hints)
	if err == nil && result != nil {
		return result.GetText(), nil
	}

	return "", fmt.Errorf("no QR code found in image")
}

// DecodeFile reads an image file from disk and parses any QR code contained within.
func DecodeFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", fmt.Errorf("failed to decode image format: %w", err)
	}

	return DecodeImage(img)
}

// DecodeBytes parses a QR code from image bytes (e.g. from clipboard or memory buffer).
func DecodeBytes(data []byte) (string, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to decode image data: %w", err)
	}
	return DecodeImage(img)
}
