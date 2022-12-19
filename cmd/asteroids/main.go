package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/go-p5/p5"
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

	step_count = 0

	solar_pulse plot_p5.Pulse_t
	dots        []plot_p5.Dot_t
	trails      []plot_p5.Trail_t
	velocities  []plot_p5.Arrow_t

	energy      plot_p5.Trail_t
	init_energy float64

	stepper integrator.Stepper_t
)

func initialise_satellites(
	a, period float64,
	n_steps, n_trails int,
	n_particles int,
) {
	solar_position := vector.Vec{X: 0, Y: 0}

	orbit = kepler.New_elliptical_orbit(a, 0.0, period)
	sim = kepler.New_simulation_parameters(n_steps, n_trails, &orbit)

	d_phi_max := math.Pi / 6.0
	d_r_max := 2.0e-2

	particles := make([]physics.Particle_t, n_particles+1)

	particles[0] = physics.New_particle("Sun", orbit.Mu/physics.G, vector.Null(), vector.Null())

	phi_jupiter := -math.Pi / 2
	jupiter_mass := particles[0].Mass * 1.0e-3
	particles[1] = kepler.New_satellite(phi_jupiter, jupiter_mass, &orbit)
	particles[1].Name = "Jupiter"

	average_mass := particles[1].Mass * 1.0e-5
	dm := 1.0

	for i := 2; i < n_particles+1; i++ {
		phi := phi_jupiter + float64(2*rand.Intn(2)-1)*rand.Float64()*d_phi_max
		if i%2 == 0 {
			phi += math.Pi / 3.0
		} else {
			phi -= math.Pi / 3.0
		}
		mass := average_mass * (1.0 + float64(2*rand.Intn(2)-1)*rand.Float64()*dm)
		particles[i] = kepler.New_satellite(phi, mass, &orbit)
		particles[i].Position = r2.Scale(1.0+float64(2*rand.Intn(2)-1)*rand.Float64()*d_r_max, particles[i].Position)
	}

	system = physics.System_t{Force: &physics.Gravity_interparticle_t{}, Particles: particles}

	init_energy = system.Energy()

	solar_pulse = plot_p5.New_pulse(color.RGBA{R: 246, G: 244, B: 129, A: 255}, n_steps, 2.0e-2, 1.0e-1)
	solar_pulse.Update_position(solar_position)
	solar_pulse.Reset_time(0.0)

	dots = plot_p5.Dots_from_system(system.Particles, sim.Dot_size)

	dots[0].Set_col(color.RGBA{R: 246, G: 244, B: 129, A: 255})
	dots[0].Set_diameter(1.0e-1)

	// dots[1].Set_col(color.RGBA{R: 246, G: 244, B: 129, A: 255})
	dots[1].Set_diameter(5.0e-2)

	trails = plot_p5.Trails_from_system(system.Particles, sim.Trail_length)
	velocities = plot_p5.Velocities_from_system(system.Particles)
}

func initialise_energy(n_trails int) {
	energy = plot_p5.New_trail(
		color.RGBA{R: 223, G: 120, B: 036, A: 255},
		n_trails,
	)
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

func highly_eccentric_settings() {
	a, period := 1.0, 1.0
	n_steps, n_trails := 200, 3
	n_particles := 32

	initialise_satellites(a, period, n_steps, n_trails, n_particles)

	stepper = integrator.New_stepper(integrator.Default_O6_algorithm())
	// stepper = integrator.New_stepper(integrator.Version_3_5_1_v_3())
	// stepper = integrator.New_stepper(integrator.Version_3_5_1_v_2())
}

func run_simulation() {
	highly_eccentric_settings()
	output_variables()
	p5.Run(setup, draw_positions_frame)
}

func setup() {
	p5.PhysCanvas(edge_pixel_count, edge_pixel_count,
		sim.X_min, sim.X_max, sim.Y_min, sim.Y_max)
	p5.Background(color.Black)

	// p5.Stroke(color.White)
	// p5.Fill(color.Transparent)
	// p5.Ellipse(0, 0, orbit.Semi_major*2, orbit.Semi_minor*2)

}

func draw_positions_frame() {

	stepper.Run(&system, sim.Step_time)

	plot_p5.Update_dots(dots, system.Particles)
	plot_p5.Update_trails(trails, system.Particles)
	// plot_p5.Update_velocities(velocities, system.Particles)

	// solar_pulse.Update_time()
	// solar_pulse.Plot()

	for i := range dots {
		dots[i].Plot()
		trails[i].Plot()
		// velocities[i].Plot()
	}

	step_count += 1

	// fmt.Printf("time = %.2e, d_energy = %.2e\n", system.Time, dE(system.Energy(), init_energy))

	// filename := fmt.Sprintf("frame_%v.png", step_count)
	// p5.Screenshot(filename)

}

func draw_energy_frame() {

	stepper.Run(&system, sim.Step_time)

	energy.Update(vector.New(system.Time, system.Energy()))

	// plot_p5.Update_dots(dots, system.Particles)
	// plot_p5.Update_trails(trails, system.Particles)
	// plot_p5.Update_velocities(velocities, system.Particles)

	// solar_pulse.Update_time()
	// solar_pulse.Plot()

	energy.Plot()

	// for i := range dots {
	// 	dots[i].Plot()
	// 	trails[i].Plot()
	// 	// velocities[i].Plot()
	// }

	step_count += 1

	// filename := fmt.Sprintf("frame_%v.png", step_count)
	// p5.Screenshot(filename)

}

func dE(E, E_0 float64) float64 {
	return 1 - E/E_0
}
