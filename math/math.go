package math

func Fibonacci() func() int {
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
