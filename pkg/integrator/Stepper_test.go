package integrator

import "testing"

func Test_make_steps_for_velocity_version_with_unique_parameters_equal(t *testing.T) {

	// given
	algorithm := Algorithm_t{
		is_velocity_version: true,
		stages:              7, fg_steps: 0, error_order: 0,
		unique_velocity_coefficients: []float64{11, 12},
		unique_force_coefficients:    []float64{21, 22},
		unique_fg_coefficients:       []float64{31, 0},
	}

	// when
	steps := make_steps(algorithm)

	// expect
	matches := is_force_and_force_gradient_step(steps[0]) &&
		is_velocity_step(steps[1]) &&
		is_force_step(steps[2]) &&
		is_velocity_step(steps[3]) &&
		is_force_step(steps[4]) &&
		is_velocity_step(steps[5]) &&
		is_force_and_force_gradient_step(steps[6])

	if !matches {
		t.Fatalf(
			"Test_make_steps: steps = %v", steps,
		)
	}

}
