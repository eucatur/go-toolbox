package format

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`[^0-9\-\s]`)

const (
	TimeDateTime = "2006-01-02"
	TimeTime     = "15:04:05"
)

// OnlyNumbers return just de numbers on string
func OnlyNumbers(s string) string {
	s = regex.ReplaceAllString(s, "")

	if len(s) > 12 {
		s = strings.Replace(s, "-", "", -1)
	}

	return s
}
