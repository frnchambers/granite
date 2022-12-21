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
	"gonum.org/v1/gonum/spatial/r2"
)

var (
	sol    physics.Particle_t
	system physics.System_t

	stepper    integrator.Stepper_t
	timestep   float64
	step_count int

	dimensions  plot_p5.Window_dimensions_t
	save_output bool

	solar_pulse plot_p5.Pulse_t
	dots        []plot_p5.Dot_t
	pulses      []plot_p5.Pulse_t
	flares      []plot_p5.Flare_t
	trails      []plot_p5.Trail_t

	background_col = color.Black
)

func main() {

	frame_rate := 60
	beats_per_minute := 95

	a, ecc, period := 1.0, 0.7, 1.0

	lead_time := -period / 4.0
	axis_offset_angle := math.Pi / 6.0

	granite_settings(
		frame_rate, beats_per_minute,
		a, ecc, period,
		lead_time, axis_offset_angle)

	p5.Run(setup, draw_frame)
}

func granite_settings(
	frame_rate, beats_per_minute int,
	a, ecc, period,
	lead_time_per_period, axis_offset_angle float64) {

	orbit := kepler.New_elliptical_orbit(a, ecc, period)

	lead_time := period * lead_time_per_period
	time_lag := period / 16.0
	offset_times := []float64{
		lead_time + 0.0*time_lag,
		lead_time + 1.0*time_lag,
		lead_time + 2.0*time_lag,
	}

	particle_offset_angle := math.Pi / 64.0
	offset_angels := []float64{
		-axis_offset_angle - 0.0*particle_offset_angle,
		-axis_offset_angle - 1.0*particle_offset_angle,
		-axis_offset_angle - 2.0*particle_offset_angle,
	}

	initialise_satellites(a, ecc, period, offset_times, offset_angels)

	beats_per_period := 8
	steps_per_period := int(float64(beats_per_period*frame_rate) * 60.0 / float64(beats_per_minute))

	fmt.Println("steps per second =", steps_per_period)

	initialise_draw_objects(&orbit, frame_rate, offset_times)

	stepper = integrator.New_stepper(integrator.Default_O6_algorithm())
	timestep = period / float64(steps_per_period)

	center_height_pc := 0.3
	dimensions = New_dimensions_1920_1080(2*orbit.Semi_major, orbit.Eccentricity/2, center_height_pc)
	save_output = false

}

func New_dimensions_1920_1080(physical_width, center_width_pc, center_height_pc float64) plot_p5.Window_dimensions_t {
	tan_theta := 1080.0 / 1920.0
	return plot_p5.Window_dimensions_t{
		Pixels_x: 1920,
		Pixels_y: 1080,
		X_min:    -physical_width * (1.0 + center_width_pc),
		X_max:    physical_width * (1.0 - center_width_pc),
		Y_min:    -physical_width * tan_theta * (1.0 + center_height_pc),
		Y_max:    physical_width * tan_theta * (1.0 - center_height_pc),
	}
}

func output_variables() {
	fmt.Print("sol = ", solar_pulse, "\n")
	fmt.Print("system: ", system, "\n")
	fmt.Printf("starting distance from body = %.2e\n", r2.Norm(r2.Sub(solar_pulse.Position(), system.Particles[0].Position)))
	fmt.Print("stepper: ", stepper, "\n")
}

func initialise_satellites(
	a, ecc, period float64,
	offset_times, offset_angles []float64,
) {
	orbit := kepler.New_elliptical_orbit(a, ecc, period)
	fmt.Print("orbit = ", orbit, "\n")

	n_particles := len(offset_times)

	particles := make([]physics.Particle_t, n_particles)
	for i := range offset_times {
		time := offset_times[i]
		rotation := offset_angles[i]
		particles[i] = initialise_satellite(&orbit, 1.0, time, rotation)
	}

	fmt.Print("orbit = ", orbit, "\n")

	sol = kepler.New_massive_particle(&orbit)

	system = physics.System_t{
		Force:     kepler.New_massive_body(&orbit),
		Particles: particles,
		Time:      0.0,
	}

}

