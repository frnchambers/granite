package plot_p5

import (
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

type Flare_t struct {
	col                          color.Color
	min_size, delta_size, width  float64
	position, reference_position vector.Vec
}

func New_flare(col color.Color, reference_position vector.Vec,
	min_size, max_size, width float64) Flare_t {
	return Flare_t{
		col:                col,
		min_size:           min_size,
		width:              width,
		delta_size:         (max_size - min_size) * math.Sqrt(width),
		reference_position: reference_position,
	}
}

func New_flare_centre_zero(col color.Color,
	min_size, max_size, width float64) Flare_t {
	return New_flare(col, vector.Null(), min_size, max_size, width)
}

func (flare *Flare_t) Reset_time(new_reference_position vector.Vec) {
	flare.reference_position = new_reference_position
}

func (flare *Flare_t) Update(new_position vector.Vec) {
	flare.position = new_position
}

func (flare *Flare_t) Plot() {
	p5.Stroke(flare.col)
	p5.Fill(flare.col)
	p5.Circle(flare.position.X, -flare.position.Y, flare.diameter())
}

func (flare *Flare_t) Position() vector.Vec {
	return flare.position
}

func (flare *Flare_t) diameter() float64 {
	distance := r2.Norm(r2.Sub(flare.position, flare.reference_position))
	return flare.min_size + flare.delta_size/(math.Pow(distance, 2)+math.Pow(flare.width, 2))
}
