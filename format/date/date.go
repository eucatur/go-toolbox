package date

import (
	"time"
)

func EUAtoBR(str string) string {

	if d, e := time.Parse("2006-01-02", str); e == nil {
		return d.Format("02/01/2006")
	}

	return ""
}

func AsEUA(dt time.Time) string {
	return dt.Format("2006-01-02")
}
