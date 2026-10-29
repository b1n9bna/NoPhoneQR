package decoder

import (
	"bytes"
	"image/png"
	"testing"

	qrcode "github.com/skip2/go-qrcode"
)

func TestDecodeImageAndBytes(t *testing.T) {
	testText := "https://github.com/b1n9bna/nophoneqr"

	// 1. Generate QR code png bytes
	pngBytes, err := qrcode.Encode(testText, qrcode.Medium, 256)
	if err != nil {
		t.Fatalf("failed to encode qr code: %v", err)
	}

	// 2. Test DecodeBytes
	decodedBytes, err := DecodeBytes(pngBytes)
	if err != nil {
		t.Fatalf("DecodeBytes failed: %v", err)
	}
	if decodedBytes != testText {
		t.Errorf("expected %q, got %q", testText, decodedBytes)
	}

	// 3. Test DecodeImage
	img, err := png.Decode(bytes.NewReader(pngBytes))
	if err != nil {
		t.Fatalf("failed to decode png into image.Image: %v", err)
	}
	decodedImg, err := DecodeImage(img)
	if err != nil {
		t.Fatalf("DecodeImage failed: %v", err)
	}
	if decodedImg != testText {
		t.Errorf("expected %q, got %q", testText, decodedImg)
	}
}
