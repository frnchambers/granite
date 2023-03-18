package physics

import (
	"fmt"

	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
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

func (system *System_t) Total_mass() (mass float64) {
	mass = 0.0
	for i := range system.Particles {
		mass += system.Particles[i].Mass
	}
	return
}

func (system *System_t) Centre_of_mass() (com vector.Vec) {
	com = vector.Null()
	M := system.Total_mass()
	for i := range system.Particles {
		com = r2.Add(com, r2.Scale(system.Particles[i].Mass/M, system.Particles[i].Position))
	}
	return
}

func (system *System_t) Average_velocity() (vel vector.Vec) {
	vel = vector.Null()
	N := float64(system.N_particles())
	for i := range system.Particles {
		vel = r2.Add(vel, system.Particles[i].Velocity)
	}
	vel = r2.Scale(1.0/N, vel)
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
	output += fmt.Sprintf("# Initial time:   %.6e\n# Initial energy: %.6e\n", system.Time, system.Energy())
	output += fmt.Sprintf("time, energy")
	for i := range system.Particles {
		output += fmt.Sprintf(", r_%d_x, r_%d_y, v_%d_x, v_%d_y, f_%d_x, f_%d_y, fg_%d_x, fg_%d_y, ke_%d", i, i, i, i, i, i, i, i, i)
	}
	output += "\n"
	return
}

func (system System_t) As_row() (output string) {
	output += fmt.Sprintf("%.6e, %.6e", system.Time, system.Energy())
	for _, p := range system.Particles {
		output += fmt.Sprintf(", %.6e, %.6e, %.6e, %.6e, %.6e, %.6e, %.6e, %.6e, %.6e",
			p.Position.X, p.Position.Y,
			p.Velocity.X, p.Velocity.Y,
			p.Force.X, p.Force.Y,
			p.Force_gradient.X, p.Force_gradient.Y,
			p.Kinetic_energy(),
		)
	}
	output += "\n"
	return
}
