package plot_p5

import (
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/vector"
)

type Trail_t struct {
	col    color.Color
	length int
	data   []vector.Vec
}

func New_trail(col color.Color, length int) Trail_t {
	return Trail_t{
		col:    col,
		length: length,
		data:   make([]vector.Vec, 0, length),
	}
}

func (trail *Trail_t) Update(next_position vector.Vec) {
	if len(trail.data) >= trail.length {
		trail.data = append([]vector.Vec{next_position}, trail.data[:len(trail.data)-1]...)
	} else {
		trail.data = append([]vector.Vec{next_position}, trail.data...)
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
