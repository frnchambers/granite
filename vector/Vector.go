package vector

import (
	"github.com/granite/comparison"
	"gonum.org/v1/gonum/spatial/r2"
)

type Vec = r2.Vec

func New(x, y float64) Vec {
	return Vec{X: x, Y: y}
}

func Destructure(vec Vec) (float64, float64) {
	return vec.X, vec.Y
}

func Null() Vec {
	return New(0, 0)
}

func Are_equal(v, w Vec) bool {
	return Are_equal_within_tolerance(v, w, comparison.DEFAULT_FLOAT64_FRACTION_DIFFERENCE_TOLERANCE)
}

func Are_equal_within_tolerance(v, w Vec, tol float64) bool {
	if !comparison.Float64_equality_within_tolerance(v.X, w.X, tol) ||
		!comparison.Float64_equality_within_tolerance(v.Y, w.Y, tol) {
		return false
	}
	return true
}
