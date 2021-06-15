package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceExistsString(t *testing.T) {
	values := []string{
		"foo",
		"anotherfoo",
		"anywherefoo",
	}

	assert.Equal(t, false, SliceExists(values, "everywherefoo"))

	assert.Equal(t, true, SliceExists(values, "anotherfoo"))

}

func Test_SliceExistsInteger(t *testing.T) {

	values := []int{
		1,
		2,
		3,
		4,
	}

	assert.Equal(t, false, SliceExists(values, 0))

	assert.Equal(t, true, SliceExists(values, 3))
}

func Test_SliceExistsFloat(t *testing.T) {

	values := []float64{
		1.5,
		44.7,
		55.0,
		33.7,
	}

	assert.Equal(t, false, SliceExists(values, 0.2))

	assert.Equal(t, true, SliceExists(values, 33.7))
}

func Test_SliceExistsMapStrings(t *testing.T) {

	values := make(map[string][]string)

	values["bla"] = []string{"blabla"}

	assert.Equal(t, false, SliceExists(values["bla"], "test"))

	values["foo"] = []string{"anything"}

	assert.Equal(t, true, SliceExists(values["foo"], "anything"))

}
