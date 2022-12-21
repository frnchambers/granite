package physics

import (
	"fmt"
)

type System_t struct {
	Time      float64
	Force     Force_t
	Particles []Particle_t
}

func (system *System_t) N_particles() int {
	return len(system.Particles)
}

func (system *System_t) Increment_time(dt float64) {
	system.Time += dt
}

func (system *System_t) Energy() float64 {
	return system.Kinetic_energy() + system.Potential_energy()
}

func (system *System_t) Kinetic_energy() (K float64) {
	K = 0
	for i := range system.Particles {
		K += system.Particles[i].Kinetic_energy()
	}
	return
}

func (system *System_t) Potential_energy() float64 {
	return system.Force.Calculate_potentials(system.Particles)
}

func (system *System_t) Calculate_forces() {
	system.Force.Calculate_forces(system.Particles)
}

func (system *System_t) Calculate_force_gradients() {
	system.Force.Calculate_force_gradients(system.Particles)
}

func (system *System_t) Increment_positions_using_velocities(timestep float64) {
	for i := range system.Particles {
		system.Particles[i].Increment_position_using_velocity(timestep)
	}
}

func (system *System_t) Increment_velocities_using_forces(timestep float64) {
	for i := range system.Particles {
		system.Particles[i].Increment_velocity_using_force(timestep)
	}
}

func (system *System_t) Increment_velocities_using_force_gradients(timestep float64) {
	for i := range system.Particles {
		system.Particles[i].Increment_velocity_using_force_gradient(timestep)
	}
}

func (system *System_t) Increment_velocities_using_forces_and_force_gradients(timestep_force, timestep_fg float64) {
	for i := range system.Particles {
		system.Particles[i].Increment_velocity_using_force_and_force_gradient(timestep_force, timestep_fg)
	}
}

func (system System_t) String() (output string) {
	output = "System_t:\n"
	output += fmt.Sprintf("Time: %.2e\nEnergy: %.2e\nForce: %v\nParticles: [\n", system.Time, system.Energy(), system.Force)
	for _, p := range system.Particles {
		output += p.String() + ",\n"
	}
	output += "]"
	return
}

func (system System_t) Table_header() (output string) {
	output = "# System_t:\n"
	output += fmt.Sprintf("# Initial time:   %.6e\n", system.Time)
	output += fmt.Sprintf("# Initial energy: %.6e\n", system.Energy())
	// output += fmt.Sprintf("# Force: %v", system.Force)
	output += fmt.Sprintf("# Time, Energy")
	for i := range system.Particles {
		output += fmt.Sprintf("r_%d_x, r_%d_y, v_%d_x, v_%d_y, f_%d_x, f_%d_y, fg_%d_x, fg_%d_y, K_%d\n", i, i, i, i, i, i, i, i, i)
	}
	return
}

func (system System_t) As_row() (output string) {
	output += fmt.Sprintf("%.6e, %.6e", system.Time, system.Energy())
	for _, p := range system.Particles {
		output += fmt.Sprintf(", %.6e, %.6e", p.Position.X, p.Position.Y)
		output += fmt.Sprintf(", %.6e, %.6e", p.Velocity.X, p.Velocity.Y)
		output += fmt.Sprintf(", %.6e, %.6e", p.Force.X, p.Force.Y)
		output += fmt.Sprintf(", %.6e, %.6e", p.Force_gradient.X, p.Force_gradient.Y)
		output += fmt.Sprintf(", %.6e", p.Kinetic_energy())
	}
	output += "\n"
	return
}
