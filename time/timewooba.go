package time

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var timewoobaZoneToIANA = map[string]string{
	"-0300": "America/Sao_Paulo",
	"-0400": "America/Manaus",
}

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
	if size := len(value); size < 26 || size > 29 {
		err = errors.New("Invalid date")
		return
	}

	value = strings.ReplaceAll(value, `\`, "")
	nsecStr := value[6:19] + strings.Repeat("0", 6)

	nsec, err := strconv.ParseInt(nsecStr, 10, 64)
	if err != nil {
		return
	}

	timewooba.Time = time.Unix(0, nsec)

	iana, ok := timewoobaZoneToIANA[value[19:24]]
	if !ok {
		err = errors.New("IANA not mapped")
		return
	}

	loc, err := time.LoadLocation(iana)
	if err != nil {
		return
	}

	timewooba.Time = timewooba.Time.In(loc)
	return
}
