package physics

import "fmt"

type System_t struct {
	Force     Force_t
	Particles []Particle_t
	Time      float64
}

func (system *System_t) N_particles() int {
	return len(system.Particles)
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
	output += fmt.Sprintf("Force: %v\nParticles: [\n", system.Force)
	for _, p := range system.Particles {
		output += p.String() + ",\n"
	}
	output += "]"
	return
}

func (system System_t) As_row() (output string) {
	output = "System_t:\n"
	output += fmt.Sprintf("Force: %v\nParticles: [\n", system.Force)
	for _, p := range system.Particles {
		output += p.String() + ",\n"
	}
	output += "]"
	return
}
