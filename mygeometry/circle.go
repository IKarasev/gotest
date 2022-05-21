package mygeometry

import "math"

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

func CirclePerimeter(r float64) float64 {
	return 2 * math.Pi * r
}

func SphereVolume(r float64) float64 {
	return 4 / 3 * math.Pi * r * r * r
}
