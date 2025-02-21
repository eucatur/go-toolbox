package gomap

import (
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected map[string]int
	}{
		{
			name: "Merge two maps with unique keys",
			input: []map[string]int{
				{"a": 1, "b": 2},
				{"c": 3, "d": 4},
			},
			expected: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name: "Merge with overlapping keys (should overwrite)",
			input: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 3, "c": 4},
			},
			expected: map[string]int{"a": 1, "b": 3, "c": 4},
		},
		{
			name: "Merge empty maps",
			input: []map[string]int{
				{},
				{},
			},
			expected: map[string]int{},
		},
		{
			name: "Merge single map",
			input: []map[string]int{
				{"x": 100},
			},
			expected: map[string]int{"x": 100},
		},
		{
			name:     "Merge with no maps",
			input:    []map[string]int{},
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeMaps(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}

}
