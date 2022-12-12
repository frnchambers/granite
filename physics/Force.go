package physics

import "gonum.org/v1/gonum/spatial/r2"

type Force_t interface {
	Calculate_forces(particles []Particle_t)
	Calculate_force_gradients(particles []Particle_t)
}

type Interaction_force_t interface {
	force_on_p_from_q(p, q *Particle_t) r2.Vec
	force_gradient_on_p_from_q(p, q *Particle_t) r2.Vec
}

type External_force_t interface {
	force(p *Particle_t) r2.Vec
	force_gradient(p *Particle_t) r2.Vec
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
