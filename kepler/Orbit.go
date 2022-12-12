package kepler

// all calculations assume that the main body is centered at (0, 0)
// the perihelion and aphelion are on the x-axis

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/spatial/r2"
	"main.go/test_tools"
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

func Tangent_to_elliplse(phi float64, orbit *Orbit_t) r2.Vec {
	r := Distance(phi, orbit)
	chi := Chi(phi, orbit.Eccentricity)
	x, y := test_tools.Destructure_vector(Position_along_elliplse(phi, orbit))
	scale := 1.0 - orbit.Eccentricity*math.Cos(phi)/chi
	y_shift := orbit.Eccentricity * r / chi
	return r2.Unit(r2.Vec{X: -y * scale, Y: x*scale + y_shift})
}

func Position_along_elliplse(phi float64, orbit *Orbit_t) r2.Vec {
	r := Distance(phi, orbit)
	return Cartesian_position_from_polar(r, phi)
}

func Distance(phi float64, orbit *Orbit_t) float64 {
	return orbit.Semi_major * (1.0 - orbit.Eccentricity*orbit.Eccentricity) / Chi(phi, orbit.Eccentricity)
}

func Speed_from_distance(distance float64, orbit *Orbit_t) float64 {
	return math.Sqrt(orbit.Mu * (2.0/distance - 1.0/orbit.Semi_major))
}

func Speed(phi float64, orbit *Orbit_t) float64 {
	return Speed_from_distance(Distance(phi, orbit), orbit)
}

func Time_to_perihelion(phi float64, orbit *Orbit_t) float64 {
	return math.Sqrt(orbit.Semi_major*orbit.Semi_major*orbit.Semi_major/orbit.Mu) * Tau(phi, orbit.Eccentricity)
}

func Phi_for_time_to_perihelion(time float64, orbit *Orbit_t) (output float64, err error) {
	if math.Abs(time) > orbit.Period {
		message := fmt.Sprintf("Invalid time, %.2e, must be less than one orbital period, %.2e", time, orbit.Period)
		return output, errors.New(message)
	}
	bisec := test_tools.New_bisection_parameters(0.0, math.Pi)
	bisec.Y_desired = time
	output, err = test_tools.Bisection(Time_to_perihelion, orbit, &bisec)
	return
}

func Chi(phi, ecc float64) float64 {
	return 1.0 + ecc*math.Cos(phi)
}

func Tau(phi, ecc float64) (output float64) {
	// one_m_eccsq := 1.0 - ecc*ecc
	return 2.0*math.Atan(math.Sqrt((1.0-ecc)/(1.0+ecc))*math.Tan(0.5*phi)) -
		ecc*math.Sqrt(1-ecc*ecc)*math.Sin(phi)/Chi(phi, ecc)
}

func Cartesian_position_from_polar(r, phi float64) r2.Vec {
	return r2.Vec{X: x_from_polar(r, phi), Y: y_from_polar(r, phi)}
}

func x_from_polar(r, phi float64) float64 {
	return r * math.Cos(phi)
}

func y_from_polar(r, phi float64) float64 {
	return r * math.Sin(phi)
}
