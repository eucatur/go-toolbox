package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	rfc3339Date = "2020-07-20T06:05:00-03:00"
	woobaDate   = "/Date(1595235900000-0300)/"
)

func TestTimewoobaUnmarshal(t *testing.T) {
	assert := assert.New(t)

	date := Timewooba{}
	date.Time, _ = time.Parse(time.RFC3339, rfc3339Date)

	st := struct {
		Date Timewooba `json:"date" db:"data"`
	}{
		Date: date,
	}

	expectedJSON := `{"date":"` + woobaDate + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimewoobaMarshal(t *testing.T) {
	assert := assert.New(t)

	date := Timewooba{}
	date.Time, _ = time.Parse(time.RFC3339, rfc3339Date)

	st := struct {
		Date Timewooba `json:"date" db:"date"`
	}{
		Date: date,
	}

	expectedJSON := `{"date":"` + woobaDate + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(woobaDate, stActual.Date.String())
}

func TestTimewoobaScan(t *testing.T) {
	assert := assert.New(t)

	dateActual := Timewooba{}

	err := dateActual.Scan(woobaDate)
	assert.Equal(nil, err)
	assert.Equal(woobaDate, dateActual.String())

	timewooba, err := ParseTimewooba(woobaDate)
	assert.Equal(nil, err)

	dateActual = Timewooba{}

	err = dateActual.Scan(timewooba.Time)
	assert.Equal(nil, err)
	assert.Equal(woobaDate, dateActual.String())

	dataTimewooba := Timewooba{Time: timewooba.Time}
	dateActual = Timewooba{}

	err = dateActual.Scan(dataTimewooba)
	assert.Equal(nil, err)
	assert.Equal(woobaDate, dateActual.String())
}

func TestTimewoobaValue(t *testing.T) {
	assert := assert.New(t)

	dateActual, err := ParseTimewooba(woobaDate)
	assert.Equal(nil, err)

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(woobaDate, valueString)
}

func TestTimewoobaIsZero(t *testing.T) {
	assert := assert.New(t)

	woobaDate := "/Date(-62135589600000-0300)/"

	dateActual, err := ParseTimewooba(woobaDate)
	assert.Equal(nil, err)

	assert.Equal(true, dateActual.IsZero())
}
