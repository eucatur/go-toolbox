package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimeClockLayout = "15:04"

type TimeClock struct{ time.Time }

func (t TimeClock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *TimeClock) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimeClockLayout, value)
	}
	return err
}

func (t TimeClock) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeClock) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case TimeClock:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimeClockLayout, v)
		}
	case []uint8:
		if len(v) > 0 {
			t.Time, err = time.Parse(TimeClockLayout, string(v))
		}
	}

	return err
}

func (t TimeClock) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeClockLayout)
}

// ParseTimeClock ...
func ParseTimeClock(value string) (timeClock TimeClock, err error) {
	timeClock.Time, err = time.Parse(TimeClockLayout, value)
	return
}
