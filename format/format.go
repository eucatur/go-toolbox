package format

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// RemoveAccents is a functions for remove accents
func RemoveAccents(s string) (r string) {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	r, _, _ = transform.String(t, s)
	r = regexp.MustCompile(`[^\w]|\s`).ReplaceAllString(r, "")
	return
}
