package plot

import (
	"image/color"

	"github.com/go-p5/p5"
	"gonum.org/v1/gonum/spatial/r2"
)

type Arrow_t struct {
	col        color.Color
	start, end r2.Vec
}

func New_arrow(lookup *r2.Vec, col color.Color) Arrow_t {
	return Arrow_t{
		col: col,
	}
}

func (arrow *Arrow_t) Update(start, end r2.Vec) {
	arrow.start = start
	arrow.end = end
}

func (arrow *Arrow_t) Plot() {
	p5.Stroke(arrow.col)
	p5.Line(arrow.start.X, -arrow.start.Y, arrow.end.X, -arrow.end.Y)
}
