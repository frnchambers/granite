package physics

import (
	"math"

	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

func (Lennard_jones_t) String() string {
	return "Lennard Jones potential {\nsigma   = %.2e\nepsilon = %.2e\n}"
}

type Lennard_jones_t struct {
	implementation Lennard_jones_potential_t
}

func New_lennard_jones(sigma, epsilon float64) Lennard_jones_t {
	return Lennard_jones_t{
		implementation: Lennard_jones_potential_t{
			sigma:   sigma,
			epsilon: epsilon,
		},
	}
}

func (lj *Lennard_jones_t) Calculate_forces(particles []Particle_t) {
	Reset_forces(particles)
	Calculate_inter_particle_forces(&lj.implementation, particles)
}

func (lj *Lennard_jones_t) Calculate_force_gradients(particles []Particle_t) {
	Reset_force_gradients(particles)
	Calculate_inter_particle_forces(&lj.implementation, particles)
}

type Lennard_jones_potential_t struct {
	sigma, epsilon float64
}

func (lj *Lennard_jones_potential_t) force_on_p_from_q(p, q *Particle_t) vector.Vec {
	var (
		direction = r2.Sub(p.Position, q.Position)
		distance  = r2.Norm(direction)
		sigma_r   = lj.sigma / distance
		sigma_r_6 = math.Pow(sigma_r, 6)
	)
	return r2.Scale(24.0*lj.epsilon/lj.sigma*sigma_r_6*(1.0-2.0*sigma_r_6)/distance, direction)
}

func (lj *Lennard_jones_potential_t) force_gradient_on_p_from_q(p, q *Particle_t) (force_gradient_q_on_p vector.Vec) {
	var ()
	return vector.Null()
}
