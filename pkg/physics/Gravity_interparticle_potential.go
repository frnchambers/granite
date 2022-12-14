package physics

import (
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

const (
	G                   = 6.67430e-11
	Gravitational_units = "(L, M, T): (AU, M_sol, days)"
)

func (Gravity_interparticle_t) String() string {
	return "Gravity from interparticle forces"
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

type Gravitational_interaction_t struct{}

func (Gravitational_interaction_t) force_on_p_from_q(p, q *Particle_t) vector.Vec {
	var (
		direction = r2.Sub(p.Position, q.Position)
		distance  = r2.Norm(direction)
	)
	return r2.Scale(-G*p.Mass*q.Mass/(distance*distance*distance), direction)
}

func (Gravitational_interaction_t) force_gradient_on_p_from_q(p, q *Particle_t) (force_gradient vector.Vec) {
	var (
		dr           = r2.Sub(p.Position, q.Position)
		da           = r2.Sub(r2.Scale(1.0/p.Mass, p.Force), r2.Scale(1.0/q.Mass, q.Force))
		dr_dot_da    = r2.Dot(dr, da)
		distance     = r2.Norm(dr)
		dist_squared = distance * distance
		dist_cubed   = distance * distance * distance
	)
	force_gradient = r2.Sub(da, r2.Scale(3.0*dr_dot_da/dist_squared, dr))
	force_gradient = r2.Scale(-2.0*G*p.Mass*q.Mass/(dist_cubed), force_gradient)
	return
}
