package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/integrator"
	"github.com/granite/kepler"
	"github.com/granite/physics"
	"github.com/granite/plot"
	"gonum.org/v1/gonum/spatial/r2"
)

const (
	edge_pixel_count = 1000
)

var (
	orbit  kepler.Orbit_t
	system physics.System_t

	sim plot.Simulation_t

	solar_dot  plot.Dot_t
	dots       []plot.Dot_t
	trails     []plot.Trail_t
	velocities []plot.Arrow_t

	stepper integrator.Stepper_t
)

func initialise_satellite() {
	a := 1.0
	ecc := 0.7
	period := 1.0

	solar_position := r2.Vec{X: 0, Y: 0}
	orbit = kepler.New_orbit(a, ecc, period)

	offset_times := []float64{
		0.0 * orbit.Period / 16.0,
		1.0 * orbit.Period / 16.0,
		2.0 * orbit.Period / 16.0,
	}
	n_particles := len(offset_times)

	axis_offset_angle := math.Pi / 6.0
	particle_offset_angle := math.Pi / 64.0

	n_steps := 360
	n_trails := 5
	sim = kepler.New_simulation_parameters(n_steps, n_trails, &orbit)

	particles := make([]physics.Particle_t, n_particles)
	for i, time := range offset_times {
		phi, err := kepler.Phi_for_time_to_perihelion(time, &orbit)
		if err != nil {
			panic(err)
		}
		particles[i] = kepler.New_satellite(phi, &orbit)
		kepler.Rotate_orbit(&particles[i], -axis_offset_angle-float64(i)*particle_offset_angle)
	}

	system = physics.System_t{Force: kepler.New_massive_body(&orbit), Particles: particles}

	solar_dot = plot.New_static_dot(color.RGBA{R: 246, G: 244, B: 129, A: 255}, 5.0e-2, solar_position)
	dots, trails, velocities = plot.From_system(system.Particles, sim.Dot_size, sim.Trail_length)
	stepper = integrator.New_stepper(integrator.Default_algorithm())
}

func output_variables() {
	fmt.Print("sol = ", solar_dot, "\n")
	fmt.Print("system: ", system, "\n")
	fmt.Printf("starting distance from body = %.2e\n", r2.Norm(r2.Sub(solar_dot.Position(), system.Particles[0].Position)))
	fmt.Print("stepper: ", stepper, "\n")
}

func main() {
	run_simulation()
	// example_plot()
}

func run_simulation() {
	initialise_satellite()
	output_variables()
	p5.Run(setup, draw_frame)
}

func setup() {
	p5.PhysCanvas(edge_pixel_count, edge_pixel_count,
		sim.X_min, sim.X_max, sim.Y_min, sim.Y_max)
	p5.Background(color.Black)
}

func draw_frame() {

	stepper.Run(&system, sim.Step_time)

	plot.Update_dots(dots, system.Particles)
	plot.Update_trails(trails, system.Particles)
	plot.Update_velocities(velocities, system.Particles)

	// p5.Stroke(color.White)
	// p5.Fill(color.Black)
	// p5.Ellipse(-orbit.Linear_eccentricity, 0.0, 2.0*orbit.Semi_major, 2.0*orbit.Semi_minor)

	solar_dot.Plot()

	for i := range dots {
		dots[i].Plot()
		trails[i].Plot()
		// velocities[i].Plot()
	}
}
