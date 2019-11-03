package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHumanizedDuration(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	assert := assert.New(t)

	t1, _ := time.Parse(layout, "2019-11-01 01:02:00")
	t2, _ := time.Parse(layout, "2019-11-02 16:06:00")

	expected := "1 dia 15 horas 4 minutos"

	actual := HumanizedDuration(t2.Sub(t1))
	assert.Equal(expected, actual)

	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-03 02:03:00")
	expected = "2 dias 1 hora 1 minuto"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-01 03:04:00")
	expected = "2 horas 2 minutos"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-01 01:05:00")
	expected = "3 minutos"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)

	t2, _ = time.Parse(layout, "2019-11-01 04:02:00")
	expected = "3 horas"
	actual = HumanizedDuration(t1.Sub(t2))
	assert.Equal(expected, actual)
}
