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

func New_arrow(col color.Color) Arrow_t {
	return Arrow_t{
		col: col,
	}
}

func (arrow *Arrow_t) Update(new_start, new_end vector.Vec) {
	arrow.start = new_start
	arrow.end = new_end
}

func (arrow *Arrow_t) Plot() {
	p5.Stroke(arrow.col)
	p5.Line(arrow.start.X, -arrow.start.Y, arrow.end.X, -arrow.end.Y)
}
