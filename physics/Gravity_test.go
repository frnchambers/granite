package physics

import (
	"testing"

	"github.com/granite/test_tools"
)

func Test_Gravitational_force_two_particles(t *testing.T) {

	// given
	m_sol, m_earth, distance := 1.0e30, 6.0e24, 1.0e11
	sun := New_particle("", m_sol, test_tools.Null_vector(), test_tools.Null_vector())
	earth := New_particle("", m_earth, test_tools.New_vector(distance, 0), test_tools.Null_vector())

	// when
	g := Gravitational_interaction_t{}
	actual_force := g.force_on_p_from_q(&earth, &sun)

	// expect
	expected_magnitude := -G * m_sol * m_earth / (distance * distance)
	expected_force := test_tools.New_vector(expected_magnitude, 0)

	if !test_tools.Vector_equality(expected_force, actual_force) {
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
		New_particle("", mass, test_tools.New_vector(0, 0), test_tools.Null_vector()),
		New_particle("", mass, test_tools.New_vector(distance, 0), test_tools.Null_vector()),
		New_particle("", mass, test_tools.New_vector(-distance, 0), test_tools.Null_vector()),
	}

	// when
	g := Gravity_interparticle_t{}
	g.Calculate_forces(particles)

	// expect
	expected_force := test_tools.Null_vector()
	actual_force := particles[0].Force

	if !test_tools.Vector_equality(expected_force, actual_force) {
		t.Fatalf(
			"Gravitaty_interparticle_force(): expect_force = %v, actual_force = %v",
			expected_force, actual_force,
		)
	}

}

func Test_Gravity_central_body_force(t *testing.T) {

	// given
	central_mass := 1.0e30
	central_body := New_particle("", central_mass, test_tools.Null_vector(), test_tools.Null_vector())

	mass, distance := 1.0e24, 1.0e11
	particles := []Particle_t{
		New_particle("", mass, test_tools.New_vector(distance, 0), test_tools.Null_vector()),
		New_particle("", mass, test_tools.New_vector(-distance, 0), test_tools.Null_vector()),
	}

	// when
	g := New_massive_body(central_body)
	g.Calculate_forces(particles)

	// expect
	expected_magnitude := G * central_mass * mass / (distance * distance)

	actual_force_0 := particles[0].Force
	expected_force_0 := test_tools.New_vector(-expected_magnitude, 0)

	actual_force_1 := particles[1].Force
	expected_force_1 := test_tools.New_vector(expected_magnitude, 0)

	if !test_tools.Vector_equality(expected_force_0, actual_force_0) ||
		!test_tools.Vector_equality(expected_force_1, actual_force_1) {
		t.Fatalf(
			"Gravitaty_interparticle_force(): (expect_force_0 = %v, actual_force_0 = %v), (expect_force_1 = %v, actual_force_1 = %v)",
			expected_force_0, actual_force_0, expected_force_1, actual_force_1,
		)
	}

}
