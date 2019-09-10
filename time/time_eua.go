package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimeEUALayout = "2006-01-02"

type TimeEUA struct{ time.Time }

func (t TimeEUA) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *TimeEUA) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimeEUALayout, value)
	}
	return err
}

func (t TimeEUA) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeEUA) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case TimeEUA:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimeEUALayout, v)
		}
	}

	return err
}

func (t TimeEUA) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeEUALayout)
}
