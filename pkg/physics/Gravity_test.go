package physics

import (
	"testing"

	"github.com/granite/pkg/vector"
)

func Test_Gravitational_force_two_particles(t *testing.T) {

	// given
	m_sol, m_earth, distance := 1.0e30, 6.0e24, 1.0e11
	sun := New_particle("", m_sol, vector.Null(), vector.Null())
	earth := New_particle("", m_earth, vector.New(distance, 0), vector.Null())

	// when
	g := Gravitational_interaction_t{}
	actual_force := g.force_on_p_from_q(&earth, &sun)

	// expect
	expected_magnitude := -G * m_sol * m_earth / (distance * distance)
	expected_force := vector.New(expected_magnitude, 0)

	if !vector.Are_equal(expected_force, actual_force) {
		t.Fatalf(
			"Gravitational_interparticle_force(): expect_force = %v, actual_force = %v",
			expected_force, actual_force,
		)
	}

}

func Test_Gravity_interparticle_force(t *testing.T) {

	// given
	mass, distance := 1.0e15, 5.0e10
	particles := []Particle_t{
		New_particle("", mass, vector.New(0, 0), vector.Null()),
		New_particle("", mass, vector.New(distance, 0), vector.Null()),
		New_particle("", mass, vector.New(-distance, 0), vector.Null()),
	}

	// when
	g := Gravity_interparticle_t{}
	g.Calculate_forces(particles)

	// expect
	expected_force := vector.Null()
	actual_force := particles[0].Force

	if !vector.Are_equal(expected_force, actual_force) {
		t.Fatalf(
			"Gravitaty_interparticle_force(): expect_force = %v, actual_force = %v",
			expected_force, actual_force,
		)
	}

}

func Test_Gravity_central_body_force(t *testing.T) {

	// given
	central_mass := 1.0e30
	central_body := New_particle("", central_mass, vector.Null(), vector.Null())

	mass, distance := 1.0e24, 1.0e11
	particles := []Particle_t{
		New_particle("", mass, vector.New(distance, 0), vector.Null()),
		New_particle("", mass, vector.New(-distance, 0), vector.Null()),
	}

	// when
	g := New_massive_body(central_body)
	g.Calculate_forces(particles)

	// expect
	expected_magnitude := G * central_mass * mass / (distance * distance)

	actual_force_0 := particles[0].Force
	expected_force_0 := vector.New(-expected_magnitude, 0)

	actual_force_1 := particles[1].Force
	expected_force_1 := vector.New(expected_magnitude, 0)

	if !vector.Are_equal(expected_force_0, actual_force_0) ||
		!vector.Are_equal(expected_force_1, actual_force_1) {
		t.Fatalf(
			"Gravitaty_interparticle_force(): (expect_force_0 = %v, actual_force_0 = %v), (expect_force_1 = %v, actual_force_1 = %v)",
			expected_force_0, actual_force_0, expected_force_1, actual_force_1,
		)
	}

}
