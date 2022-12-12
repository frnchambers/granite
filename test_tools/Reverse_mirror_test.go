package test_tools

import "testing"

func Test_reverse_array(t *testing.T) {

	// given
	array := []int{1, 2, 3, 4}

	// when
	actual := Make_reverse_array(array)

	// expect
	expected := []int{4, 3, 2, 1}

	if !Are_int_slices_equal(expected, actual) {
		t.Fatalf(
			"Test_reverse_array: expected_array = %v, actual_array = %v",
			expected, actual,
		)
	}
}

func Test_mirror_even_array(t *testing.T) {

	// given
	is_even := true
	array := []int{1, 2, 3, 4}

	// when
	actual := Make_mirror_array(array, is_even)

	// expect
	expected := []int{1, 2, 3, 4, 4, 3, 2, 1}

	if !Are_int_slices_equal(expected, actual) {
		t.Fatalf(
			"Test_mirror_even_array: expected_array = %v, actual_array = %v",
			expected, actual,
		)
	}
}

func Test_mirror_odd_array(t *testing.T) {

	// given
	is_even := false
	array := []int{1, 2, 3, 4}

	// when
	actual := Make_mirror_array(array, is_even)

	// expect
	expected := []int{1, 2, 3, 4, 3, 2, 1}

	if !Are_int_slices_equal(expected, actual) {
		t.Fatalf(
			"Test_mirror_even_array: expected_array = %v, actual_array = %v",
			expected, actual,
		)
	}
}
