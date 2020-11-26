package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFullClockUnmarshal(t *testing.T) {
	assert := assert.New(t)

	data := TimeFullClock{Time: time.Now()}

	st := struct {
		Data TimeFullClock `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeClockLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimeFullClockMarshal(t *testing.T) {
	assert := assert.New(t)

	dataTime := time.Now()
	data := TimeFullClock{Time: dataTime}

	st := struct {
		Data TimeFullClock `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeFullClockLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(dataTime.Format(TimeFullClockLayout), stActual.Data.Format(TimeFullClockLayout))
}

func TestTimeFullClockScan(t *testing.T) {
	assert := assert.New(t)

	dateString := "08:15"
	dateActual := TimeFullClock{}

	err := dateActual.Scan(dateString)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dateTime, _ := time.Parse(TimeFullClockLayout, dateString)
	dateActual = TimeFullClock{}

	err = dateActual.Scan(dateTime)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dataTimeClock := TimeFullClock{Time: dateTime}
	dateActual = TimeFullClock{}

	err = dateActual.Scan(dataTimeClock)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())
}

func TestTimeFullClockValue(t *testing.T) {
	assert := assert.New(t)

	dateString := "08:15"
	dateTime, _ := time.Parse(TimeFullClockLayout, dateString)
	dateActual := TimeFullClock{Time: dateTime}

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(dateString, valueString)
}
