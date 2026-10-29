package cleaner

import (
	"testing"
)

func TestCleanURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "clean simple URL without params",
			input:    "https://github.com/b1n9bna/nophoneqr",
			expected: "https://github.com/b1n9bna/nophoneqr",
		},
		{
			name:     "strip UTM parameters",
			input:    "https://example.com/article?utm_source=twitter&utm_medium=social&utm_campaign=launch&id=123",
			expected: "https://example.com/article?id=123",
		},
		{
			name:     "strip WeChat tracking params",
			input:    "https://mp.weixin.qq.com/s/abcdef?scene=21&chksm=123456&version=1.0#wechat_redirect",
			expected: "https://mp.weixin.qq.com/s/abcdef#wechat_redirect",
		},
		{
			name:     "strip Bilibili tracking params",
			input:    "https://www.bilibili.com/video/BV1xx411c7mD?spm_id_from=333.999.0.0&vd_source=abc12345",
			expected: "https://www.bilibili.com/video/BV1xx411c7mD",
		},
		{
			name:     "unwrap Zhihu intermediate redirect",
			input:    "https://link.zhihu.com/?target=https%3A%2F%2Fgithub.com%2Fb1n9bna%2Fnophoneqr",
			expected: "https://github.com/b1n9bna/nophoneqr",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CleanURL(tc.input)
			if got != tc.expected {
				t.Errorf("CleanURL(%q) = %q, expected %q", tc.input, got, tc.expected)
			}
		})
	}
}
