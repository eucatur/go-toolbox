package time

import (
	"time"
)

//FirstDayOfMonth Gets the first day of the month informed
func FirstDayOfMonth(month, year int) (firstOfMonth time.Time) {
	now := time.Now()
	currentLocation := now.Location()
	firstOfMonth = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentLocation)
	return

}

//LastDayOfMonth Gets the last day of the month informed
func LastDayOfMonth(month, year int) (lastOfMonth time.Time) {
	firstOfMonth := FirstDayOfMonth(month, year)
	lastOfMonth = firstOfMonth.AddDate(0, 1, -1)
	return
}
