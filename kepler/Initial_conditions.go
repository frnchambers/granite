package kepler

import (
	"github.com/granite/physics"
	"github.com/granite/plot"
	"github.com/granite/test_tools"
	"gonum.org/v1/gonum/spatial/r2"
)

func New_satellite(phi float64, orbit *Orbit_t) physics.Particle_t {

	position := Position_along_elliplse(phi, orbit)
	velocoity := r2.Scale(Speed(phi, orbit), Tangent_to_elliplse(phi, orbit))

	return physics.New_particle("", 1.0, position, velocoity)
}

func New_satellite_from_perihelion(orbit *Orbit_t) physics.Particle_t {
	position := Cartesian_position_from_polar(orbit.Perihelion, 0.0)
	velocoity := r2.Vec{X: 0.0, Y: orbit.V_perihelion}
	return physics.New_particle("", 1.0, position, velocoity)
}

func New_massive_body(orbit *Orbit_t) physics.Gravity_massive_body_t {
	return physics.New_massive_body(
		physics.New_particle("", orbit.Mu/physics.G, test_tools.Null_vector(), test_tools.Null_vector()))
}

func New_simulation_parameters(n_steps, n_trails int, orbit *Orbit_t) plot.Simulation_t {
	// n_trails := 20
	// n_steps := 500
	dt := orbit.Period / float64(n_steps)

	return plot.Simulation_t{
		Trail_length: n_trails,
		Dot_size:     2.0e-2,
		X_min:        -2*orbit.Semi_major - orbit.Linear_eccentricity,
		X_max:        2*orbit.Semi_major - orbit.Linear_eccentricity,
		Y_min:        -orbit.Semi_major * 2.0,
		Y_max:        orbit.Semi_major * 2.0,
		Step_time:    dt,
	}
}

func Rotate_orbit(particle *physics.Particle_t, phi float64) {
	centre := r2.Vec{X: 0, Y: 0}
	particle.Position = r2.Rotate(particle.Position, phi, centre)
	particle.Velocity = r2.Rotate(particle.Velocity, phi, centre)
}

// func Satellite_in_orbit(a, ecc, period float64, d_phi float64) (physics.System_t, r2.Vec, plot.Simulation_t) {
// 	return Satellites_in_orbit(a, ecc, period, 1, d_phi, 0.0)
// }

// func Satellites_in_orbit(a, ecc, period float64, n_particles int, init_phi, d_phi float64) (
// 	system physics.System_t,
// 	body_position r2.Vec,
// 	sim_params plot.Simulation_t,
// ) {

// 	body_position = r2.Vec{X: 0, Y: 0}

// 	lin_ecc := ecc * a
// 	mu := math.Pow(2.0*math.Pi/period, 2.0) * math.Pow(a, 3.0)
// 	body_mass := mu / physics.G
// 	r_perihelion := a * (1 - ecc)
// 	v_perihelion := math.Sqrt(mu) * math.Sqrt(2.0/r_perihelion-1.0/a)

// 	particles := make([]physics.Particle_t, 0, n_particles)

// 	n_trails := 20
// 	n_steps := 500
// 	dt := period / float64(n_steps)

// 	sim_params = plot.Simulation_t{
// 		Trail_length: n_trails,
// 		Dot_size:     2.0e-2,
// 		X_min:        2*-a - lin_ecc,
// 		X_max:        2*a - lin_ecc,
// 		Y_min:        -a * 2.0,
// 		Y_max:        +a * 2.0,
// 		Step_time:    dt,
// 	}

// 	for i := 0; i < n_particles; i++ {
// 		position := r2.Rotate(r2.Add(body_position, r2.Vec{X: r_perihelion, Y: 0}), init_phi+float64(i)*d_phi, body_position)
// 		velocity := r2.Rotate(r2.Vec{X: 0, Y: v_perihelion}, init_phi+float64(i)*d_phi, body_position)

// 		particles = append(particles, physics.New_particle(
// 			"Satelite", 1.0,
// 			position,
// 			velocity,
// 		))
// 	}

// 	central_body := physics.New_particle("Sun", body_mass, body_position, body_position)

// 	system = physics.System_t{Force: physics.New_massive_body(central_body), Particles: particles}

// 	return
// }
