package plot_p5

import (
	"image/color"

	"github.com/granite/physics"
	"gonum.org/v1/gonum/spatial/r2"
)

const (
	DEFAULT_VELOCITY_SCALE = 0.2
)

func Update_dots(dots []Dot_t, particles []physics.Particle_t) {
	for i := range particles {
		dots[i].Update(particles[i].Position)
	}
}

func Update_trails(trails []Trail_t, particles []physics.Particle_t) {
	for i := range particles {
		trails[i].Update(particles[i].Position)
	}
}

func Update_velocities(arrows []Arrow_t, particles []physics.Particle_t) {
	for i := range particles {
		normalised_velocity := r2.Scale(DEFAULT_VELOCITY_SCALE, particles[i].Velocity)
		arrows[i].Update(particles[i].Position, r2.Add(particles[i].Position, normalised_velocity))
	}
}

func From_system(particles []physics.Particle_t, dot_size float64, trail_length int) (
	dots []Dot_t,
	trails []Trail_t,
	velocities []Arrow_t,
) {
	n_particles := len(particles)
	trails = make([]Trail_t, 0, n_particles)
	dots = make([]Dot_t, 0, n_particles)
	velocities = make([]Arrow_t, 0, n_particles)

	for i := range particles {
		dots = append(dots, New_dot(
			&particles[i].Position,
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
			dot_size,
		))
		trails = append(trails, New_trail(
			&particles[i].Position,
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
			trail_length,
		))
		velocities = append(velocities, New_arrow(
			&particles[i].Position,
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
		))
	}
	return
}
