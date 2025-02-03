package utils

import (
	"net/url"
)

func ValidateURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}
