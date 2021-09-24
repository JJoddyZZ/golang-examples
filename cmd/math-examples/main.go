package main

import (
	"fmt"
	"math"

	"github.com/JJoddyZZ/golang-examples/somemath"
)

func main() {
	z, i := somemath.Sqrt(2)
	fmt.Printf("Own sqrt result: %g in %d iterations, math.Sqrt result: %g", z, i, math.Sqrt(2))
}
