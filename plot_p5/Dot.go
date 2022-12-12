package plot_p5

import (
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/vector"
)

type Dot_t struct {
	col      color.Color
	diameter float64
	data     vector.Vec
}

func New_dot(lookup *vector.Vec, col color.Color, diameter float64) Dot_t {
	return Dot_t{
		col:      col,
		diameter: diameter,
	}
}

func New_static_dot(col color.Color, diameter float64, position vector.Vec) (dot Dot_t) {
	dot = Dot_t{
		col:      col,
		diameter: diameter,
	}
	dot.Update(position)
	return
}

func (dot *Dot_t) Update(position vector.Vec) {
	dot.data = position
}

func (dot *Dot_t) Plot() {
	p5.Stroke(dot.col)
	p5.Fill(dot.col)
	p5.Circle(dot.data.X, -dot.data.Y, dot.diameter)
}

func (dot *Dot_t) Position() vector.Vec {
	return dot.data
}
