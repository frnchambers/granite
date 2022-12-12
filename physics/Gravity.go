package physics

import (
	"fmt"

	"github.com/granite/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

const (
	G                   = 6.67430e-11
	Gravitational_units = "(L, M, T): (AU, M_sol, days)"
)

func (Gravity_interparticle_t) String() string {
	return "Gravity from interparticle forces"
}
func (g Gravity_massive_body_t) String() (output string) {
	return fmt.Sprintf("Gravity from central massive body:\n%v", g.Body())
}

type Gravity_interparticle_t struct{}

func (Gravity_interparticle_t) Calculate_forces(particles []Particle_t) {
	Reset_forces(particles)
	Calculate_inter_particle_forces(Gravitational_interaction_t{}, particles)
}

func (Gravity_interparticle_t) Calculate_force_gradients(particles []Particle_t) {
	Reset_force_gradients(particles)
	Calculate_inter_particle_forces(Gravitational_interaction_t{}, particles)
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

type Gravitational_interaction_t struct{}

func (Gravitational_interaction_t) force_on_p_from_q(p, q *Particle_t) (force_q_on_p vector.Vec) {
	var (
		direction = r2.Sub(p.Position, q.Position)
		distance  = r2.Norm(direction)
	)
	force_q_on_p = r2.Scale(-G*p.Mass*q.Mass/(distance*distance*distance), direction)
	return
}

func (Gravitational_interaction_t) force_gradient_on_p_from_q(p, q *Particle_t) (force_gradient_q_on_p vector.Vec) {
	var (
		dr           = r2.Sub(p.Position, q.Position)
		da           = r2.Sub(r2.Scale(1.0/p.Mass, p.Force), r2.Scale(1.0/q.Mass, q.Force))
		dr_dot_da    = r2.Dot(dr, da)
		distance     = r2.Norm(dr)
		dist_squared = distance * distance
		dist_cubed   = distance * distance * distance
	)
	force_gradient_q_on_p = r2.Sub(da, r2.Scale(3.0*dr_dot_da/dist_squared, dr))
	force_gradient_q_on_p = r2.Scale(-2.0*G*p.Mass*q.Mass/(dist_cubed), force_gradient_q_on_p)
	return
}
