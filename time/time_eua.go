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

// AddDate returns the time corresponding to adding the given number of years, months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns March 4, 2010.
// AddDate normalizes its result in the same way that Date does, so, for example, adding one month to October 31 yields December 1, the normalized form for November 31.
func (t TimeEUA) AddDate(years int, months int, days int) TimeEUA {
	t.Time = t.Time.AddDate(years, months, days)
	return t
}

// ParseTimeEUA parses a formatted string and returns the time value it represents.
func ParseTimeEUA(value string) (timeEUA TimeEUA, err error) {
	timeEUA.Time, err = time.Parse(TimeEUALayout, value)
	return
}
