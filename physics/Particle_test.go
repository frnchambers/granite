package physics

import (
	"testing"

	"github.com/granite/test_tools"
)

func Test_Increment_position_using_velocity(t *testing.T) {

	// given
	timestep, velocity, p_0 := 3.0, 1.0, 5.0
	particle := New_blank_particle()
	particle.Position = test_tools.New_vector(p_0, 0)
	particle.Velocity = test_tools.New_vector(velocity, 0)

	// when
	particle.Increment_position_using_velocity(timestep)

	// expect
	expected_position := test_tools.New_vector(p_0+timestep*velocity, 0)
	actual_position := particle.Position

	if !test_tools.Vector_equality(expected_position, actual_position) {
		t.Fatalf(
			"Increment_position_using_velocity(dt = %v): expect_position = %v, actual_position = %v",
			timestep, expected_position, actual_position,
		)
	}
}

func Test_Increment_velocity_using_force(t *testing.T) {

	// given
	mass, timestep, force, v_0 := 2.0, 3.0, 6.0, 3.0
	particle := New_blank_particle()
	particle.Mass = mass
	particle.Velocity = test_tools.New_vector(v_0, 0)
	particle.Force = test_tools.New_vector(force, 0)

	// when
	particle.Increment_velocity_using_force(timestep)

	// expect
	expected_velocity := test_tools.New_vector(v_0+timestep/mass*force, 0)
	actual_velocity := particle.Velocity

	if !test_tools.Vector_equality(expected_velocity, actual_velocity) {
		t.Fatalf(
			"Increment_velocity_using_force(dt = %v): expect_velocity = %v, actual_velocity = %v",
			timestep, expected_velocity, actual_velocity,
		)

	}
}

func Test_Increment_velocity_using_force_gradient(t *testing.T) {

	// given
	mass, timestep, force_gradient, v_0 := 2.0, 3.0, 1.0, 4.0
	particle := New_blank_particle()
	particle.Mass = mass
	particle.Velocity = test_tools.New_vector(v_0, 0)
	particle.Force_gradient = test_tools.New_vector(force_gradient, 0)

	// when
	particle.Increment_velocity_using_force_gradient(timestep)

	// expect
	expected_velocity := test_tools.New_vector(v_0+timestep/mass*force_gradient, 0)
	actual_velocity := particle.Velocity

	if !test_tools.Vector_equality(expected_velocity, actual_velocity) {
		t.Fatalf(
			"Increment_velocity_using_force_gradient(dt = %v): expect_velocity = %v, actual_velocity = %v",
			timestep, expected_velocity, actual_velocity,
		)

	}
}
