package time

import (
	"time"
)

// WeekDayPTBr Returns the day of the week in Brazilian Portuguese
func WeekDayPTBr(date time.Time) (weekDay string) {
	switch date.Weekday() {
	case time.Sunday:
		weekDay = "Domingo"
	case time.Monday:
		weekDay = "Segunda-feira"
	case time.Tuesday:
		weekDay = "Terça-feira"
	case time.Wednesday:
		weekDay = "Quarta-feira"
	case time.Thursday:
		weekDay = "Quinta-feira"
	case time.Friday:
		weekDay = "Sexta-feira"
	case time.Saturday:
		weekDay = "Sábado"
	default:
		weekDay = ""
	}

	return
}
