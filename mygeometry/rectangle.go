package mygeometry

import "math"

func RecPerimeter(sideA float64, sideB float64) float64 {
	return sideA + sideB
}

func RecArea(sideA float64, sideB float64) float64 {
	return sideA * sideB
}

func RecDiagonal(sideA float64, sideB float64) float64 {
	return math.Sqrt(sideA*sideA + sideB*sideB)
}
