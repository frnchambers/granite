package plot_p5

import (
	"fmt"

	"github.com/granite/pkg/vector"
)

type Window_dimensions_t struct {
	Pixels_x, Pixels_y         int
	Trail_length               int
	Dot_size, Step_time        float64
	X_min, X_max, Y_min, Y_max float64
}

func (win Window_dimensions_t) String() string {
	return fmt.Sprintf("Pixels: (%d, %d), X limits: (%.2e, %.2e), Y limits: (%.2e, %.2e)",
		win.Pixels_x, win.Pixels_y, win.X_min, win.X_max, win.Y_min, win.Y_max)
}

type Plottable_t interface {
	Update(vector.Vec)
	Plot()
}

func New_dimensions(pixels_x, pixels_y int, physical_width, center_width_pc, center_height_pc float64) Window_dimensions_t {
	tan_theta := float64(pixels_y) / float64(pixels_x)
	return Window_dimensions_t{
		Pixels_x: pixels_x,
		Pixels_y: pixels_y,
		// X_min:    -physical_width * (1.0 + center_width_pc),
		// X_max:    physical_width * (1.0 - center_width_pc),
		// Y_min:    -physical_width * tan_theta * (1.0 + center_height_pc),
		// Y_max:    physical_width * tan_theta * (1.0 - center_height_pc),
		X_min: -physical_width * center_width_pc,
		X_max: physical_width * (1.0 - center_width_pc),
		Y_min: -physical_width * tan_theta * (1.0 - center_height_pc),
		Y_max: physical_width * tan_theta * center_height_pc,
	}
}
