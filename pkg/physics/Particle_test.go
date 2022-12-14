package physics

import (
	"testing"

	"github.com/granite/pkg/vector"
)

func Test_Increment_position_using_velocity(t *testing.T) {

	// given
	timestep, velocity, p_0 := 3.0, 1.0, 5.0
	particle := New_blank_particle()
	particle.Position = vector.New(p_0, 0)
	particle.Velocity = vector.New(velocity, 0)

	// when
	particle.Increment_position_using_velocity(timestep)

	// expect
	expected_position := vector.New(p_0+timestep*velocity, 0)
	actual_position := particle.Position

	if !vector.Are_equal(expected_position, actual_position) {
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
	particle.Velocity = vector.New(v_0, 0)
	particle.Force = vector.New(force, 0)

	// when
	particle.Increment_velocity_using_force(timestep)

	// expect
	expected_velocity := vector.New(v_0+timestep/mass*force, 0)
	actual_velocity := particle.Velocity

	if !vector.Are_equal(expected_velocity, actual_velocity) {
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
	particle.Velocity = vector.New(v_0, 0)
	particle.Force_gradient = vector.New(force_gradient, 0)

	// when
	particle.Increment_velocity_using_force_gradient(timestep)

	// expect
	expected_velocity := vector.New(v_0+timestep/mass*force_gradient, 0)
	actual_velocity := particle.Velocity

	if !vector.Are_equal(expected_velocity, actual_velocity) {
		t.Fatalf(
			"Increment_velocity_using_force_gradient(dt = %v): expect_velocity = %v, actual_velocity = %v",
			timestep, expected_velocity, actual_velocity,
		)

	}
}
