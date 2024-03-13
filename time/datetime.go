package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"emperror.dev/errors"
	"github.com/jinzhu/now"
)

var CurrentTime = time.Now

const (
	timezoneformat   = "Z07:00"
	infoFormatAccept = "yyyy-MM-ddTHH:mm:ssZ"
)

var (
	errDateTimeInvalid = func(valuePerformed interface{}) error {
		return errors.Errorf("The value %v is not a datetime valid. Value date time accept %s", valuePerformed, infoFormatAccept)
	}

	isMinDateTime = func(compare time.Time) bool {
		minDateTime := time.Date(1, time.January, 1, 0, 0, 0, 0, compare.Location())
		return minDateTime.Equal(compare)
	}
)

type DateTime string

type Date DateTime

type Period []DateTime

type tpDateTime struct {
	datetime time.Time
	location *time.Location
	timezone string
	date     string
	time     string
}

func Set[T time.Time | string](value T) DateTime {

	dt := tpDateTime{}

	switch v := any(value).(type) {
	case time.Time:
		dt = tryParse(v.Format(time.RFC3339))
	default:
		dt = tryParse(fmt.Sprint(value))

	}

	if dt.datetime.IsZero() || isMinDateTime(dt.datetime) {
		return ""
	}

	dtFormated := dt.datetime.Format(time.RFC3339)

	return DateTime(dtFormated)

}

func (strdt *DateTime) SetLocation(location *time.Location) {

	dt := tryParse(string(*strdt))

	dtNewLocation := dt.datetime.In(location).Format(timezoneformat)

	newDt := Set(fmt.Sprintf("%s %s%s", dt.date, dt.time, dtNewLocation))

	*strdt = newDt

}

func (strdt *DateTime) ChangeTimezone(location *time.Location) {

	dt := tryParse(string(*strdt))

	dt = *newDateTime(dt.datetime.In(location))

	*strdt = DateTime(dt.datetime.Format(time.RFC3339))

}

func (strdt DateTime) IsValid() error {

	dt := tryParse(string(strdt))

	if dt.datetime.IsZero() {
		return fmt.Errorf("the value %#v DateTime is invalid", strdt)
	}

	return nil

}

func (strdt DateTime) String() string {

	dt := tryParse(string(strdt))

	return dt.datetime.Format(time.DateTime)
}

func (strdt DateTime) GetDateTimeWithTimezone() string {

	dt := tryParse(string(strdt))

	if dt.datetime.IsZero() || isMinDateTime(dt.datetime) {
		return ""
	}

	return dt.datetime.Format(time.RFC3339)
}

func (strdt DateTime) GetTimezone() string {

	dt := tryParse(string(strdt))

	if dt.datetime.IsZero() || isMinDateTime(dt.datetime) {
		return ""
	}

	return dt.timezone
}

func (strdt DateTime) GetTime() string {

	dt := tryParse(string(strdt))

	return dt.time
}

func (strdt DateTime) GetDate() string {

	dt := tryParse(string(strdt))

	return dt.date
}

func (strdt DateTime) GetStdTime() time.Time {

	dt := tryParse(string(strdt))

	return dt.datetime
}

func (strdt DateTime) GetLocation() time.Location {

	dt := tryParse(string(strdt))

	return *dt.location
}

func newDateTime(dt time.Time) *tpDateTime {

	return &tpDateTime{
		datetime: dt,
		location: dt.Location(),
		timezone: dt.Format(timezoneformat),
		date:     dt.Format(time.DateOnly),
		time:     dt.Format(time.TimeOnly),
	}
}

