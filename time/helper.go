package time

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/eucatur/go-toolbox/check"
)

func HumanizedDuration(d time.Duration) string {
	hours, _ := math.Modf(d.Hours())
	minutes := d.Minutes() - (hours * 60)
	days, _ := math.Modf(hours / 24)

	days = math.Abs(days)
	hours = math.Abs(hours)
	minutes = math.Abs(minutes)

	hours = hours - (days * 24)

	iDays := int(days)
	iHours := int(hours)
	iMinutes := int(minutes)

	output := []string{}

	if iDays > 0 {
		output = append(output, fmt.Sprint(iDays))
		output = append(output, check.If(iDays == 1, "dia", "dias").(string))
	}

	if iHours > 0 {
		output = append(output, fmt.Sprint(iHours))
		output = append(output, check.If(iHours == 1, "hora", "horas").(string))
	}

	if iMinutes > 0 {
		output = append(output, fmt.Sprint(iMinutes))
		output = append(output, check.If(iMinutes == 1, "minuto", "minutos").(string))
	}

	return strings.Join(output, " ")
}
