package random

import (
	"testing"
)

func Test_random_sign(t *testing.T) {

	n_tests := 15

	for i := 0; i < n_tests; i++ {

		// when
		sample := sign()

		if !(sample == -1 || sample == +1) {
			t.Fatalf(
				"Test, %d, failed: expected = [-1, 1], actual = %v",
				i, sample,
			)
		}
	}
}
