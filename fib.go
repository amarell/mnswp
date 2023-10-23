package main

import "fmt"

func fibonacci() func() int {
	sum := 0
	a := -1
	b := 1

	return func() int {
		sum = 0
		sum += a + b
		a = b
		b = sum
		return sum
	}
}

func testFib() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
