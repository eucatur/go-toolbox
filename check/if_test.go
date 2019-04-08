package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	res := If(2 > 1, true, false)
	assert.Equal(t, true, res)

	res = If(2 < 1, true, false)
	assert.Equal(t, false, res)

	trueValue := "true"
	falseValue := "false"
	result := ""

	result = If(2 > 1, trueValue, falseValue).(string)
	assert.Equal(t, trueValue, result)

	result = If(2 < 1, trueValue, falseValue).(string)
	assert.Equal(t, falseValue, result)
}

func TestIfFunc(t *testing.T) {
	trueFunc := func() interface{} { return "true" }
	falseFunc := func() interface{} { return "false" }
	result := ""

	result = IfFunc(2 > 1, trueFunc, falseFunc).(string)
	assert.Equal(t, "true", result)

	result = IfFunc(2 < 1, trueFunc, falseFunc).(string)
	assert.Equal(t, "false", result)
}
