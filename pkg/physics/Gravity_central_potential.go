package physics

import (
	"fmt"

	"github.com/granite/pkg/vector"
)

func (g Gravity_massive_body_t) String() (output string) {
	return fmt.Sprintf("Gravity from central massive body:\n%v", g.Body())
}

type Gravity_massive_body_t struct{ Gravitation_massive_body Gravitation_massive_body_t }

func (g *Gravity_massive_body_t) Body_mass() float64 {
	return g.Gravitation_massive_body.body.Mass
}
func (g *Gravity_massive_body_t) Centre() vector.Vec {
	return g.Gravitation_massive_body.body.Position
}
func (g *Gravity_massive_body_t) Body() Particle_t {
	return g.Gravitation_massive_body.body
}

func New_massive_body(particle Particle_t) Gravity_massive_body_t {
	return Gravity_massive_body_t{
		Gravitation_massive_body: Gravitation_massive_body_t{
			body: New_particle(particle.Name, particle.Mass, particle.Position, vector.Null()),
		},
	}
}

func (g Gravity_massive_body_t) Calculate_forces(particles []Particle_t) {
	Reset_forces(particles)
	Calculate_external_forces(&g.Gravitation_massive_body, particles)
}

func (g Gravity_massive_body_t) Calculate_force_gradients(particles []Particle_t) {
	Reset_force_gradients(particles)
	Calculate_external_forces(&g.Gravitation_massive_body, particles)
}

type Gravitation_massive_body_t struct{ body Particle_t }

func (g *Gravitation_massive_body_t) Body() Particle_t {
	return g.body
}
func (g *Gravitation_massive_body_t) Body_mass() float64 {
	return g.body.Mass
}
func (g *Gravitation_massive_body_t) Body_position() vector.Vec {
	return g.body.Position
}

func (central_mass *Gravitation_massive_body_t) force(p *Particle_t) vector.Vec {
	return Gravitational_interaction_t{}.force_on_p_from_q(p, &central_mass.body)
}

func (central_mass *Gravitation_massive_body_t) force_gradient(p *Particle_t) vector.Vec {
	return Gravitational_interaction_t{}.force_gradient_on_p_from_q(p, &central_mass.body)
}
