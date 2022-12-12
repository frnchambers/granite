package integrator

import (
	"testing"

	"github.com/granite/comparison"
)

func Test_stages_for_velocity_version_with_unique_parameters_equal(t *testing.T) {

	// given
	algorithm := Algorithm_t{
		is_velocity_version: true,
		stages:              7, fg_steps: 4, error_order: 0,
		unique_velocity_coefficients: []float64{11, 12},
		unique_force_coefficients:    []float64{21, 22},
		unique_fg_coefficients:       []float64{31, 32},
	}

	// when
	v_actual := algorithm.velocity_stages()
	f_actual := algorithm.force_stages()
	fg_actual := algorithm.force_gradient_stages()

	// expect
	v_expected := []float64{11, 12, 11}
	f_expected := []float64{21, 22, 22, 21}
	fg_expected := []float64{31, 32, 32, 31}

	if !(comparison.Are_float64_slices_equal(v_expected, v_actual)) {
		t.Fatalf(
			"Test_stage_array: v_expected = %v, v_actual = %v",
			v_expected, v_actual,
		)
	}

	if !(comparison.Are_float64_slices_equal(f_expected, f_actual)) {
		t.Fatalf(
			"Test_stage_array: f_expected = %v, f_actual = %v",
			f_expected, f_actual,
		)
	}

	if !(comparison.Are_float64_slices_equal(fg_expected, fg_actual)) {
		t.Fatalf(
			"Test_stage_array: fg_expected = %v, fg_actual = %v",
			fg_expected, fg_actual,
		)
	}
}

func Test_stages_for_position_version_with_unique_parameters_unequal(t *testing.T) {

	// given
	algorithm := Algorithm_t{
		is_velocity_version: false,
		stages:              5, fg_steps: 2, error_order: 2,
		unique_velocity_coefficients: []float64{11, 12},
		unique_force_coefficients:    []float64{21},
		unique_fg_coefficients:       []float64{31},
	}

	// when
	v_actual := algorithm.velocity_stages()
	f_actual := algorithm.force_stages()
	fg_actual := algorithm.force_gradient_stages()

	// expect
	v_expected := []float64{11, 12, 11}
	f_expected := []float64{21, 21}
	fg_expected := []float64{31, 31}

	if !(comparison.Are_float64_slices_equal(v_expected, v_actual)) {
		t.Fatalf(
			"Test_stage_array: v_expected = %v, v_actual = %v",
			v_expected, v_actual,
		)
	}

	if !(comparison.Are_float64_slices_equal(f_expected, f_actual)) {
		t.Fatalf(
			"Test_stage_array: f_expected = %v, f_actual = %v",
			f_expected, f_actual,
		)
	}

	if !(comparison.Are_float64_slices_equal(fg_expected, fg_actual)) {
		t.Fatalf(
			"Test_stage_array: fg_expected = %v, fg_actual = %v",
			fg_expected, fg_actual,
		)
	}
}
