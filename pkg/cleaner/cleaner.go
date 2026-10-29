package cleaner

import (
	"net/url"
	"strings"
)

// trackingParams contains common tracking query parameter prefixes and exact matches.
var trackingParamPrefixes = []string{
	"utm_",
	"spm_",
}

var trackingExactParams = map[string]bool{
	// General tracking & ads
	"fbclid":      true,
	"gclid":       true,
	"msclkid":     true,
	"yclid":       true,
	"ttclid":      true,
	"twclid":      true,
	"igshid":      true,
	"si":          true,
	"ref":         true,
	"source":      true,
	"_hsenc":      true,
	"_hsmi":       true,
	"mc_cid":      true,
	"mc_eid":      true,
	// WeChat
	"chksm":       true,
	"scene":       true,
	"subscene":    true,
	"sessionid":   true,
	"sender":      true,
	"clicktime":   true,
	"enterid":     true,
	"ascene":      true,
	"devicetype":  true,
	"version":     true,
	"nettype":     true,
	"lang":        true,
	"pass_ticket": true,
	"wx_header":   true,
	// Bilibili
	"from_source":      true,
	"from_spmid":       true,
	"vd_source":        true,
	"share_source":     true,
	"share_medium":     true,
	"share_plat":       true,
	"share_session_id": true,
	"share_tag":        true,
	// Xiaohongshu / Weibo
	"xhsshare": true,
	"appuid":   true,
	"apptime":  true,
	"wm":       true,
}

// CleanURL sanitizes a URL by unwrapping redirects and stripping tracking parameters.
func CleanURL(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return ""
	}

	// Unwrap known redirect wrappers first
	unwrapped := unwrapRedirect(rawURL)
	if unwrapped != "" {
		rawURL = unwrapped
	}

	u, err := url.Parse(rawURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return rawURL
	}

	q := u.Query()
	modified := false

	for param := range q {
		lower := strings.ToLower(param)
		if trackingExactParams[lower] || hasTrackingPrefix(lower) {
			q.Del(param)
			modified = true
		}
	}

	if modified {
		u.RawQuery = q.Encode()
	}

	return u.String()
}

func hasTrackingPrefix(param string) bool {
	for _, prefix := range trackingParamPrefixes {
		if strings.HasPrefix(param, prefix) {
			return true
		}
	}
	return false
}

// unwrapRedirect extracts the target destination URL from known intermediate redirects.
func unwrapRedirect(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	host := strings.ToLower(u.Host)

	// Zhihu outbound link wrapper: link.zhihu.com/?target=http%3A%2F%2F...
	if strings.Contains(host, "zhihu.com") && strings.HasPrefix(u.Path, "/question/") == false {
		target := u.Query().Get("target")
		if target != "" {
			if decoded, err := url.QueryUnescape(target); err == nil && isHTTP(decoded) {
				return decoded
			}
		}
	}

	// QQ / WeChat outbound warning wrappers
	if strings.Contains(host, "qq.com") {
		for _, key := range []string{"url", "pfurl", "target"} {
			if val := u.Query().Get(key); val != "" {
				if decoded, err := url.QueryUnescape(val); err == nil && isHTTP(decoded) {
					return decoded
				}
			}
		}
	}

	return ""
}

func isHTTP(raw string) bool {
	return strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://")
}
