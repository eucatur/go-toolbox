package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const TimestampLayout = "2006-01-02 15:04:05"

type Timestamp struct{ time.Time }

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
