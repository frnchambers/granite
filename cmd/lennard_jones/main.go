package main

import (
	"fmt"
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/integrator"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/plot_p5"
	"github.com/granite/pkg/random"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
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

	sigma, epsilon := 1.0, 1.0
	n_particles := 2
	box_size := 2.0

	dt := 1.0e-2

	n_trails := 10
	dot_size := box_size * 1.0e-2

	initialise_molecules(sigma, epsilon, n_particles, box_size, dt)
	initialise_integrator(dt)
	initialise_window(box_size)
	initialise_draw_objects(dot_size, n_trails)

	output_variables()

	p5.Run(setup, draw_frame)
}

func initialise_molecules(sigma, epsilon float64, n_particles int, box_size, timestep float64) {

	max_position := box_size / 2.0
	// max_velocity := 1.0e-2 * box_size / timestep
	max_velocity := 0.0

	particles := make([]physics.Particle_t, n_particles)
	for i := range particles {
		particles[i] = new_random_molecule(max_position, max_velocity)
	}

	system = physics.System_t{Force: physics.New_lennard_jones(sigma, epsilon), Particles: particles, Time: 0.0}

	centre_of_mass := system.Centre_of_mass()
	average_velocity := system.Average_velocity()

	for i := range system.Particles {
		particles[i].Position = r2.Sub(particles[i].Position, centre_of_mass)
		particles[i].Velocity = r2.Sub(particles[i].Velocity, average_velocity)
	}
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

	// system.Time += sim.Step_time
	// fmt.Println("t =", system.Time)
}
