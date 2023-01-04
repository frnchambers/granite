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
)

const (
	edge_pixel_count = 1000
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

func initialise_molecules() {

	sigma, epsilon := 1.0, 1.0

	n_trails := 10
	n_particles := 5

	dt := 1.0e-3

	window_size := 20.0
	dot_size := window_size * 1.0e-2

	max_velocity := 1.0e-2 * window_size

	sim = plot_p5.Window_dimensions_t{
		Trail_length: n_trails,
		Dot_size:     dot_size,
		X_min:        -window_size,
		X_max:        window_size,
		Y_min:        -window_size,
		Y_max:        window_size,
		Step_time:    dt,
	}

	particles := make([]physics.Particle_t, n_particles)
	for i := range particles {
		particles[i] = new_random_molecule(window_size/2.0, max_velocity)
	}

	system = physics.System_t{Force: physics.New_lennard_jones(sigma, epsilon), Particles: particles, Time: 0.0}

	dots = plot_p5.Dots_from_system(system.Particles, sim.Dot_size)
	trails = plot_p5.Trails_from_system(system.Particles, sim.Trail_length)
	velocities = plot_p5.Velocities_from_system(system.Particles)

	stepper = integrator.New_stepper(integrator.Velocity_verlet_algorithm())
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

func main() {
	run_simulation()
}

func run_simulation() {
	initialise_molecules()
	p5.Run(setup, draw_frame)
}

func setup() {
	p5.PhysCanvas(edge_pixel_count, edge_pixel_count,
		sim.X_min, sim.X_max, sim.Y_min, sim.Y_max)
	p5.Background(color.Black)
}

func draw_frame() {

	stepper.Run(&system, sim.Step_time)

	plot_p5.Update_dots(dots, system.Particles)
	// plot_p5.Update_trails(trails, system.Particles)
	plot_p5.Update_velocities(velocities, system.Particles)

	for i := 0; i < system.N_particles(); i++ {
		dots[i].Plot()
		// trails[i].Plot()
		velocities[i].Plot()
	}

	// system.Time += sim.Step_time
	// fmt.Println("t =", system.Time)
}
