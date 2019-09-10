package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeISO8601Unmarshal(t *testing.T) {
	assert := assert.New(t)

	data := TimeISO8601{Time: time.Now()}

	st := struct {
		Data TimeISO8601 `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format("2006-01-02T15:04:05Z") + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimeISO8601Marshal(t *testing.T) {
	assert := assert.New(t)

	dataTime := time.Now()
	data := TimeISO8601{Time: dataTime}

	st := struct {
		Data TimeISO8601 `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format("2006-01-02T15:04:05Z") + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(dataTime.Format("2006-01-02T15:04:05Z"), stActual.Data.Format("2006-01-02T15:04:05Z"))
}

func TestTimeISO8601Scan(t *testing.T) {
	assert := assert.New(t)

	dateString := "2018-12-08T12:50:00Z"
	dateActual := TimeISO8601{}

	err := dateActual.Scan(dateString)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dateTime, _ := time.Parse(TimeISO8601Layout, dateString)
	dateActual = TimeISO8601{}

	err = dateActual.Scan(dateTime)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dataTimeISO8601 := TimeISO8601{Time: dateTime}
	dateActual = TimeISO8601{}

	err = dateActual.Scan(dataTimeISO8601)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())
}

func TestTimeISO8601Value(t *testing.T) {
	assert := assert.New(t)

	dateString := "2018-12-08T12:50:00Z"
	dateTime, _ := time.Parse(TimeISO8601Layout, dateString)
	dateActual := TimeISO8601{Time: dateTime}

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(dateString, valueString)
}
