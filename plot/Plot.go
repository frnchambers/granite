package plot

import "gonum.org/v1/gonum/spatial/r2"

type Simulation_t struct {
	Trail_length               int
	Dot_size, Step_time        float64
	X_min, X_max, Y_min, Y_max float64
}

type Plottable_t interface {
	Update(r2.Vec)
	Plot()
}
