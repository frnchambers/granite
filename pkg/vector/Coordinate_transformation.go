package vector

import "math"

func Cartesian_position_from_polar(r, phi float64) Vec {
	return Vec{X: X_from_polar(r, phi), Y: Y_from_polar(r, phi)}
}

func X_from_polar(r, phi float64) float64 {
	return r * math.Cos(phi)
}

func Y_from_polar(r, phi float64) float64 {
	return r * math.Sin(phi)
}
