package phone

import (
	"regexp"
	"strings"
)

var regexPhone = regexp.MustCompile(`[^0-9\-\s]`)

func Short(phone string) string {
	phone = regexPhone.ReplaceAllString(phone, "")

	if len(phone) > 12 {
		phone = strings.Replace(phone, "-", "", -1)
	}

	return phone
}
