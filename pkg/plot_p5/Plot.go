package plot_p5

import "github.com/granite/pkg/vector"

type Window_dimensions_t struct {
	Pixels_x, Pixels_y         int
	Trail_length               int
	Dot_size, Step_time        float64
	X_min, X_max, Y_min, Y_max float64
}

type Plottable_t interface {
	Update(vector.Vec)
	Plot()
}
