package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimeISO8601Layout = "2006-01-02T15:04:05Z"

type TimeISO8601 struct{ time.Time }

func (t TimeISO8601) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *TimeISO8601) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimeISO8601Layout, value)
	}
	return err
}

func (t TimeISO8601) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeISO8601) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case TimeISO8601:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimeISO8601Layout, v)
		}
	}

	return err
}

func (t TimeISO8601) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeISO8601Layout)
}
