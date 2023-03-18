package physics

import (
	"testing"

	"github.com/granite/pkg/vector"
)

func Test_Centre_of_mass(t *testing.T) {

	// given
	particles := []Particle_t{
		New_particle("", 2.0, vector.New(1.0, 0.0), vector.Null()),
		New_particle("", 2.0, vector.New(-1.0, 0.0), vector.Null()),
		New_particle("", 1.0, vector.New(0.0, 5.0), vector.Null()),
	}
	// n_particles := len(particles)
	system := System_t{Force: Gravity_interparticle_t{}, Particles: particles, Time: 0.0}

	// when
	centre_of_mass := system.Centre_of_mass()

	// expect
	expected_position := vector.New(0.0, 1.0)

	if !vector.Are_equal(expected_position, centre_of_mass) {
		t.Fatalf(
			"Centre_of_mass(): expect_position = %v, centre_of_mass = %v",
			expected_position, centre_of_mass,
		)
	}
}

func Test_Average_velocity(t *testing.T) {

	// given
	particles := []Particle_t{
		New_particle("", 1.0, vector.Null(), vector.New(3.0, -1.0)),
		New_particle("", 3.0, vector.Null(), vector.New(3.0, 2.5)),
		New_particle("", 4.0, vector.Null(), vector.New(0.0, -1.5)),
	}
	// n_particles := len(particles)
	system := System_t{Force: Gravity_interparticle_t{}, Particles: particles, Time: 0.0}

	// when
	average_velocity := system.Average_velocity()

	// expect
	expected_velocity := vector.New(2.0, 0.0)

	if !vector.Are_equal(expected_velocity, average_velocity) {
		t.Fatalf(
			"Centre_of_mass(): expect_position = %v, centre_of_mass = %v",
			expected_velocity, average_velocity,
		)
	}
}
