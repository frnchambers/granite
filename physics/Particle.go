package physics

import (
	"fmt"
	"math"

	"github.com/granite/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

type Particle_t struct {
	Name           string
	Mass           float64
	Position       vector.Vec
	Velocity       vector.Vec
	Force          vector.Vec
	Force_gradient vector.Vec
}

func (p *Particle_t) Reset_force_to_zero() {
	p.Force = vector.Vec{X: 0, Y: 0}
}

func (p *Particle_t) Reset_force_gradient_to_zero() {
	p.Force_gradient = vector.Vec{X: 0, Y: 0}
}

func (p *Particle_t) Increment_position_using_velocity(timestep float64) {
	p.Position = r2.Add(p.Position, r2.Scale(timestep, p.Velocity))
}

func (p *Particle_t) Increment_velocity_using_force(timestep float64) {
	p.Velocity = r2.Add(p.Velocity, r2.Scale(timestep/p.Mass, p.Force))
}

func (p *Particle_t) Increment_velocity_using_force_gradient(timestep float64) {
	p.Velocity = r2.Add(p.Velocity, r2.Scale(timestep/p.Mass, p.Force_gradient))
}

func (p *Particle_t) Increment_velocity_using_force_and_force_gradient(timestep_force, timestep_fg float64) {
	p.Velocity = r2.Add(p.Velocity, r2.Scale(timestep_force/p.Mass, p.Force))
	p.Velocity = r2.Add(p.Velocity, r2.Scale(timestep_fg/p.Mass, p.Force_gradient))
}

func New_particle(name string, mass float64, initial_pos, initial_vel vector.Vec) Particle_t {
	return Particle_t{
		Name:           name,
		Mass:           mass,
		Position:       initial_pos,
		Velocity:       initial_vel,
		Force:          vector.Vec{},
		Force_gradient: vector.Vec{},
	}
}

func New_blank_particle() Particle_t {
	return Particle_t{
		Name:           "",
		Mass:           0.0,
		Position:       vector.Vec{},
		Velocity:       vector.Vec{},
		Force:          vector.Vec{},
		Force_gradient: vector.Vec{},
	}
}

func (particle Particle_t) String() (output string) {
	output = "Particle_t {\n"
	output += fmt.Sprintf("Name =            %s\n", particle.Name)
	output += fmt.Sprintf("Mass =            %.2e\n", particle.Mass)
	output += fmt.Sprintf("Position =        %v\n", particle.Position)
	output += fmt.Sprintf("Velocity =        %v\n", particle.Velocity)
	output += fmt.Sprintf("Speed =           %.2e\n", r2.Norm(particle.Velocity))
	output += fmt.Sprintf("Angle betwee x+ = %.2e pi\n", math.Acos(r2.Cos(particle.Velocity, vector.Vec{X: 1, Y: 0}))/math.Pi)
	output += fmt.Sprintf("Force =           %v\n", particle.Force)
	output += fmt.Sprintf("Force_gradient =  %v\n", particle.Force_gradient)
	output += "}"
	return
}
