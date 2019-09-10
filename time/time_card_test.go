package time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeCardUnmarshal(t *testing.T) {
	assert := assert.New(t)

	data := TimeCard{Time: time.Now()}

	st := struct {
		Data TimeCard `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeCardLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))
}

func TestTimeCardMarshal(t *testing.T) {
	assert := assert.New(t)

	dataTime := time.Now()
	data := TimeCard{Time: dataTime}

	st := struct {
		Data TimeCard `json:"data" db:"data"`
	}{
		Data: data,
	}

	expectedJSON := `{"data":"` + data.Format(TimeCardLayout) + `"}`

	actualJSON, err := json.Marshal(st)
	assert.Equal(nil, err)
	assert.Equal(expectedJSON, string(actualJSON))

	stActual := st

	err = json.Unmarshal(actualJSON, &stActual)
	assert.Equal(nil, err)
	assert.Equal(dataTime.Format(TimeCardLayout), stActual.Data.Format(TimeCardLayout))
}

func TestTimeCardScan(t *testing.T) {
	assert := assert.New(t)

	dateString := "2019-09"
	dateActual := TimeCard{}

	err := dateActual.Scan(dateString)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dateTime, _ := time.Parse(TimeCardLayout, dateString)
	dateActual = TimeCard{}

	err = dateActual.Scan(dateTime)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())

	dataTimeCard := TimeCard{Time: dateTime}
	dateActual = TimeCard{}

	err = dateActual.Scan(dataTimeCard)
	assert.Equal(nil, err)
	assert.Equal(dateString, dateActual.String())
}

func TestTimeCardValue(t *testing.T) {
	assert := assert.New(t)

	dateString := "2019-09"
	dateTime, _ := time.Parse(TimeCardLayout, dateString)
	dateActual := TimeCard{Time: dateTime}

	value, err := dateActual.Value()
	assert.Equal(nil, err)

	valueString, ok := value.(string)
	assert.Equal(true, ok)
	assert.Equal(dateString, valueString)
}
