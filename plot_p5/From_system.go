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

func Update_pulses(pulses []Pulse_t, particles []physics.Particle_t) {
	for i := range particles {
		pulses[i].Update(particles[i].Position)
	}
}

func Update_flares(flares []Flare_t, particles []physics.Particle_t) {
	for i := range particles {
		flares[i].Update(particles[i].Position)
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

func Dots_from_system(particles []physics.Particle_t, dot_size float64) (
	dots []Dot_t,
) {
	n_particles := len(particles)
	dots = make([]Dot_t, 0, n_particles)

	for range particles {
		dots = append(dots, New_dot(
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
			dot_size,
		))
	}
	return
}

func Pulses_from_system(particles []physics.Particle_t, cycle_lenth int, min_size, max_size float64) (
	pulses []Pulse_t,
) {
	n_particles := len(particles)
	pulses = make([]Pulse_t, 0, n_particles)

	for range particles {
		pulses = append(pulses, New_pulse(
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
			cycle_lenth, min_size, max_size,
		))
	}
	return
}

func Flares_from_system(particles []physics.Particle_t, min_size, max_size, width float64) (
	flares []Flare_t,
) {
	n_particles := len(particles)
	flares = make([]Flare_t, 0, n_particles)

	for range particles {
		flares = append(flares, New_flare_centre_zero(
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
			min_size, max_size, width,
		))
	}
	return
}

func Trails_from_system(particles []physics.Particle_t, trail_length int) (
	trails []Trail_t,
) {
	n_particles := len(particles)
	trails = make([]Trail_t, 0, n_particles)

	for range particles {
		trails = append(trails, New_trail(
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
			trail_length,
		))
	}
	return
}

func Velocities_from_system(particles []physics.Particle_t) (
	velocities []Arrow_t,
) {
	n_particles := len(particles)
	velocities = make([]Arrow_t, 0, n_particles)

	for range particles {
		velocities = append(velocities, New_arrow(
			color.RGBA{R: 223, G: 120, B: 036, A: 255},
		))
	}
	return
}
