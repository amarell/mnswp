package util

import (
	"reflect"
	"testing"
)

func TestBogo(t *testing.T) {
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned",
		},
		//Reversed slice
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
			name:     "Reversed Unsigned",
		},
		//Sorted slice
		{
			input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Signed",
		},
		//Reversed slice
		{
			input:    []int{2, 1, -1, -2, -3},
			expected: []int{-3, -2, -1, 1, 2},
			name:     "Reversed Signed ",
		},
		//Random order with repetitions
		{
			input:    []int{-5, 7, 4, -2, 7},
			expected: []int{-5, -2, 4, 7, 7},
			name:     "Random order Signed",
		},
		//Single-entry slice
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Singleton",
		},
		// Empty slice
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice",
		},
	}

	for _, test := range sortTests {

		actual := Bogo(test.input)
		sorted := reflect.DeepEqual(actual, test.expected)
		if !sorted {
			t.Fatalf("Array wasn't sorted - actual: %v - expected: %v", actual, test.expected)
		}
	}
}
