package mymath

import "math"

func Root(f float64, r float64) float64 {
	return math.Pow(f, 1/r)
}
