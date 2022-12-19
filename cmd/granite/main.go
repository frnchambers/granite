package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/comparison"
	"github.com/granite/pkg/integrator"
	"github.com/granite/pkg/kepler"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/plot_p5"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

const (
	edge_pixel_count = 1000
)

var (
	orbit  kepler.Elliptical_orbit_t
	system physics.System_t

	sim plot_p5.Simulation_t

	solar_pulse plot_p5.Pulse_t
	dots        []plot_p5.Dot_t
	pulses      []plot_p5.Pulse_t
	flares      []plot_p5.Flare_t
	trails      []plot_p5.Trail_t
	velocities  []plot_p5.Arrow_t

	stepper    integrator.Stepper_t
	step_count = 0
)

// func initialise_satellite(
// 	orbit *kepler.Orbit_t,
// 	time_to_perihelion, axis_angle float64,
// ) (particle physics.Particle_t) {

// 	phi := 0.0
// 	if !comparison.Float64_equality(time_to_perihelion, 0.0) {
// 		var err error = nil
// 		phi, err = kepler.Phi_for_time_to_perihelion(time_to_perihelion, orbit)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	particle = kepler.New_satellite(phi, 1.0, orbit)

// 	if !comparison.Float64_equality(axis_angle, 0.0) {
// 		kepler.Rotate_orbit(&particle, axis_angle)
// 	}

// 	return
// }

func initialise_satellites(
	a, ecc, period float64,
	n_steps, n_trails int,
	axis_offset_angle, particle_offset_angle float64,
	offset_times []float64,
) {
	solar_position := vector.Vec{X: 0, Y: 0}

	orbit = kepler.New_elliptical_orbit(a, ecc, period)
	n_particles := len(offset_times)
	sim = kepler.New_simulation_parameters(n_steps, n_trails, &orbit)

	particles := make([]physics.Particle_t, n_particles)
	for i, time := range offset_times {
		phi := 0.0
		if !comparison.Float64_equality(time, 0.0) {
			var err error = nil
			phi, err = kepler.Phi_for_time_to_perihelion(time, &orbit)
			if err != nil {
				panic(err)
			}
		}
		particles[i] = kepler.New_satellite(phi, 1.0, &orbit)
		if !comparison.Float64_equality(axis_offset_angle, 0.0) &&
			!comparison.Float64_equality(particle_offset_angle, 0.0) {
			kepler.Rotate_orbit(&particles[i], -axis_offset_angle-float64(i)*particle_offset_angle)
		}
	}

	system = physics.System_t{Force: kepler.New_massive_body(&orbit), Particles: particles}

	solar_pulse = plot_p5.New_pulse(color.RGBA{R: 246, G: 244, B: 129, A: 255}, n_steps, 2.0e-2, 1.0e-1)
	solar_pulse.Update_position(solar_position)
	solar_pulse.Reset_time(-int(offset_times[0] / sim.Step_time))

	dots = plot_p5.Dots_from_system(system.Particles, sim.Dot_size)
	width := 5.0e-1
	flares = plot_p5.Flares_from_system(system.Particles, sim.Dot_size*5.0e-1, sim.Dot_size*1.5, width)
	pulses = plot_p5.Pulses_from_system(system.Particles, n_steps, sim.Dot_size/2.0, sim.Dot_size*2.0)
	for i, time := range offset_times {
		pulses[i].Reset_time(-int(time / sim.Step_time))
	}
	trails = plot_p5.Trails_from_system(system.Particles, sim.Trail_length)
	velocities = plot_p5.Velocities_from_system(system.Particles)
}

func output_variables() {
	fmt.Print("sol = ", solar_pulse, "\n")
	fmt.Print("orbit = ", orbit, "\n")
	fmt.Print("system: ", system, "\n")
	fmt.Printf("starting distance from body = %.2e\n", r2.Norm(r2.Sub(solar_pulse.Position(), system.Particles[0].Position)))
	fmt.Print("stepper: ", stepper, "\n")
}

func main() {
	run_simulation()
}

func granite_settings() {
	a, ecc, period := 1.0, 0.7, 1.0

	// n_steps, n_trails := 360, 7
	// axis_offset_angle, particle_offset_angle := math.Pi/6.0, math.Pi/64.0

	// time_lag := period / 16.0
	// offset_times := []float64{
	// 	0.0 * time_lag,
	// 	1.0 * time_lag,
	// 	2.0 * time_lag,
	// }

	// initialise_satellites(a, ecc, period, n_steps, n_trails, axis_offset_angle, particle_offset_angle, offset_times)

	var beats_per_second float64 = 95.0 / 60.0
	steps_per_second := 90
	beats_per_period := 8

	steps_per_period := int(float64(beats_per_period*steps_per_second) / beats_per_second)

	fmt.Println("steps per second =", steps_per_period)

	n_trails := 10
	axis_offset_angle, particle_offset_angle := math.Pi/6.0, math.Pi/64.0

	lead_time := period / 4.0
	time_lag := period / 16.0
	offset_times := []float64{
		0.0*time_lag + lead_time,
		1.0*time_lag + lead_time,
		2.0*time_lag + lead_time,
	}

	initialise_satellites(a, ecc, period, steps_per_period, n_trails, axis_offset_angle, particle_offset_angle, offset_times)

	stepper = integrator.New_stepper(integrator.Default_O4_algorithm())
}

func run_simulation() {
	granite_settings()
	// highly_eccentric_settings()
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
	step_count += 1

	plot_p5.Update_dots(dots, system.Particles)
	plot_p5.Update_pulses(pulses, system.Particles)
	plot_p5.Update_flares(flares, system.Particles)
	plot_p5.Update_trails(trails, system.Particles)
	plot_p5.Update_velocities(velocities, system.Particles)

	solar_pulse.Update_time()
	solar_pulse.Plot()

	for i := range dots {
		// dots[i].Plot()
		// pulses[i].Plot()
		flares[i].Plot()
		trails[i].Plot()
		// velocities[i].Plot()
	}

	// filename := fmt.Sprintf("frames/frame_%.6d.png", step_count)
	// p5.Screenshot(filename)

}
