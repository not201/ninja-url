package utils

import (
	"net/url"
	"strings"
)

func IsValidUrl(urlStr string) bool {
	urlStr = strings.TrimSpace(urlStr)

	if urlStr == "" {
		return false
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}
