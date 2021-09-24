package somemath

import (
	"math"

	"github.com/JJoddyZZ/golang-examples/someerrs"
)

// Sqrt calculates manually sqrt using Newton's method
func Sqrt(x float64) (float64, int) {
	errLimit, itersLimit := 0.0000000001, 10
	z := x / 2
	prevZ := z
	for i := 0; i < itersLimit; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prevZ) <= errLimit {
			return z, i
		}
		prevZ = z
	}
	return z, itersLimit
}

// NoNegativeSqrt calculates manually sqrt using Newton's method
// returns error if given argument is equal or lower than 0
func NoNegativeSqrt(x someerrs.ErrNegativeSqrt) (float64, int, error) {
	if x <= 0 {
		return 0.0, 0, x
	}
	errLimit, itersLimit := 0.0000000001, 10
	z := float64(x) / 2
	prevZ := z
	for i := 0; i < itersLimit; i++ {
		z -= (z*z - float64(x)) / (2 * z)
		if math.Abs(z-prevZ) <= errLimit {
			return z, i, nil
		}
		prevZ = z
	}
	return z, itersLimit, nil
}
