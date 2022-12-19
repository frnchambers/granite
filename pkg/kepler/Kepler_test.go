package kepler

import (
	"math"
	"testing"

	"github.com/granite/pkg/comparison"
	"github.com/granite/pkg/vector"
)

func Test_Tau_circular(t *testing.T) {

	// given
	ecc := 0.0
	sample_phis := []float64{0.25, 0.7, 1.1}

	for i := range sample_phis {

		// given
		phi := sample_phis[i]

		// when
		actual := tau(phi, ecc)

		// expect
		expect := phi

		if !comparison.Float64_equality(expect, actual) {
			t.Fatalf(
				"Test %v, phi = %v, failed: expect = %v, actual = %v",
				i, phi, expect, actual,
			)
		}
	}
}

func Test_Tau_elliptical(t *testing.T) {

	// given
	ecc := 0.3
	sample_phis := []float64{0.0, math.Pi}
	expected_tau := []float64{0.0, math.Pi}

	for i := range sample_phis {

		// given
		phi := sample_phis[i]

		// when
		actual := tau(phi, ecc)

		// expect
		expect := expected_tau[i]

		if !comparison.Float64_equality(expect, actual) {
			t.Fatalf(
				"Test %v, phi = %v, failed: expect = %v, actual = %v",
				i, phi, expect, actual,
			)
		}
	}
}

func Test_Distance(t *testing.T) {

	// given
	orbit := New_elliptical_orbit(1.0, 0.4, 1.0)
	sample_phis := []float64{0.0, math.Pi, -math.Pi}
	expected_distances := []float64{orbit.Perihelion, orbit.Aphelion, orbit.Aphelion}

	for i := range sample_phis {

		// given
		phi := sample_phis[i]

		// when
		actual := Distance_from_centre(phi, &orbit)

		// expect
		expect := expected_distances[i]

		if !comparison.Float64_equality(expect, actual) {
			t.Fatalf(
				"Test %v, phi = %v, failed: expect = %v, actual = %v",
				i, phi, expect, actual,
			)
		}
	}
}

func Test_Speed(t *testing.T) {

	// given
	orbit := New_elliptical_orbit(1.0, 0.4, 1.0)
	sample_phis := []float64{0.0, math.Pi, -math.Pi}
	expected_speeds := []float64{orbit.V_perihelion, orbit.V_aphelion, orbit.V_aphelion}

	for i := range sample_phis {

		// given
		phi := sample_phis[i]

		// when
		actual := Speed_along_ellipse(phi, &orbit)

		// expect
		expect := expected_speeds[i]

		if !comparison.Float64_equality(expect, actual) {
			t.Fatalf(
				"Test %v, phi = %v, failed: expect = %v, actual = %v",
				i, phi, expect, actual,
			)
		}

	}

}

func Test_Time_to_perihelion(t *testing.T) {

	// given
	orbit := New_elliptical_orbit(1.0, 0.4, 1.0)
	sample_phis := []float64{math.Pi, -math.Pi}
	expected_speeds := []float64{orbit.Period / 2.0, orbit.Period / 2.0}

	for i := range sample_phis {

		// given
		phi := sample_phis[i]

		// when
		actual := math.Abs(Time_to_perihelion(phi, &orbit))

		// expect
		expect := expected_speeds[i]

		if !comparison.Float64_equality(expect, actual) {
			t.Fatalf(
				"Test %v, phi = %v, failed: expect = %v, actual = %v",
				i, phi, expect, actual,
			)
		}
	}
}

func Test_Phi_for_time_to_perihelion(t *testing.T) {

	// given
	orbit := New_elliptical_orbit(1.0, 0.4, 1.0)
	time := orbit.Period / 2.0

	// when
	actual, err := Phi_for_time_to_perihelion(time, &orbit)

	// expect
	expect := []float64{math.Pi, -math.Pi}

	if err != nil {
		t.Fatalf("Test failed: no error expected but returned: %v", err)
	}

	if !comparison.Float64_equality(expect[0], actual) || !comparison.Float64_equality(expect[0], actual) {
		t.Fatalf(
			"Test t = %v, failed: expect = +/- %v, actual = %v",
			time, expect[0], actual,
		)
	}

}

func Test_Tangent_to_ellipse(t *testing.T) {

	// given
	orbit := New_elliptical_orbit(1.0, 0.4, 1.0)
	sample_phis := []float64{
		0.0,
		math.Pi,
		math.Pi - math.Atan(orbit.Semi_minor/orbit.Linear_eccentricity),
	}

	expected_tangents := []vector.Vec{
		{X: 0.0, Y: 1.0},
		{X: 0.0, Y: -1.0},
		{X: -1.0, Y: 0.0},
	}

	for i := range sample_phis {

		// given
		phi := sample_phis[i]

		// when
		actual := Tangent_along_ellipse(phi, &orbit)

		// expect
		expect := expected_tangents[i]

		if !vector.Are_equal(expect, actual) {
			t.Fatalf(
				"Test %v, phi = %v, failed: expect = %v, actual = %v",
				i, phi, expect, actual,
			)
		}
	}
}
