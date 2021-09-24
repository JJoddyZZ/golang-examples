package somemath

import (
	"math"
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
