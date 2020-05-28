package time

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/eucatur/go-toolbox/text"
)

// Timewooba ...
type Timewooba struct{ time.Time }

// MarshalJSON ...
func (t Timewooba) MarshalJSON() (bytes []byte, err error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

// UnmarshalJSON ...
func (t *Timewooba) UnmarshalJSON(bytes []byte) (err error) {
	value, err := strconv.Unquote(string(bytes))
	if err != nil {
		return
	}

	timewooba, err := ParseTimewooba(value)
	if err != nil {
		return
	}

	t.Time = timewooba.Time
	return
}

// Value ...
func (t Timewooba) Value() (value driver.Value, err error) {
	if t.IsZero() {
		return
	}

	return t.String(), nil
}

// Scan ...
func (t *Timewooba) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case Timewooba:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			var timewooba Timewooba

			timewooba, err = ParseTimewooba(v)
			if err != nil {
				return
			}

			t.Time = timewooba.Time
		}
	}

	return
}

// String ...
func (t Timewooba) String() string {
	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("/Date(%.13s%s)/", strconv.FormatInt(t.UnixNano(), 10), t.Format("-0700"))
}

// ParseTimewooba ...
func ParseTimewooba(value string) (timewooba Timewooba, err error) {
	rgx := regexp.MustCompile(`(\/Date\()(-?[0-9]{13,15})([+|-][0-9]{4})(\)\/)`)

	matches := rgx.FindStringSubmatch(value)
	if len(matches) != 5 {
		err = errors.New("Invalid date")
		return
	}

	nsecStr := text.RPad(matches[2], 19, "0")

	nsec, err := strconv.ParseInt(nsecStr, 10, 64)
	if err != nil {
		return
	}

	if nsec < 1 {
		return
	}

	timewooba.Time = time.Unix(0, nsec)

	dateTime := timewooba.Time.Format("2006-01-02 15:04:05")
	timeZone := matches[3]

	valueToParse := fmt.Sprintf("%s %s", dateTime, timeZone)

	timewooba.Time, err = time.Parse("2006-01-02 15:04:05 -0700", valueToParse)
	return
}
