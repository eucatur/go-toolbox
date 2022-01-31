package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHumanizedDuration(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	assert := assert.New(t)

	t1, _ := time.Parse(layout, "2019-11-01 01:02:00")
	t2, _ := time.Parse(layout, "2019-11-02 16:06:00")

	expected := "1 dia 15 horas 4 minutos"

	actual := HumanizedDuration(t2.Sub(t1))
	assert.Equal(expected, actual)

	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-03 02:03:00")
	expected = "2 dias 1 hora 1 minuto"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-01 03:04:00")
	expected = "2 horas 2 minutos"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-01 01:05:00")
	expected = "3 minutos"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-01 04:02:00")
	expected = "3 horas"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)
}

func TestGetTryParseDate(t *testing.T) {

	assert := assert.New(t)

	// Layouts for tests
	layouts := map[string]string{
		"EUAlayoutFullDateTimeZone": "2006-01-02T15:04:05Z",
		"EUAlayoutFullDateTime":     "2006-01-02 15:04:05",
		"EUAlayoutFullDate":         "2006-01-02",
		"BRlayoutFullDateTimeZone":  "02/01/2006T15:04:05Z",
		"BRlayoutFullDateTime":      "02/01/2006 15:04:05",
		"BRlayoutFullDate":          "02/01/2006",
	}

	now := time.Now()

	assertDateTimeFormatEUAFullDateTimeZone := now.Format(layouts["EUAlayoutFullDateTimeZone"])
	assertDateTimeFormatEUAFullDateTime := now.Format(layouts["EUAlayoutFullDateTime"])
	assertDateTimeFormatEUAFullDate := now.Format(layouts["EUAlayoutFullDate"])
	assertDateTimeFormatBRFullDateTimeZone := now.Format(layouts["BRlayoutFullDateTimeZone"])
	assertDateTimeFormatBRFullDateTime := now.Format(layouts["BRlayoutFullDateTime"])
	assertDateTimeFormatBRFullDate := now.Format(layouts["BRlayoutFullDate"])

	// EUA
	parsed := GetTryParseDate(now.String(), "").Format("2006-01-02T15:04:05Z")

	assert.Equal(assertDateTimeFormatEUAFullDateTimeZone, parsed)

	parsed = GetTryParseDate(now.String(), "").Format("2006-01-02 15:04:05")

	assert.Equal(assertDateTimeFormatEUAFullDateTime, parsed)

	parsed = GetTryParseDate(now.String(), "").Format("2006-01-02")

	assert.Equal(assertDateTimeFormatEUAFullDate, parsed)

	// BR
	parsed = GetTryParseDate(now.String(), "").Format("02/01/2006T15:04:05Z")

	assert.Equal(assertDateTimeFormatBRFullDateTimeZone, parsed)

	parsed = GetTryParseDate(now.String(), "").Format("02/01/2006 15:04:05")

	assert.Equal(assertDateTimeFormatBRFullDateTime, parsed)

	parsed = GetTryParseDate(now.String(), "").Format("02/01/2006")

	assert.Equal(assertDateTimeFormatBRFullDate, parsed)

	// Parsed with value setted

	valueDateTime := "20/07/2020 20:05:06"

	parseTestValueDateTime, err := time.Parse(layouts["BRlayoutFullDateTime"], valueDateTime)

	if err != nil {
		t.Error(err)
		return
	}

	assertDateTimeFormatEUAFullDateTimeZone = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTimeZone"])
	assertDateTimeFormatEUAFullDateTime = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTime"])
	assertDateTimeFormatEUAFullDate = parseTestValueDateTime.Format(layouts["EUAlayoutFullDate"])
	assertDateTimeFormatBRFullDateTimeZone = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTimeZone"])
	assertDateTimeFormatBRFullDateTime = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTime"])
	assertDateTimeFormatBRFullDate = parseTestValueDateTime.Format(layouts["BRlayoutFullDate"])

	// EUA
	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02T15:04:05Z")

	assert.Equal(assertDateTimeFormatEUAFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02 15:04:05")

	assert.Equal(assertDateTimeFormatEUAFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02")

	assert.Equal(assertDateTimeFormatEUAFullDate, parsed)

	// BR
	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006T15:04:05Z")

	assert.Equal(assertDateTimeFormatBRFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006 15:04:05")

	assert.Equal(assertDateTimeFormatBRFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006")

	assert.Equal(assertDateTimeFormatBRFullDate, parsed)

	// Parsed with short value setted

	valueDateTime = "20/07/2020"

	parseTestValueDateTime, err = time.Parse(layouts["BRlayoutFullDate"], valueDateTime)

	if err != nil {
		t.Error(err)
		return
	}

	assertDateTimeFormatEUAFullDateTimeZone = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTimeZone"])
	assertDateTimeFormatEUAFullDateTime = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTime"])
	assertDateTimeFormatEUAFullDate = parseTestValueDateTime.Format(layouts["EUAlayoutFullDate"])
	assertDateTimeFormatBRFullDateTimeZone = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTimeZone"])
	assertDateTimeFormatBRFullDateTime = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTime"])
	assertDateTimeFormatBRFullDate = parseTestValueDateTime.Format(layouts["BRlayoutFullDate"])

	// EUA
	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02T15:04:05Z")

	assert.Equal(assertDateTimeFormatEUAFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02 15:04:05")

	assert.Equal(assertDateTimeFormatEUAFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02")

	assert.Equal(assertDateTimeFormatEUAFullDate, parsed)

	// BR
	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006T15:04:05Z")

	assert.Equal(assertDateTimeFormatBRFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006 15:04:05")

	assert.Equal(assertDateTimeFormatBRFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006")

	assert.Equal(assertDateTimeFormatBRFullDate, parsed)

	// Parsed with value setted

	valueDateTime = "2020-07-20 20:05:06"

	parseTestValueDateTime, err = time.Parse(layouts["EUAlayoutFullDateTime"], valueDateTime)

	if err != nil {
		t.Error(err)
		return
	}

	assertDateTimeFormatEUAFullDateTimeZone = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTimeZone"])
	assertDateTimeFormatEUAFullDateTime = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTime"])
	assertDateTimeFormatEUAFullDate = parseTestValueDateTime.Format(layouts["EUAlayoutFullDate"])
	assertDateTimeFormatBRFullDateTimeZone = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTimeZone"])
	assertDateTimeFormatBRFullDateTime = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTime"])
	assertDateTimeFormatBRFullDate = parseTestValueDateTime.Format(layouts["BRlayoutFullDate"])

	// EUA
	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02T15:04:05Z")

	assert.Equal(assertDateTimeFormatEUAFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02 15:04:05")

	assert.Equal(assertDateTimeFormatEUAFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02")

	assert.Equal(assertDateTimeFormatEUAFullDate, parsed)

	// BR
	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006T15:04:05Z")

	assert.Equal(assertDateTimeFormatBRFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006 15:04:05")

	assert.Equal(assertDateTimeFormatBRFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006")

	assert.Equal(assertDateTimeFormatBRFullDate, parsed)

	// Parsed with short value setted

	valueDateTime = "2020-07-20"

	parseTestValueDateTime, err = time.Parse(layouts["EUAlayoutFullDate"], valueDateTime)

	if err != nil {
		t.Error(err)
		return
	}

	assertDateTimeFormatEUAFullDateTimeZone = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTimeZone"])
	assertDateTimeFormatEUAFullDateTime = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTime"])
	assertDateTimeFormatEUAFullDate = parseTestValueDateTime.Format(layouts["EUAlayoutFullDate"])
	assertDateTimeFormatBRFullDateTimeZone = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTimeZone"])
	assertDateTimeFormatBRFullDateTime = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTime"])
	assertDateTimeFormatBRFullDate = parseTestValueDateTime.Format(layouts["BRlayoutFullDate"])

	// EUA
	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02T15:04:05Z")

	assert.Equal(assertDateTimeFormatEUAFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02 15:04:05")

	assert.Equal(assertDateTimeFormatEUAFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("2006-01-02")

	assert.Equal(assertDateTimeFormatEUAFullDate, parsed)

	// BR
	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006T15:04:05Z")

	assert.Equal(assertDateTimeFormatBRFullDateTimeZone, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006 15:04:05")

	assert.Equal(assertDateTimeFormatBRFullDateTime, parsed)

	parsed = GetTryParseDate(valueDateTime, "").Format("02/01/2006")

	assert.Equal(assertDateTimeFormatBRFullDate, parsed)

	// ---------------- tests with layout setted -------------------------
	valueDateTime = "20/07/2020 20:05:06"

	parseTestValueDateTime, err = time.Parse(layouts["BRlayoutFullDateTime"], valueDateTime)

	if err != nil {
		t.Error(err)
		return
	}

	assertDateTimeFormatEUAFullDateTimeZone = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTimeZone"])
	assertDateTimeFormatEUAFullDateTime = parseTestValueDateTime.Format(layouts["EUAlayoutFullDateTime"])
	assertDateTimeFormatEUAFullDate = parseTestValueDateTime.Format(layouts["EUAlayoutFullDate"])
	assertDateTimeFormatBRFullDateTimeZone = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTimeZone"])
	assertDateTimeFormatBRFullDateTime = parseTestValueDateTime.Format(layouts["BRlayoutFullDateTime"])
	assertDateTimeFormatBRFullDate = parseTestValueDateTime.Format(layouts["BRlayoutFullDate"])

	// EUA
	parsed = GetTryParseDate("2020-07-20T20:05:06Z", "2006-01-02T15:04:05Z").Format("2006-01-02T15:04:05Z")

	assert.Equal(assertDateTimeFormatEUAFullDateTimeZone, parsed)

	parsed = GetTryParseDate("2020-07-20 20:05:06", "2006-01-02 15:04:05").Format("2006-01-02 15:04:05")

	assert.Equal(assertDateTimeFormatEUAFullDateTime, parsed)

	parsed = GetTryParseDate("2020-07-20", "2006-01-02").Format("2006-01-02")

	assert.Equal(assertDateTimeFormatEUAFullDate, parsed)

	// BR
	parsed = GetTryParseDate("20/07/2020T20:05:06Z", "02/01/2006T15:04:05Z").Format("02/01/2006T15:04:05Z")

	assert.Equal(assertDateTimeFormatBRFullDateTimeZone, parsed)

	parsed = GetTryParseDate("20/07/2020 20:05:06", "02/01/2006 15:04:05").Format("02/01/2006 15:04:05")

	assert.Equal(assertDateTimeFormatBRFullDateTime, parsed)

	parsed = GetTryParseDate("20/07/2020", "02/01/2006").Format("02/01/2006")

	assert.Equal(assertDateTimeFormatBRFullDate, parsed)

	parsed = GetTryParseDate("20210622111237", "").Format("2006-01-02 15:04:05")

	assert.Equal("2021-06-22 11:12:37", parsed)

	parsed = GetTryParseDate("20210622112654", "").Format("2006-01-02 15:04:05")

	assert.Equal("2021-06-22 11:26:54", parsed)

	parsed = GetTryParseDate("2022-01-31T10:23:54+05:00", "").Format("2006-01-02 15:04:05")

	assert.Equal("2022-01-31 10:23:54", parsed)
}
