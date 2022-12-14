package comparison

import "testing"

func Test_compare_float64(t *testing.T) {

	default_tol := DEFAULT_FLOAT64_FRACTION_DIFFERENCE_TOLERANCE

	tests := []struct {
		a, b, tol float64
		expect    bool
	}{
		{a: 1.0, b: 1.0, tol: default_tol, expect: true},
		{a: 1.0, b: 2.0, tol: default_tol, expect: false},
		{a: 1.0, b: 1.0 * (1.0 + default_tol), tol: default_tol, expect: false},
		{a: 1.0, b: 0.0, tol: default_tol, expect: false},
		{a: 0.0, b: 0.0, tol: default_tol, expect: true},
		{a: 1.0, b: default_tol, tol: default_tol, expect: false},
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
