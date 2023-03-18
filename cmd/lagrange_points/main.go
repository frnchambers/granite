package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/integrator"
	"github.com/granite/pkg/kepler"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/plot_p5"
	"github.com/granite/pkg/random"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

var (
	system physics.System_t

	stepper    integrator.Stepper_t
	timestep   float64
	step_count int

	dimensions  plot_p5.Window_dimensions_t
	save_frames bool

	dots   []plot_p5.Dot_t
	trails []plot_p5.Trail_t

	// energy      plot_p5.Trail_t
	// init_energy float64

	background_col = color.Black
)

func main() {
	// save_output := false

	a, period := 1.0, 1.0
	n_particles := 20
	steps_per_period := 250
	n_trails := int(float64(steps_per_period) * 2.0e-2)

	initialise_satellites(a, period, n_particles)
	initialise_integrator(period, steps_per_period)
	initialise_window(a)
	initialise_draw_objects(n_trails)

	p5.Run(setup, draw_frame)
}

func setup() {
	p5.PhysCanvas(
		dimensions.Pixels_x, dimensions.Pixels_y,
		dimensions.X_min, dimensions.X_max,
		dimensions.Y_min, dimensions.Y_max)
	p5.Background(background_col)
}

func initialise_satellites(a, period float64, n_particles int) {

	orbit := kepler.New_elliptical_orbit(a, 0.0, period)

	particles := make([]physics.Particle_t, n_particles+2)

	particles[0] = physics.New_particle("Sun", orbit.Mu/physics.G, vector.Null(), vector.Null())

	phi_jupiter := -math.Pi / 2
	jupiter_mass := particles[0].Mass * 1.0e-3
	particles[1] = kepler.New_satellite(phi_jupiter, jupiter_mass, &orbit)
	particles[1].Name = "Jupiter"

	average_mass := particles[1].Mass * 1.0e-5
	// dm := 1.0
	d_phi_max := math.Pi / 18.0
	d_r_max := 5.0e-2

	for i := 2; i < len(particles); i++ {

		phi := phi_jupiter + random.Signed_float64(d_phi_max)

		if i%2 == 0 {
			phi += math.Pi / 3.0
		} else {
			phi -= math.Pi / 3.0
		}
		// mass := average_mass * (1.0 + float64(2*rand.Intn(2)-1)*rand.Float64()*dm)
		mass := average_mass

		particles[i] = kepler.New_satellite(phi, mass, &orbit)
		particles[i].Position = r2.Scale(1.0+random.Signed_float64(d_r_max), particles[i].Position)

		distance := r2.Norm(particles[i].Position)
		speed := kepler.Speed_along_circle(distance, orbit.Mu)
		particles[i].Velocity = r2.Scale(speed/r2.Norm(particles[i].Velocity), particles[i].Velocity)
	}

	system = physics.System_t{Force: &physics.Gravity_interparticle_t{}, Particles: particles}

	// init_energy = system.Energy()
}

func initialise_draw_objects(n_trails int) {
	// jupiter orbital distance: 7.4e11 m
	// solar radius: 7.0e8 m
	// jupiter radius: 7.0e7 km
	// solar diameter = 2.0e-3 dJ
	// jupiter diameter = 2.0e-4 dJ

	solar_col := color.RGBA{R: 255, G: 255, B: 51, A: 255}
	jupiter_col := color.RGBA{R: 255, G: 95, B: 26, A: 255}
	particle_col := color.RGBA{R: 166, G: 166, B: 166, A: 255}

	sol_size := 6.0e-2
	jupiter_size := sol_size * 2.0e-1
	asteroid_size := jupiter_size * 0.5
	// sol_size := 2.0e-3
	// jupiter_size := sol_size * 1.0e-1
	// asteroid_size := jupiter_size * 0.5

	dots = plot_p5.Dots_from_system(system.Particles, particle_col, asteroid_size)

	dots[0].Set_col(solar_col)
	dots[0].Set_diameter(sol_size)

	dots[1].Set_col(jupiter_col)
	dots[1].Set_diameter(jupiter_size)

	trails = plot_p5.Trails_from_system(system.Particles, particle_col, n_trails)

	trails[0].Set_length(0)

	trails[1].Set_color(jupiter_col)

}

func initialise_window(a float64) {
	physical_width := 3.5 * a
	center_width_frac, center_height_frac := 0.5, 0.5
	dimensions = plot_p5.New_dimensions(1200, 1000, physical_width, center_width_frac, center_height_frac)
	fmt.Println("Window: ", dimensions)
}

func initialise_integrator(period float64, steps_per_period int) {
	stepper = integrator.New_stepper(integrator.Default_O6_algorithm())
	timestep = period / float64(steps_per_period)
}

// func initialise_energy(n_trails int) {
// 	energy = plot_p5.New_trail(
// 		color.RGBA{R: 223, G: 120, B: 036, A: 255},
// 		n_trails,
// 	)
// }

func output_variables() {
	fmt.Print("system: ", system, "\n")
	fmt.Print("stepper: ", stepper, "\n")
}

func draw_frame() {

	stepper.Run(&system, timestep)

	plot_p5.Update_dots(dots, system.Particles)
	plot_p5.Update_trails(trails, system.Particles)

	for i := range dots {
		dots[i].Plot()
		trails[i].Plot()
	}

	step_count += 1

	// fmt.Printf("time = %.2e, d_energy = %.2e\n", system.Time, dE(system.Energy(), init_energy))

	// filename := fmt.Sprintf("frame_%v.png", step_count)
	// p5.Screenshot(filename)

}

// func draw_energy_frame() {

// 	stepper.Run(&system, timestep)

// 	energy.Update(vector.New(system.Time, system.Energy()))

// 	// plot_p5.Update_dots(dots, system.Particles)
// 	// plot_p5.Update_trails(trails, system.Particles)
// 	// plot_p5.Update_velocities(velocities, system.Particles)

// 	// solar_pulse.Update_time()
// 	// solar_pulse.Plot()

// 	energy.Plot()

// 	// for i := range dots {
// 	// 	dots[i].Plot()
// 	// 	trails[i].Plot()
// 	// 	// velocities[i].Plot()
// 	// }

// 	step_count += 1

// 	// filename := fmt.Sprintf("frame_%v.png", step_count)
// 	// p5.Screenshot(filename)

// }

// func dE(E, E_0 float64) float64 {
// 	return 1 - E/E_0
// }
