package root_solver

import (
	"testing"

	"github.com/granite/comparison"
)

func Test_bisection_expect_success(t *testing.T) {

	// given
	params := linear_params{m: 1.0, c: -1.5}
	bisec := New_bisection_parameters(-10.0, 10.0)

	// when
	actual_root, err := Bisection(linear, &params, &bisec)

	// expect
	expected_root := 1.5

	if !(err == nil) {
		t.Fatalf("Test_bisection: error threw unexpectedly")
	}

	if !(comparison.Float64_equality_within_tolerance(expected_root, actual_root, bisec.Tolerance*10.0)) {
		t.Fatalf(
			"Test_bisection: expected_root = %v, actual_root = %v",
			expected_root, actual_root,
		)
	}

}

func Test_bisection_single_variable_expect_success(t *testing.T) {

	// given
	bisec := New_bisection_parameters(-10.0, 10.0)

	// when
	actual_root, err := Bisection_single_variable(linear_single_valued, &bisec)

	// expect
	expected_root := 1.5

	if !(err == nil) {
		t.Fatalf("Test_bisection: error threw unexpectedly")
	}

	if !(comparison.Float64_equality_within_tolerance(expected_root, actual_root, bisec.Tolerance*10.0)) {
		t.Fatalf(
			"Test_bisection: expected_root = %v, actual_root = %v",
			expected_root, actual_root,
		)
	}

}

func Test_bisection_not_straddling_root(t *testing.T) {

	// given
	params := linear_params{m: 1.0, c: -1.5}
	bisec := New_bisection_parameters(-20.0, -10.0)

	// when
	_, err := Bisection(linear, &params, &bisec)

	// expect

	if err == nil {
		t.Fatalf("Test_bisection_not_straddling_root: bisection did not throw")
	}
}

func Test_bisection_too_many_iterations(t *testing.T) {

	// given
	params := linear_params{m: 1.0, c: -1.5}
	bisec := New_bisection_parameters(-10.0, 10.0)
	bisec.Max_iterations = 2

	// when
	_, err := Bisection(linear, &params, &bisec)

	// expect
	if err == nil {
		t.Fatalf("Test_bisection_too_many_iteractions: bisection did not throw")
	}
}

type linear_params struct{ m, c float64 }

func linear(x float64, p *linear_params) float64 {
	return p.m*x + p.c
}

func linear_single_valued(x float64) float64 {
	return x - 1.5
}
