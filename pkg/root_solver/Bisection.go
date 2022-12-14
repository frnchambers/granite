package root_solver

import (
	"errors"
	"fmt"

	"github.com/granite/pkg/comparison"
)

const (
	DEFAULT_MAXIMUM_ITERATIONS = 100
	DEFAULT_TOLERANCE          = comparison.DEFAULT_FLOAT64_FRACTION_DIFFERENCE_TOLERANCE
)

func single_variable_handler(x float64, fn *func(float64) float64) float64 {
	return (*fn)(x)
}

type Bisection_parameters_t struct {
	X_0, X_1, Y_desired float64
	Tolerance           float64
	Max_iterations      int
}

func New_bisection_parameters(x_0, x_1 float64) Bisection_parameters_t {
	return Bisection_parameters_t{
		X_0:            x_0,
		X_1:            x_1,
		Y_desired:      0.0,
		Tolerance:      DEFAULT_TOLERANCE,
		Max_iterations: DEFAULT_MAXIMUM_ITERATIONS,
	}
}

func Bisection_single_variable(
	fn func(float64) float64, bisec *Bisection_parameters_t,
) (root float64, err error) {
	return Bisection(single_variable_handler, &fn, bisec)
}

func Bisection[fn_parameters_t any](
	fn func(float64, *fn_parameters_t) float64,
	fn_parameters *fn_parameters_t,
	bisec *Bisection_parameters_t,
) (root float64, err error) {

	fn_0, fn_1 := fn(bisec.X_0, fn_parameters)-bisec.Y_desired, fn(bisec.X_1, fn_parameters)-bisec.Y_desired
	if comparison.Float64_equality_within_tolerance(fn_0, 0.0, bisec.Tolerance) {
		return bisec.X_0, nil
	}
	if comparison.Float64_equality_within_tolerance(fn_1, 0.0, bisec.Tolerance) {
		return bisec.X_1, nil
	}

	if fn_0*fn_1 > 0.0 {
		message := fmt.Sprintf("Bisection: given points do not straddle zero: (x_0, fn_0) = (%.2e, %.2e), (x_1, fn_1) = (%.2e, %.2e)",
			bisec.X_0, fn_0, bisec.X_1, fn_1)
		return root, errors.New(message)
	}

	var fn_root float64
	is_converged, iter := false, 0
	for !is_converged && iter <= bisec.Max_iterations {

		root = 0.5 * (bisec.X_0 + bisec.X_1)
		fn_root = fn(root, fn_parameters) - bisec.Y_desired

		if fn_root*fn_0 > 0 {
			bisec.X_0 = root
			fn_0 = fn_root
		} else if fn_root*fn_1 > 0 {
			bisec.X_1 = root
			fn_1 = fn_root
		}

		if comparison.Float64_equality_within_tolerance(fn_root, 0.0, bisec.Tolerance) {
			is_converged = true
		}

		iter++
	}

	if !is_converged {
		message := fmt.Sprintf("Bisection: exhausted max number of iteraction (%d) zero: (x_0, fn_0) = (%.2e, %.2e), (x_1, fn_1) = (%.2e, %.2e)",
			bisec.Max_iterations, bisec.X_0, fn_0, bisec.X_1, fn_1)
		return root, errors.New(message)
	}

	return
}
