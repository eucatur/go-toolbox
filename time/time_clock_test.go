package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeClockUnmarshal(t *testing.T) {
	assert := assert.New(t)

	data := TimeClock{Time: time.Now()}

	st := struct {
		Data TimeClock `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeClockLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimeClockMarshal(t *testing.T) {
	assert := assert.New(t)

	dataTime := time.Now()
	data := TimeClock{Time: dataTime}

	st := struct {
		Data TimeClock `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeClockLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(dataTime.Format(TimeClockLayout), stActual.Data.Format(TimeClockLayout))
}

func TestTimeClockScan(t *testing.T) {
	assert := assert.New(t)

	dateString := "08:15"
	dateActual := TimeClock{}

	err := dateActual.Scan(dateString)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dateTime, _ := time.Parse(TimeClockLayout, dateString)
	dateActual = TimeClock{}

	err = dateActual.Scan(dateTime)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dataTimeClock := TimeClock{Time: dateTime}
	dateActual = TimeClock{}

	err = dateActual.Scan(dataTimeClock)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())
}

func TestTimeClockValue(t *testing.T) {
	assert := assert.New(t)

	dateString := "08:15"
	dateTime, _ := time.Parse(TimeClockLayout, dateString)
	dateActual := TimeClock{Time: dateTime}

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(dateString, valueString)
}
