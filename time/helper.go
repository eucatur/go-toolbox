package time

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/eucatur/go-toolbox/check"
	"github.com/eucatur/go-toolbox/text"
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

// GetTryParseDate -- Obter o padrão de data e hora EUA e/ou BR
/*
dateTime: data para tentativa de conversão
layout:
"EUAlayoutFullDateTimeZone": "2006-01-02T15:04:05Z",
"EUAlayoutFullDateTime":     "2006-01-02 15:04:05",
"EUAlayoutFullDate":         "2006-01-02",
"BRlayoutFullDateTimeZone":  "02/01/2006T15:04:05Z",
"BRlayoutFullDateTime":      "02/01/2006 15:04:05",
"BRlayoutFullDate":          "02/01/2006",
Se não for definido o layout será realizado tentativas de parse com base na lista de layouts
*/
func GetTryParseDate(dateTime string, layout string) (parsedDateTime time.Time) {

	layouts := map[string]string{
		"EUAlayoutFullDateTimeZone": "2006-01-02T15:04:05Z",
		"EUAlayoutFullDateTime":     "2006-01-02 15:04:05",
		"EUAlayoutFullDate":         "2006-01-02",
		"BRlayoutFullDateTimeZone":  "02/01/2006T15:04:05Z",
		"BRlayoutFullDateTime":      "02/01/2006 15:04:05",
		"BRlayoutFullDate":          "02/01/2006",
	}

	internalLaytout := map[string]string{
		"ANSIC":                      time.ANSIC,
		"UnixDate":                   time.UnixDate,
		"RubyDate":                   time.RubyDate,
		"RFC822":                     time.RFC822,
		"RFC822Z":                    time.RFC822Z,
		"RFC850":                     time.RFC850,
		"RFC1123":                    time.RFC1123,
		"RFC1123Z":                   time.RFC1123Z,
		"RFC3339":                    time.RFC3339,
		"RFC3339Nano":                time.RFC3339Nano,
		"Kitchen":                    time.Kitchen,
		"Stamp":                      time.Stamp,
		"StampMilli":                 time.StampMilli,
		"StampMicro":                 time.StampMicro,
		"StampNano":                  time.StampNano,
		"FullLayoutNow":              "2006-01-02 15:04:05.999999 -0700 -07 m=+0.000000000",
		"LayoutTimeFullZone":         "2006-01-02 15:04:05 -0700 -07",
		"LayoutTimeZone":             "2006-01-02 15:04:05 -0700",
		"LayoutTimeUTC":              "2006-01-02 15:04:05 -0700 UTC",
		"LayoutDateTimeCombined":     "20060102150405",
		"LayoutDateTimePlusTimeZone": "2006-01-02T15:04:05+00:00",
		"LayoutDateTimeJustInfoT":    "2006-01-02T15:04:05",
	}

	if !text.StringIsEmptyOrWhiteSpace(layout) {
		setLayout, ok := layouts[layout]

		if !ok {
			setLayout = layout
		}

		getDateTime, err := time.Parse(setLayout, dateTime)

		if err != nil {
			return
		}

		parsedDateTime = getDateTime

		return
	}

	tryInternal := false

	for _, layout := range layouts {

		if !tryInternal {

			for _, internal := range internalLaytout {

				tryParseDateTime, err := time.Parse(internal, dateTime)

				if err != nil {
					continue
				}

				parsedDateTime = tryParseDateTime
				return

			}

			tryInternal = true
		}

		getDateTime, err := time.Parse(layout, dateTime)

		if err != nil {
			continue
		}

		parsedDateTime = getDateTime
		break

	}

	return
}
