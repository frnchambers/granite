package integrator

import (
	"testing"

	"github.com/granite/pkg/comparison"
)

func Test_reverse_array(t *testing.T) {

	tests := []struct {
		input, expect []int
	}{
		{input: []int{1, 2, 3, 4}, expect: []int{4, 3, 2, 1}},
		{input: []int{1, 2, 5}, expect: []int{5, 2, 1}},
	}

	for i, test := range tests {

		// when
		actual := make_reverse_array(test.input)

		if !comparison.Are_int_slices_equal(actual, test.expect) {
			t.Fatalf(
				"Test, i = %v, %v, failed : expect = %v, actual = %v",
				i, test, test.expect, actual,
			)
		}
	}
}

func Test_mirror_array(t *testing.T) {

	tests := []struct {
		input, expect []int
		is_even       bool
	}{
		{input: []int{1, 2, 3, 4}, is_even: true, expect: []int{1, 2, 3, 4, 4, 3, 2, 1}},
		{input: []int{1, 2, 5}, is_even: true, expect: []int{1, 2, 5, 5, 2, 1}},
		{input: []int{3, 4, 6}, is_even: false, expect: []int{3, 4, 6, 4, 3}},
		{input: []int{1, 3, 5}, is_even: false, expect: []int{1, 3, 5, 3, 1}},
	}

	for i, test := range tests {

		// when
		actual := make_mirror_array(test.input, test.is_even)

		if !comparison.Are_int_slices_equal(actual, test.expect) {
			t.Fatalf(
				"Test, i = %v, %v, failed : expect = %v, actual = %v",
				i, test, test.expect, actual,
			)
		}
	}
}