func initialise_satellite(
	orbit *kepler.Elliptical_orbit_t,
	mass, time_to_perihelion, axis_rotation float64,
) (particle physics.Particle_t) {

	phi := 0.0
	if !comparison.Float64_equality(time_to_perihelion, 0.0) {
		var err error = nil
		phi, err = kepler.Phi_for_time_to_perihelion(time_to_perihelion, orbit)
		if err != nil {
			panic(err)
		}
	}

	particle = kepler.New_satellite(phi, mass, orbit)

	if !comparison.Float64_equality(axis_rotation, 0.0) {
		kepler.Rotate_orbit(&particle, axis_rotation)
	}

	return
}

func initialise_draw_objects(
	orbit *kepler.Elliptical_orbit_t,
	frame_rate int,
	offset_times []float64,
) {

	sol_col := color.RGBA{R: 246, G: 244, B: 129, A: 255}
	sol_size, sol_max := 3.0e-2, 9.0e-2
	initialise_sol(sol_col, sol_size, sol_max, orbit.Period, offset_times[0])

	particle_col := color.RGBA{R: 223, G: 120, B: 036, A: 255}
	particle_size_relative_to_sol := 2.0e-1
	particle_size := sol_size * particle_size_relative_to_sol
	initialise_dots(particle_col, particle_size)

	particle_max := sol_max * particle_size_relative_to_sol
	initialise_pulses(particle_col, particle_size, particle_max, orbit.Period, offset_times)

	n_trails := int(float64(frame_rate) / 10.0)
	initialise_trails(particle_col, n_trails)

	flare_width := 5.0e-1
	initialise_flares(particle_col, particle_size, particle_max, flare_width)
}

func initialise_sol(
	col color.Color,
	min_size, max_size float64,
	period, start_time float64,
) {
	solar_pulse = plot_p5.New_pulse(col, period, min_size, max_size)
	solar_pulse.Update_position(sol.Position)
	solar_pulse.Reset_time(start_time)
}

func initialise_dots(col color.Color, dot_size float64) {
	dots = plot_p5.Dots_from_system(system.Particles, col, dot_size)
}

func initialise_flares(
	col color.Color,
	min_size, max_size, width float64,
) {
	flares = plot_p5.Flares_from_system(system.Particles, sol.Position, col, min_size, max_size, width)
}

func initialise_pulses(
	col color.Color,
	min_size, max_size float64,
	period float64, offset_times []float64,
) {
	pulses = plot_p5.Pulses_from_system(system.Particles, col, period, min_size, max_size)
	for i, time := range offset_times {
		pulses[i].Reset_time(time)
	}
}

func initialise_trails(col color.Color, trail_length int) {
	trails = plot_p5.Trails_from_system(system.Particles, col, trail_length)
}

func draw_frame() {

	stepper.Run(&system, timestep)
	step_count += 1

	plot_p5.Update_dots(dots, system.Particles)
	plot_p5.Update_pulses(pulses, system.Particles, timestep)
	plot_p5.Update_flares(flares, system.Particles)
	plot_p5.Update_trails(trails, system.Particles)

	solar_pulse.Update_time(timestep)
	solar_pulse.Plot()

	for i := range dots {
		// dots[i].Plot()
		// pulses[i].Plot()
		flares[i].Plot()
		trails[i].Plot()
		// velocities[i].Plot()
	}

	if save_output {
		filename := fmt.Sprintf("frames/frame_%.6d.png", step_count)
		p5.Screenshot(filename)
	}
}

func setup() {
	p5.PhysCanvas(dimensions.Pixels_x, dimensions.Pixels_y,
		dimensions.X_min, dimensions.X_max, dimensions.Y_min, dimensions.Y_max)
	p5.Background(background_col)
}
