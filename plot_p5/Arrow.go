package plot_p5

import (
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/vector"
)

type Arrow_t struct {
	col        color.Color
	start, end vector.Vec
}

func New_arrow(lookup *vector.Vec, col color.Color) Arrow_t {
	return Arrow_t{
		col: col,
	}
}

func (arrow *Arrow_t) Update(start, end vector.Vec) {
	arrow.start = start
	arrow.end = end
}

func (arrow *Arrow_t) Plot() {
	p5.Stroke(arrow.col)
	p5.Line(arrow.start.X, -arrow.start.Y, arrow.end.X, -arrow.end.Y)
}
