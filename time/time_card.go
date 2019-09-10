package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimeCardLayout = "2006-01"

type TimeCard struct{ time.Time }

func (t TimeCard) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *TimeCard) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimeCardLayout, value)
	}
	return err
}

func (t TimeCard) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeCard) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case TimeCard:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimeCardLayout, v)
		}
	case []uint8:
		t.Time, err = time.Parse(TimeCardLayout, string([]byte(v[:])))
	}

	return err
}

func (t TimeCard) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeCardLayout)
}
