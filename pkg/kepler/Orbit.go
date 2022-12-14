package kepler

// all calculations assume that the main body is centered at (0, 0)
// the perihelion and aphelion are on the x-axis

import (
	"errors"
	"fmt"
	"math"

	"github.com/granite/pkg/root_solver"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

type Orbit_t struct {
	Semi_major, Semi_minor, Eccentricity, Linear_eccentricity,
	Mu, Period, Energy_per_mass,
	Aphelion, Perihelion,
	V_aphelion, V_perihelion float64
}

func New_orbit(a, ecc, period float64) (orbit Orbit_t) {
	orbit.Semi_major = a
	orbit.Eccentricity = ecc
	orbit.Period = period

	orbit.Semi_minor = a * math.Sqrt(1-ecc*ecc)
	orbit.Linear_eccentricity = a * ecc
	orbit.Perihelion = a * (1 - ecc)
	orbit.Aphelion = a * (1 + ecc)
	orbit.Mu = math.Pow(2.0*math.Pi/period, 2.0) * math.Pow(a, 3.0)
	orbit.Energy_per_mass = orbit.Mu / (2.0 * a)

	orbit.V_perihelion = Speed_from_distance(orbit.Perihelion, &orbit)
	orbit.V_aphelion = Speed_from_distance(orbit.Aphelion, &orbit)

	return
}

func Tangent_along_ellipse(phi float64, orbit *Orbit_t) vector.Vec {
	r := Distance_from_centre(phi, orbit)
	chi := chi(phi, orbit.Eccentricity)
	x, y := vector.Destructure(Position_along_elliplse(phi, orbit))
	scale := 1.0 - orbit.Eccentricity*math.Cos(phi)/chi
	y_shift := orbit.Eccentricity * r / chi
	return r2.Unit(vector.Vec{X: -y * scale, Y: x*scale + y_shift})
}

func Position_along_elliplse(phi float64, orbit *Orbit_t) vector.Vec {
	r := Distance_from_centre(phi, orbit)
	return vector.Cartesian_position_from_polar(r, phi)
}

func Distance_from_centre(phi float64, orbit *Orbit_t) float64 {
	return orbit.Semi_major * (1.0 - orbit.Eccentricity*orbit.Eccentricity) / chi(phi, orbit.Eccentricity)
}

func Speed_from_distance(distance float64, orbit *Orbit_t) float64 {
	return math.Sqrt(orbit.Mu * (2.0/distance - 1.0/orbit.Semi_major))
}

func Speed_along_ellipse(phi float64, orbit *Orbit_t) float64 {
	return Speed_from_distance(Distance_from_centre(phi, orbit), orbit)
}

func Speed_along_circle(a, mu float64) float64 {
	return math.Sqrt(mu / a)
}

func Time_to_perihelion(phi float64, orbit *Orbit_t) float64 {
	return math.Sqrt(orbit.Semi_major*orbit.Semi_major*orbit.Semi_major/orbit.Mu) * tau(phi, orbit.Eccentricity)
}

func Phi_for_time_to_perihelion(time float64, orbit *Orbit_t) (output float64, err error) {
	if math.Abs(time) > orbit.Period {
		message := fmt.Sprintf("Invalid time, %.2e, must be less than one orbital period, %.2e", time, orbit.Period)
		return output, errors.New(message)
	}
	bisec := root_solver.New_bisection_parameters(0.0, math.Pi)
	bisec.Y_desired = time
	output, err = root_solver.Bisection(Time_to_perihelion, orbit, &bisec)
	return
}

func chi(phi, ecc float64) float64 {
	return 1.0 + ecc*math.Cos(phi)
}

func tau(phi, ecc float64) (output float64) {
	// one_m_eccsq := 1.0 - ecc*ecc
	return 2.0*math.Atan(math.Sqrt((1.0-ecc)/(1.0+ecc))*math.Tan(0.5*phi)) -
		ecc*math.Sqrt(1-ecc*ecc)*math.Sin(phi)/chi(phi, ecc)
}

func (orbit Orbit_t) String() (output string) {
	output = fmt.Sprintf("Orbit_t: {")
	output += fmt.Sprintf("Semi-major axis = %.2e\n", orbit.Semi_major)
	output += fmt.Sprintf("Semi-minor axis = %.2e\n", orbit.Semi_minor)
	output += fmt.Sprintf("Eccentricity = %.2e\n", orbit.Eccentricity)
	output += fmt.Sprintf("Linear-eccentricity = %.2e\n", orbit.Linear_eccentricity)
	output += fmt.Sprintf("Mu axis             = %.2e\n", orbit.Mu)
	output += fmt.Sprintf("Period          = %.2e\n", orbit.Period)
	output += fmt.Sprintf("Energy per mass = %.2e\n", orbit.Energy_per_mass)
	output += fmt.Sprintf("R_aphelion   = %.2e\n", orbit.Aphelion)
	output += fmt.Sprintf("R_perihelion = %.2e\n", orbit.Perihelion)
	output += fmt.Sprintf("V_aphelion   = %.2e\n", orbit.V_aphelion)
	output += fmt.Sprintf("V_perihelion = %.2e\n", orbit.V_perihelion)
	output += fmt.Sprintf("}")
	return
}
