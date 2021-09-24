package somefunc

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	a := 0
	b := 0
	return func() int {
		sum := a + b
		if sum == 0 {
			sum = 1
		}
		a = b
		b = sum
		return sum
	}
}
