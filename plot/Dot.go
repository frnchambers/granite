package plot

import (
	"image/color"

	"github.com/go-p5/p5"
	"gonum.org/v1/gonum/spatial/r2"
)

type Dot_t struct {
	col      color.Color
	diameter float64
	data     r2.Vec
}

func New_dot(lookup *r2.Vec, col color.Color, diameter float64) Dot_t {
	return Dot_t{
		col:      col,
		diameter: diameter,
	}
}

func New_static_dot(col color.Color, diameter float64, position r2.Vec) (dot Dot_t) {
	dot = Dot_t{
		col:      col,
		diameter: diameter,
	}
	dot.Update(position)
	return
}

func (dot *Dot_t) Update(position r2.Vec) {
	dot.data = position
}

func (dot *Dot_t) Plot() {
	p5.Stroke(dot.col)
	p5.Fill(dot.col)
	p5.Circle(dot.data.X, -dot.data.Y, dot.diameter)
}

func (dot *Dot_t) Position() r2.Vec {
	return dot.data
}
