package physics

import (
	"fmt"
	"math"

	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

func (lj Lennard_jones_t) String() string {
	return fmt.Sprintf("Lennard Jones potential {\nsigma   = %.2e\nepsilon = %.2e\n}", lj.implementation.sigma, lj.implementation.epsilon)
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

func (lj Lennard_jones_t) Calculate_potentials(particles []Particle_t) float64 {
	return Calculate_inter_particle_potentials(&lj.implementation, particles)
}

func (lj Lennard_jones_t) Calculate_forces(particles []Particle_t) {
	Reset_forces(particles)
	Calculate_inter_particle_forces(&lj.implementation, particles)
}

func (lj Lennard_jones_t) Calculate_force_gradients(particles []Particle_t) {
	Reset_force_gradients(particles)
	Calculate_inter_particle_forces(&lj.implementation, particles)
}

type Lennard_jones_potential_t struct {
	sigma, epsilon float64
}

func (lj *Lennard_jones_potential_t) potential_between_p_and_q(p, q *Particle_t) float64 {
	distance := r2.Norm(r2.Sub(p.Position, q.Position))
	return lj.epsilon * LJ_j(distance/lj.sigma)
}

func (lj *Lennard_jones_potential_t) force_on_p_from_q(p, q *Particle_t) vector.Vec {
	var (
		direction = r2.Sub(p.Position, q.Position)
		distance  = r2.Norm(direction)
	)
	return r2.Scale(-lj.epsilon/lj.sigma*LJ_j_p(distance/lj.sigma), direction)
}

func (lj *Lennard_jones_potential_t) force_gradient_on_p_from_q(p, q *Particle_t) (force_gradient_q_on_p vector.Vec) {
	var ()
	return vector.Null()
}

var (
	LJ_x0   = 1
	LJ_x1   = math.Pow(2.0, 1.0/6.0)      // 1,122462
	LJ_x2   = math.Pow(26.0/7.0, 1.0/6.0) // 1,2444551
	LJ_j1   = 1.0
	LJ_j1pp = 12.0 * math.Pow(4.0, 3.0)
)

func LJ_j(x float64) float64 {
	xm6 := math.Pow(x, -6.0)
	return -4.0 * xm6 * (1.0 - xm6)
}

func LJ_j_p(x float64) float64 {
	xm6 := math.Pow(x, -6.0)
	return 24.0 * xm6 / x * (1.0 - 2.0*xm6)
}

func LJ_j_pp(x float64) float64 {
	xm6 := math.Pow(x, -6.0)
	return 168.0 * xm6 / (x * x) * (1.0 - 26.0/7.0*xm6)
}
