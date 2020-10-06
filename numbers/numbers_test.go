package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimals(t *testing.T) {
	assert.Equal(t, 1., Decimals(1.123, 0))
	assert.Equal(t, 1.1, Decimals(1.123, 1))
	assert.Equal(t, 1.12, Decimals(1.123, 2))
	assert.Equal(t, 1.123, Decimals(1.123, 3))
	assert.Equal(t, 1.123, Decimals(1.123, 4))
}

func TestIntExistsInSlice(t *testing.T) {
	assert.Equal(t, true, IntExistsInSlice(1, []int{1, 2, 3}))
	assert.Equal(t, true, IntExistsInSlice(2, []int{1, 2, 3}))
	assert.Equal(t, true, IntExistsInSlice(3, []int{1, 2, 3}))

	assert.Equal(t, false, IntExistsInSlice(0, []int{1, 2, 3}))
}

func TestUniqueInts(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, UniqueInts([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, UniqueInts([]int{1, 1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, UniqueInts([]int{1, 1, 2, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, UniqueInts([]int{1, 1, 2, 2, 3, 3}))
	assert.Equal(t, []int{1, 2, 3}, UniqueInts([]int{1, 2, 3, 1, 2, 3}))

	assert.Equal(t, []int{1}, UniqueInts([]int{1, 1, 1, 1}))
	assert.Equal(t, []int{0, 1}, UniqueInts([]int{0, 1, 1, 1}))
}
