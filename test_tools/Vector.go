package test_tools

import (
	"math"

	"gonum.org/v1/gonum/spatial/r2"
)

const (
	FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE = 1.0e-9
)

func New_vector(x, y float64) r2.Vec {
	return r2.Vec{X: x, Y: y}
}

func Destructure_vector(vec r2.Vec) (float64, float64) {
	return vec.X, vec.Y
}

func Null_vector() r2.Vec {
	return New_vector(0, 0)
}

func Float64_equality(a, b float64) bool {
	return Float64_equality_within_tolerance(a, b, FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE)
}

func Float64_equality_within_tolerance(a, b, tol float64) bool {
	difference := math.Abs(a - b)
	switch {
	case a == b:
		return true
	case math.Abs(b) < tol:
		return difference < tol
	case math.Abs(difference/b) > tol:
		return false
	default:
		return true
	}
}

func Vector_equality(v, w r2.Vec) bool {
	return Vector_equality_within_tolerance(v, w, FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE)
}

func Vector_equality_within_tolerance(v, w r2.Vec, tol float64) bool {
	if !Float64_equality_within_tolerance(v.X, w.X, tol) ||
		!Float64_equality_within_tolerance(v.Y, w.Y, tol) {
		return false
	}
	return true
}

func Are_float64_slices_equal(arr_1, arr_2 []float64) bool {
	if len(arr_1) != len(arr_2) {
		return false
	}
	for i := range arr_1 {
		if !Float64_equality(arr_1[i], arr_2[i]) {
			return false
		}
	}
	return true
}

func Are_int_slices_equal(arr_1, arr_2 []int) bool {
	if len(arr_1) != len(arr_2) {
		return false
	}
	for i := range arr_1 {
		if arr_1[i] != arr_2[i] {
			return false
		}
	}
	return true
}

func Make_mirror_array[T any](array []T, is_even bool) []T {
	rev := Make_reverse_array(array)
	if is_even {
		return append(array, rev...)
	} else {
		return append(array, rev[1:]...)
	}
}

func Make_reverse_array[T any](array []T) (reverse_array []T) {
	n := len(array)
	reverse_array = make([]T, n)
	for i := 0; i < n; i++ {
		reverse_array[i] = array[n-1-i]
	}
	return
}
