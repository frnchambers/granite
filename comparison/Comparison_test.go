package comparison

import "testing"

func Test_compare_float64(t *testing.T) {

	tests := []struct {
		a, b, tol float64
		expect    bool
	}{
		{a: 1.0, b: 1.0, tol: DEFAULT_FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE, expect: true},
		{a: 1.0, b: 2.0, tol: DEFAULT_FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE, expect: false},
		{a: 0.0, b: 1.0, tol: DEFAULT_FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE, expect: false},
		{a: 1.0, b: 0.0, tol: DEFAULT_FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE, expect: false},
		{a: 1.0, b: 1.0 + 1.0e-9, tol: 1.0e-3, expect: true},
		// {a: 1.0, b: DEFAULT_FLOAT64_PERCENTAGE_DIFFERENCE_TOLERANCE, expect: false},
	}

	for i, test := range tests {
		// when
		actual := Float64_equality_within_tolerance(
			test.a, test.b, test.tol)

		if actual != test.expect {
			t.Fatalf(
				"Test, i = %v, %v, failed, : expect = %v, actual = %v",
				i, test, test.expect, actual,
			)
		}
	}
}
