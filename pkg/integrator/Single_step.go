package integrator

import (
	"fmt"

	"github.com/granite/pkg/physics"
)

type single_step_t interface {
	run(system *physics.System_t, timestep float64)
	String() string
}

type velocity_step_t struct{ coefficient float64 }
type force_step_t struct{ coefficient float64 }
type force_and_force_gradient_step_t struct{ f_coefficient, fg_coefficient float64 }

func (s *velocity_step_t) run(system *physics.System_t, timestep float64) {
	system.Increment_positions_using_velocities(s.coefficient * timestep)
}

func (s *force_step_t) run(system *physics.System_t, timestep float64) {
	system.Calculate_forces()
	system.Increment_velocities_using_forces(s.coefficient * timestep)
}

func (s *force_and_force_gradient_step_t) run(system *physics.System_t, timestep float64) {
	system.Calculate_forces()
	system.Increment_velocities_using_forces(s.f_coefficient * timestep)
	system.Calculate_force_gradients()
	system.Increment_velocities_using_force_gradients(s.fg_coefficient * timestep * timestep * timestep)
}

func (s *velocity_step_t) String() string {
	return fmt.Sprintf("exp(%.2e dt A)", s.coefficient)
}

func (s *force_step_t) String() string {
	return fmt.Sprintf("exp(%.2e dt B)", s.coefficient)
}

func (s *force_and_force_gradient_step_t) String() string {
	return fmt.Sprintf("exp(%.2e dt B + %.2e dt^3 C)", s.f_coefficient, s.fg_coefficient)
}

func is_force_and_force_gradient_step(step single_step_t) bool {
	_, ok := step.(*force_and_force_gradient_step_t)
	return ok
}

func is_force_step(step single_step_t) bool {
	_, ok := step.(*force_step_t)
	return ok
}

func is_velocity_step(step single_step_t) bool {
	_, ok := step.(*velocity_step_t)
	return ok
}
