package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeEUAUnmarshal(t *testing.T) {
	assert := assert.New(t)

	data := TimeEUA{Time: time.Now()}

	st := struct {
		Data TimeEUA `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeEUALayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimeEUAMarshal(t *testing.T) {
	assert := assert.New(t)

	dataTime := time.Now()
	data := TimeEUA{Time: dataTime}

	st := struct {
		Data TimeEUA `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeEUALayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(dataTime.Format(TimeEUALayout), stActual.Data.Format(TimeEUALayout))
}

func TestTimeEUAScan(t *testing.T) {
	assert := assert.New(t)

	dateString := "2018-12-08"
	dateActual := TimeEUA{}

	err := dateActual.Scan(dateString)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dateTime, _ := time.Parse(TimeEUALayout, dateString)
	dateActual = TimeEUA{}

	err = dateActual.Scan(dateTime)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dataTimeEUA := TimeEUA{Time: dateTime}
	dateActual = TimeEUA{}

	err = dateActual.Scan(dataTimeEUA)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())
}

func TestTimeEUAValue(t *testing.T) {
	assert := assert.New(t)

	dateString := "2018-12-08"
	dateTime, _ := time.Parse(TimeEUALayout, dateString)
	dateActual := TimeEUA{Time: dateTime}

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(dateString, valueString)
}
