package util

import (
	"testing"
)

func TestArrayContainsMethod(t *testing.T) {
	sortTests := []struct {
		array    []int
		target   int
		expected bool
		name     string
	}{
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   5,
			expected: true,
			name:     "Element exists in the array",
		},
		{
			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   420,
			expected: false,
			name:     "Element does not exist in the array",
		},
	}

	for _, test := range sortTests {
		actual := ArrayContains(test.array, test.target)
		if actual != test.expected {
			t.Fatalf("Test failed. Expected %v, Actual %v --- Test name: %v", actual, test.expected, test.name)
		}
	}
}
