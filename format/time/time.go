package time

import (
	"fmt"
	"time"
)

func HourMin(d time.Duration) string {
	return fmt.Sprintf("%02d:%02d", int64(d.Hours()), int64(d.Minutes())%60)
}

func AsDefault(dt time.Time) string {
	return dt.Format("15:04:05")
}

func ToShort(dt string) string {
	return dt[0:5]
}
