package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimeFullClockLayout = "15:04:05"

type TimeFullClock struct{ time.Time }

func (t TimeFullClock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *TimeFullClock) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimeFullClockLayout, value)
	}
	return err
}

func (t TimeFullClock) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeFullClock) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case TimeFullClock:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimeFullClockLayout, v)
		}
	case []uint8:
		if len(v) > 0 {
			t.Time, err = time.Parse(TimeFullClockLayout, string(v))
		}
	}

	return err
}

func (t TimeFullClock) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeFullClockLayout)
}

// ParseTimeFullClock ...
func ParseTimeFullClock(value string) (timeFullClock TimeFullClock, err error) {
	timeFullClock.Time, err = time.Parse(TimeFullClockLayout, value)
	return
}
