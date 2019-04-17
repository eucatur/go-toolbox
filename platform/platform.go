package platform

import (
	"regexp"
	"strings"
)

const (
	Platform   = "platform"
	Mobile     = "Mobile"
	Aplicativo = "Aplicativo"
	Desktop    = "Desktop"
)

// FromUserAgent ...
func FromUserAgent(userAgent string) (p string, err error) {
	userAgent = strings.ToLower(userAgent)

	regexAplicativo := regexp.MustCompile(`(; wv\)|aplicativo)`)
	if regexAplicativo.MatchString(userAgent) {
		return Aplicativo, nil
	}

	rgxMobile := regexp.MustCompile(`(android|avantgo|blackberry|bolt|boost|cricket|docomo|fone|hiptop|mini|mobi|palm|phone|pie|up\.browser|up\.link|webos|wos)`)
	if rgxMobile.MatchString(userAgent) {
		return Mobile, nil
	}

	return Desktop, nil
}
