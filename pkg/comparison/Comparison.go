package comparison

import "math"

const (
	DEFAULT_FLOAT64_FRACTION_DIFFERENCE_TOLERANCE = 1.0e-9
)

func Float64_equality(a, b float64) bool {
	return Float64_equality_within_tolerance(a, b, DEFAULT_FLOAT64_FRACTION_DIFFERENCE_TOLERANCE)
}

func Float64_equality_within_tolerance(a, b, tol float64) bool {
	difference := math.Abs(a - b)
	switch {
	case a == b:
		return true
	case math.Abs(b) < tol:
		return difference < tol
	default:
		return math.Abs(difference/b) < tol
	}
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
