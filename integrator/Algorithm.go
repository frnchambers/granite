package integrator

import (
	"fmt"
	"math"

	"main.go/test_tools"
)

const (
	FORCE_GRADIENT_STEP_TOLERANCE = 1.0e-9
)

type Algorithm_t struct {
	is_velocity_version           bool
	stages, fg_steps, error_order int
	unique_velocity_coefficients  []float64
	unique_force_coefficients     []float64
	unique_fg_coefficients        []float64
}

func (algorithm *Algorithm_t) String() (description string) {
	description = fmt.Sprintf("%d-stage, O(%d) ", algorithm.stages, algorithm.error_order)
	if algorithm.is_velocity_version {
		description += "velocity"
	} else {
		description += "position"
	}
	description += fmt.Sprintf(" version algorithm, with %d force-gradient step", algorithm.fg_steps)
	return
}

func should_perform_fg_step(force_gradient_coefficient float64) bool {
	if math.Abs(force_gradient_coefficient) < FORCE_GRADIENT_STEP_TOLERANCE {
		return false
	}
	return true
}

func (algorithm *Algorithm_t) velocity_stages() []float64 {
	is_even := !(algorithm.is_velocity_version == algorithm.are_unique_lengths_equal())
	return test_tools.Make_mirror_array(algorithm.unique_velocity_coefficients, is_even)
}

func (algorithm *Algorithm_t) force_stages() []float64 {
	is_even := (algorithm.is_velocity_version == algorithm.are_unique_lengths_equal())
	return test_tools.Make_mirror_array(algorithm.unique_force_coefficients, is_even)
}

func (algorithm *Algorithm_t) force_gradient_stages() []float64 {
	is_even := (algorithm.is_velocity_version == algorithm.are_unique_lengths_equal())
	return test_tools.Make_mirror_array(algorithm.unique_fg_coefficients, is_even)
}

func (algorithm *Algorithm_t) are_unique_lengths_equal() bool {
	if len(algorithm.unique_velocity_coefficients) == len(algorithm.unique_force_coefficients) {
		return true
	}
	return false
}
