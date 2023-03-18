package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/integrator"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/plot_p5"
	"github.com/granite/pkg/random"
	"github.com/granite/pkg/vector"
)

var (
	system physics.System_t

	stepper integrator.Stepper_t

	timestep   float64
	step_count int

	dimensions  plot_p5.Window_dimensions_t
	save_frames bool

	dots       []plot_p5.Dot_t
	trails     []plot_p5.Trail_t
	velocities []plot_p5.Arrow_t

	background_col = color.Black
)

func main() {

	mass, sigma, epsilon := 1.0, 1.0, 1.0
	box_size := 2.0

	omega_sq := epsilon * physics.LJ_j1pp / (mass * sigma * sigma)
	period := 2.0 * math.Pi / math.Sqrt(omega_sq)
	dt := 5.0e-2 * period

	n_trails := 10
	dot_size := box_size * 2.0e-2

	initialise_molecules(mass, sigma, epsilon)
	initialise_integrator(dt)
	initialise_window(box_size)
	initialise_draw_objects(dot_size, n_trails)

	output_variables()

	p5.Run(setup, draw_frame)
}

func initialise_molecules(mass, sigma, epsilon float64) {

	r_0 := sigma * physics.LJ_x2 * (1. + 1.9)
	n_particles := 2

	particles := make([]physics.Particle_t, n_particles)

	particles[0] = physics.New_particle("", mass, vector.New(r_0/2.0, 0.0), vector.Null())
	particles[1] = physics.New_particle("", mass, vector.New(-r_0/2.0, 0.0), vector.Null())

	system = physics.System_t{Force: physics.New_lennard_jones(sigma, epsilon), Particles: particles, Time: 0.0}
}

func initialise_integrator(step_size float64) {
	timestep = step_size
	stepper = integrator.New_stepper(integrator.Velocity_verlet_algorithm())
}

func initialise_window(box_size float64) {
	physical_width := 2.0 * box_size
	center_width_frac, center_height_frac := 0.5, 0.5
	dimensions = plot_p5.New_dimensions(1200, 1000, physical_width, center_width_frac, center_height_frac)
	fmt.Println("Window: ", dimensions)
}

func initialise_draw_objects(dot_size float64, trail_length int) {
	col := color.RGBA{R: 255, G: 95, B: 26, A: 255}
	dots = plot_p5.Dots_from_system(system.Particles, col, dot_size)
	trails = plot_p5.Trails_from_system(system.Particles, col, trail_length)
	velocities = plot_p5.Velocities_from_system(system.Particles, col)
}

func new_random_molecule(box_size, max_speed float64) physics.Particle_t {
	return physics.New_particle("", 1.0, random.Position(box_size), random.Velocity(max_speed))
}

func new_random_static_molecule(box_size, max_speed float64) physics.Particle_t {
	return physics.New_particle("", 1.0, random.Position(box_size), vector.Null())
}

func output_variables() {
	fmt.Print("system: ", system, "\n")
	fmt.Print("stepper: ", stepper, "\n")
	fmt.Print("timestep: ", timestep, "\n")
}

func setup() {
	p5.PhysCanvas(
		dimensions.Pixels_x, dimensions.Pixels_y,
		dimensions.X_min, dimensions.X_max,
		dimensions.Y_min, dimensions.Y_max)
	p5.Background(background_col)
}

func draw_frame() {

	stepper.Run(&system, timestep)

	plot_p5.Update_dots(dots, system.Particles)
	// plot_p5.Update_trails(trails, system.Particles)
	// plot_p5.Update_velocities(velocities, system.Particles)

	for i := 0; i < system.N_particles(); i++ {
		dots[i].Plot()
		// trails[i].Plot()
		// velocities[i].Plot()
	}

	// fmt.Print(system.As_row())
}
