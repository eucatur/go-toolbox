package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestampUnmarshal(t *testing.T) {
	assert := assert.New(t)

	data := Timestamp{Time: time.Now()}

	st := struct {
		Data Timestamp `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimestampLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimestampMarshal(t *testing.T) {
	assert := assert.New(t)

	dataTime := time.Now()
	data := Timestamp{Time: dataTime}

	st := struct {
		Data Timestamp `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimestampLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(dataTime.Format(TimestampLayout), stActual.Data.Format(TimestampLayout))
}

func TestTimestampScan(t *testing.T) {
	assert := assert.New(t)

	dateString := "2019-09-06 08:42:38"
	dateActual := Timestamp{}

	err := dateActual.Scan(dateString)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dateTime, _ := time.Parse(TimestampLayout, dateString)
	dateActual = Timestamp{}

	err = dateActual.Scan(dateTime)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dataTimestamp := Timestamp{Time: dateTime}
	dateActual = Timestamp{}

	err = dateActual.Scan(dataTimestamp)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())
}

func TestTimestampValue(t *testing.T) {
	assert := assert.New(t)

	dateString := "2019-09-06 08:42:38"
	dateTime, _ := time.Parse(TimestampLayout, dateString)
	dateActual := Timestamp{Time: dateTime}

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(dateString, valueString)
}
