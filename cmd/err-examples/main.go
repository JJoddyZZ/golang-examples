package main

import (
	"fmt"
	"math"

	"github.com/JJoddyZZ/golang-examples/somemath"
)

func main() {
	z, i, _ := somemath.NoNegativeSqrt(-2)
	fmt.Printf("Own sqrt result: %g in %d iterations, math.Sqrt result: %g\n", z, i, math.Sqrt(2))

	a, b, err := somemath.NoNegativeSqrt(-2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Own sqrt result: %g in %d iterations, math.Sqrt result: %g", a, b, math.Sqrt(-2))
}
