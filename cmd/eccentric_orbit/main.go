package main

import (
	"fmt"
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/integrator"
	"github.com/granite/pkg/kepler"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/plot_p5"
	"gonum.org/v1/gonum/spatial/r2"
)

/*

By definition:
  mu    := m1 m2 / (m1 + m2),
  alpha := m2 / m1,

Use: m2 = alpha m1, to remove m2 for:
  mu = alpha m1 / (1 + alpha),

Solve for m1:
  m1 = mu (1 + alpha) / alpha,
  m1 = mu (1 + alpha).

By definition: r := r2 - r1.

Can choose centre of mass frame such that,
  c := (m1 r1 + m2 r2) / (m1 + m2) = 0,
so that,
  0 = m1 r1 + m2 (r + r1)

Useful to note,
           m2 / (m1 + m2) = alpha / (1 + alpha),
and,
  1 - alpha / (1 + alpha) = 1 / (1 + alpha).

Solve for position:
  r1 = - r alpha / (1 + alpha),
  r2 = r / (1 + alpha).

*/

var (
	system physics.System_t

	stepper    integrator.Stepper_t
	timestep   float64
	step_count int

	dimensions  plot_p5.Window_dimensions_t
	save_frames bool

	dots       []plot_p5.Dot_t
	trails     []plot_p5.Trail_t
	velocities []plot_p5.Arrow_t

	// energy      plot_p5.Trail_t
	// init_energy float64

	background_col = color.Black
)

func main() {

	a, ecc, period, alpha := 1.0, 0.0, 1.0, 10.0e-1
	step_size := 3.0e-4 / period
	trail_fraction := 1.0
	n_trails := int(period / step_size * trail_fraction)

	initialise_bodies(a, ecc, period, alpha)
	initialise_integrator(step_size)
	initialise_window(a)
	initialise_draw_objects(n_trails)

	output_variables()

	p5.Run(setup, draw_frame)
}

func setup() {
	p5.PhysCanvas(
		dimensions.Pixels_x, dimensions.Pixels_y,
		dimensions.X_min, dimensions.X_max,
		dimensions.Y_min, dimensions.Y_max)
	p5.Background(background_col)
}

func initialise_bodies(a, ecc, period, alpha float64) {

	// centre_of_mass := vector.Vec{X: 0, Y: 0}

	orbit := kepler.New_elliptical_orbit(a, ecc, period)

	r := kepler.Position_along_elliplse(0.0, &orbit)
	v := kepler.Velocity_along_ellipse(0.0, &orbit)

	m1, m2 := orbit.Mu*(1.0+alpha)/alpha, orbit.Mu*(1.0+alpha)
	pos_1, pos_2 := r2.Scale(-alpha/(1.0+alpha), r), r2.Scale(1.0/(1.0+alpha), r)
	vel_1, vel_2 := r2.Scale(-alpha/(1.0+alpha), v), r2.Scale(1.0/(1.0+alpha), v)

	particles := make([]physics.Particle_t, 2)
	particles[0] = physics.New_particle("", m1/physics.G, pos_1, vel_1)
	particles[1] = physics.New_particle("", m2/physics.G, pos_2, vel_2)

	system = physics.System_t{Force: &physics.Gravity_interparticle_t{}, Particles: particles}
}

func initialise_draw_objects(n_trails int) {

	particle_col := color.RGBA{R: 255, G: 95, B: 26, A: 255}
	dot_size := 1.5e-2

	dots = plot_p5.Dots_from_system(system.Particles, particle_col, dot_size)
	trails = plot_p5.Trails_from_system(system.Particles, particle_col, n_trails)

	velocities = plot_p5.Velocities_from_system(system.Particles, particle_col)
}

func initialise_window(a float64) {
	physical_width := 2.0 * a
	center_width_frac, center_height_frac := 0.5, 0.5
	dimensions = plot_p5.New_dimensions(1200, 1000, physical_width, center_width_frac, center_height_frac)
	fmt.Println("Window: ", dimensions)
}

// func initialise_integrator(period float64, steps_per_period int) {
func initialise_integrator(step_size float64) {
	stepper = integrator.New_stepper(integrator.Default_O6_algorithm())
	timestep = step_size
}

func output_variables() {
	fmt.Print("system: ", system, "\n")
	fmt.Print("stepper: ", stepper, "\n")
}

func draw_frame() {

	stepper.Run(&system, timestep)

	plot_p5.Update_dots(dots, system.Particles)
	plot_p5.Update_trails(trails, system.Particles)
	// plot_p5.Update_velocities(velocities, system.Particles)

	for i := range dots {
		dots[i].Plot()
		trails[i].Plot()
		// velocities[i].Plot()
	}
}
