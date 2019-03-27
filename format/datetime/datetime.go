package datetime

import (
	"time"
)

func EUAtoBR(str string) string {

	if d, e := time.Parse("2006-01-02 15:04:05", str); e == nil {
		return d.Format("02/01/2006 15:04:05")
	}

	return ""
}

func EUAtoBRShort(str string) string {

	if d, e := time.Parse("2006-01-02 15:04:05", str); e == nil {
		return d.Format("02/01/2006 15:04")
	}

	return ""
}

func AsBRShort(dt time.Time) string {
	return dt.Format("02/01/2006 15:04")
}
