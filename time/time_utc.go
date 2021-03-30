package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimeUTCLayout = "2006-01-02T15:04:05-07:00"

type TimeUTC struct{ time.Time }

func (t TimeUTC) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *TimeUTC) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimeUTCLayout, value)
	}
	return err
}

func (t TimeUTC) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeUTC) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case TimeUTC:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimeUTCLayout, v)
		}
	}

	return err
}

func (t TimeUTC) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeUTCLayout)
}
