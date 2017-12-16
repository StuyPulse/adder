package adder

func Add(n1, n2 int) int {
	return n1 + n2
}

func Multiply(n1, n2 int) int {
	return n1 * n2
}

func Divide(dividend, divisor int) int {
	if divisor == 0 {
		return 0
	}

	return dividend / divisor
}
