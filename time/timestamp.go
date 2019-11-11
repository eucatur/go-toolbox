package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimestampLayout = "2006-01-02 15:04:05"

type Timestamp struct{ time.Time }

// NowTimestamp returns the current local timestamp.
func NowTimestamp() Timestamp {
	return Timestamp{Time: time.Now()}
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *Timestamp) UnmarshalJSON(bytes []byte) error {
	value, err := strconv.Unquote(string(bytes))
	if err == nil && value != "" {
		t.Time, err = time.Parse(TimestampLayout, value)
	}
	return err
}

func (t Timestamp) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *Timestamp) Scan(value interface{}) error {
	var err error

	switch v := value.(type) {
	case Timestamp:
		t.Time = v.Time
	case time.Time:
		t.Time = v
	case string:
		if v != "" {
			t.Time, err = time.Parse(TimestampLayout, v)
		}
	}

	return err
}

func (t Timestamp) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimestampLayout)
}

// AddDate returns the time corresponding to adding the given number of years, months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns March 4, 2010.
// AddDate normalizes its result in the same way that Date does, so, for example, adding one month to October 31 yields December 1, the normalized form for November 31.
func (t Timestamp) AddDate(years int, months int, days int) Timestamp {
	t.Time = t.Time.AddDate(years, months, days)
	return t
}

// ParseTimestamp ...
func ParseTimestamp(value string) (timestamp Timestamp, err error) {
	timestamp.Time, err = time.Parse(TimestampLayout, value)
	return
}
