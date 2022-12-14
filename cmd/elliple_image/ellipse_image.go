package main

import (
	"fmt"
	"math"

	"github.com/granite/pkg/kepler"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	orbit := kepler.New_orbit(1.0, 0.3, 1.0)
	n := 81

	lineData := elliplse_points(&orbit, n)
	centreData := centre_point()

	p := plot.New()
	p.Title.Text = fmt.Sprintf("Orbital path (a = %.2e, e = %.2e)", orbit.Semi_major, orbit.Eccentricity)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	max_limit := orbit.Aphelion * (1.0 + 5.0e-2)
	p.X.Min, p.X.Max = -max_limit, max_limit
	p.Y.Min, p.Y.Max = -max_limit, max_limit

	p.Add(plotter.NewGrid())
	s, _ := plotter.NewScatter(centreData)
	l, _ := plotter.NewLine(lineData)
	p.Add(s, l)

	real_size := 20 * vg.Centimeter
	if err := p.Save(real_size, real_size, "points.png"); err != nil {
		panic(err)
	}
}

func elliplse_points(orbit *kepler.Orbit_t, n int) (points plotter.XYs) {
	points = make(plotter.XYs, n)

	d_phi := 2.0 * math.Pi / float64(n-1)
	phi := -math.Pi
	for i := range points {
		points[i].X, points[i].Y = vector.Destructure(
			kepler.Position_along_elliplse(phi, orbit))
		phi += d_phi
	}

	return
}

func centre_point() (points plotter.XYs) {
	points = make(plotter.XYs, 1)
	points[0].X, points[0].Y = 0.0, 0.0
	return
}
