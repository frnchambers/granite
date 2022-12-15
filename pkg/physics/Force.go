package physics

import (
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

type Force_t interface {
	Calculate_potentials(particles []Particle_t) float64
	Calculate_forces(particles []Particle_t)
	Calculate_force_gradients(particles []Particle_t)
}

type Interaction_force_t interface {
	potential_between_p_and_q(p, q *Particle_t) float64
	force_on_p_from_q(p, q *Particle_t) vector.Vec
	force_gradient_on_p_from_q(p, q *Particle_t) vector.Vec
}

type External_force_t interface {
	potential(p *Particle_t) float64
	force(p *Particle_t) vector.Vec
	force_gradient(p *Particle_t) vector.Vec
}

func Calculate_inter_particle_potentials(interaction Interaction_force_t, particles []Particle_t) (V float64) {
	V = 0.0
	n_particles := len(particles)
	for i := range particles {
		for j := i + 1; j < n_particles; j++ {
			p := &particles[i]
			q := &particles[j]

			V += interaction.potential_between_p_and_q(p, q)

		}
	}
	return
}

func Calculate_inter_particle_forces(interaction Interaction_force_t, particles []Particle_t) {
	n_particles := len(particles)
	for i := range particles {
		for j := i + 1; j < n_particles; j++ {
			p := &particles[i]
			q := &particles[j]

			force := interaction.force_on_p_from_q(p, q)

			p.Force = r2.Add(p.Force, force)
			q.Force = r2.Sub(q.Force, force)
		}
	}
}

func Calculate_inter_particle_force_gradients(interaction Interaction_force_t, particles []Particle_t) {
	n_particles := len(particles)
	for i := range particles {
		for j := i + 1; j < n_particles; j++ {
			p := &particles[i]
			q := &particles[j]

			force_gradient := interaction.force_gradient_on_p_from_q(p, q)

			p.Force_gradient = r2.Add(p.Force_gradient, force_gradient)
			q.Force_gradient = r2.Sub(q.Force_gradient, force_gradient)
		}
	}
}

func Calculate_external_potentials(external External_force_t, particles []Particle_t) (V float64) {
	V = 0
	for i := range particles {
		V += external.potential(&particles[i])
	}
	return
}

func Calculate_external_forces(external External_force_t, particles []Particle_t) {
	for i := range particles {
		p := &particles[i]
		p.Force = r2.Add(p.Force, external.force(p))
	}
}

func Calculate_external_force_gradients(external External_force_t, particles []Particle_t) {
	for i := range particles {
		p := &particles[i]
		p.Force_gradient = r2.Add(p.Force_gradient, external.force_gradient(p))
	}
}

func Reset_forces(particles []Particle_t) {
	for i := range particles {
		particles[i].Reset_force_to_zero()
	}
}

func Reset_force_gradients(particles []Particle_t) {
	for i := range particles {
		particles[i].Reset_force_gradient_to_zero()
	}
}
