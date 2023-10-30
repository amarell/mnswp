package math

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	fibSequence := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	
	f := Fibonacci()

	for i := 0; i < 10; i++ {
		next := f()

		if(next != fibSequence[i]) {
			t.Fatal("Not the right sequence!")
		}
	}
}