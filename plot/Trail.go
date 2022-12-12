package plot

import (
	"image/color"

	"github.com/go-p5/p5"
	"gonum.org/v1/gonum/spatial/r2"
)

type Trail_t struct {
	col    color.Color
	length int
	data   []r2.Vec
}

func New_trail(lookup *r2.Vec, col color.Color, length int) Trail_t {
	return Trail_t{
		col:    col,
		length: length,
		data:   make([]r2.Vec, 0, length),
	}
}

func (trail *Trail_t) Update(position r2.Vec) {
	if len(trail.data) >= trail.length {
		trail.data = append([]r2.Vec{position}, trail.data[:len(trail.data)-1]...)
	} else {
		trail.data = append([]r2.Vec{position}, trail.data...)
	}
}

func (trail *Trail_t) Plot() {
	for i := 0; i < len(trail.data)-1; i++ {
		pos_i := trail.data[i]
		pos_j := trail.data[i+1]
		p5.Stroke(trail.col)
		p5.Line(pos_i.X, -pos_i.Y, pos_j.X, -pos_j.Y)
	}
}
