package dispatcher

import (
	"testing"
)

func TestDispatch(t *testing.T) {
	// 1. Test URL (autoOpen = false for tests so browser doesn't pop up during unit tests)
	res, err := Dispatch("https://mp.weixin.qq.com/s/test?chksm=123&scene=1", false)
	if err != nil {
		t.Fatalf("Dispatch URL failed: %v", err)
	}
	if res.Type != "url" {
		t.Errorf("expected type 'url', got %q", res.Type)
	}
	if res.Processed != "https://mp.weixin.qq.com/s/test" {
		t.Errorf("expected cleaned URL, got %q", res.Processed)
	}

	// 2. Test Wi-Fi
	res, err = Dispatch("WIFI:S:Office-5G;T:WPA;P:SuperSecretPass123;;", false)
	if err != nil {
		t.Fatalf("Dispatch Wi-Fi failed: %v", err)
	}
	if res.Type != "wifi" {
		t.Errorf("expected type 'wifi', got %q", res.Type)
	}
	if res.CopiedText != "SuperSecretPass123" {
		t.Errorf("expected copied password 'SuperSecretPass123', got %q", res.CopiedText)
	}

	// 3. Test OTP 2FA
	res, err = Dispatch("otpauth://totp/GitHub:b1n9?secret=JBSWY3DPEHPK3PXP&issuer=GitHub", false)
	if err != nil {
		t.Fatalf("Dispatch OTP failed: %v", err)
	}
	if res.Type != "otp" {
		t.Errorf("expected type 'otp', got %q", res.Type)
	}
	if res.CopiedText != "JBSWY3DPEHPK3PXP" {
		t.Errorf("expected secret copied, got %q", res.CopiedText)
	}

	// 4. Test Plain text
	res, err = Dispatch("Hello, World!", false)
	if err != nil {
		t.Fatalf("Dispatch text failed: %v", err)
	}
	if res.Type != "text" || res.CopiedText != "Hello, World!" {
		t.Errorf("unexpected result for plain text: %+v", res)
	}
}