func tryParse(dateTime string) tpDateTime {

	now.TimeFormats = append(now.TimeFormats,
		fmt.Sprintf("%s %s%s", time.DateOnly, time.TimeOnly, timezoneformat),
		"02/01/2006T15:04:05Z",
		"02/01/2006 15:04:05",
		"02/01/2006",
		"2006-01-02",
		"2006-01-02 15:04:05.999999 -0700 -07 m=+0.000000000",
		"2006-01-02 15:04:05 -0700 -07",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05 -0700 UTC",
		"20060102150405",
		"2006-01-02T15:04:05+00:00",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02 15:04:05",
		"02/01/06T15:04:05Z",
		"02/01/06 15:04:05",
		"02/01/06",
		"06-01-02",
		"06-01-02 15:04:05.999999 -0700 -07 m=+0.000000000",
		"06-01-02 15:04:05 -0700 -07",
		"06-01-02 15:04:05 -0700",
		"06-01-02 15:04:05 -0700 UTC",
		"060102150405",
		"06-01-02T15:04:05+00:00",
		"06-01-02T15:04:05",
		"06-01-02T15:04",
		"06-01-02 15:04:05",
	)

	dt, err := now.Parse(dateTime)

	if err != nil {
		dt = GetTryParseDate(dateTime, "")
		if dt.IsZero() {
			return tpDateTime{}
		}
	}

	if dt.IsZero() || isMinDateTime(dt) {
		return tpDateTime{}
	}

	newDt := newDateTime(dt)

	return *newDt

}

func (dt DateTime) MarshalJSON() ([]byte, error) {

	return []byte(fmt.Sprintf(`"%s"`, dt.GetDateTimeWithTimezone())), nil

}

func (dt *DateTime) UnmarshalJSON(bytes []byte) error {

	value := string(bytes)

	value, err := strconv.Unquote(value)

	if err != nil {
		value = string(bytes)
	}

	*dt = Set(value)

	return nil

}

func (dt DateTime) Value() (driver.Value, error) {

	return dt.GetDateTimeWithTimezone(), nil

}

func (dt *DateTime) Scan(value interface{}) (err error) {

	defer func() {

		if errRecover := recover(); errRecover != nil {

			err = errDateTimeInvalid(value)

			*dt = ""

		}

	}()

	switch data := value.(type) {
	case []uint8:

		value := string(data)
		*dt = Set(value)

	case time.Time:

		*dt = Set(data)

	case string:

		*dt = Set(data)

	default:
		err = errDateTimeInvalid(value)
		*dt = Set("")
	}

	return

}

func (period Period) PeriodValid() bool {

	if len(period) != 2 {
		return false
	}

	for _, dt := range period {
		if dt.IsValid() != nil {
			return false
		}

	}

	return true

}

func (period Period) GetPeriodStartlyEnd() (from, to DateTime) {

	if !period.PeriodValid() {
		return
	}

	from = period[0]
	to = period[1]

	return

}

func (dt Date) MarshalJSON() ([]byte, error) {

	return []byte(fmt.Sprintf(`"%s"`, DateTime(dt).GetDate())), nil
}

func (dt *Date) UnmarshalJSON(bytes []byte) error {

	thisDateTime := DateTime(*dt)

	err := thisDateTime.UnmarshalJSON(bytes)

	*dt = Date(thisDateTime)

	return err

}

func (strdt Date) Value() (driver.Value, error) {

	dt := tryParse(string(strdt))

	return dt.datetime.Format(time.DateOnly), nil

}

func (dt *Date) Scan(value interface{}) (err error) {

	thisDateTime := DateTime(*dt)

	err = thisDateTime.Scan(value)

	*dt = Date(thisDateTime)

	return

}

func ParseToFullYear(year int64) int64 {

	if year >= 1000 && year <= 9999 {
		return year
	}

	currentYear := CurrentTime().Year()

	century := (currentYear / 100) * 100

	if year == 0 {
		return int64(century)
	}

	fullYear := int64(century) + year

	return fullYear
}

func (dt DateTime) GetFullYear() int64 {

	return ParseToFullYear(int64(dt.GetStdTime().Year()))

}

func (dt DateTime) GetShortYear() int64 {

	year, _ := strconv.Atoi(dt.GetStdTime().Format("06"))

	return int64(year)

}

func (dt DateTime) GetMonth() string {

	return dt.GetStdTime().Format("01")

}
