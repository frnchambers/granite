package plot_p5

import (
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/tree"
)

func Draw_tree(node *tree.Node_t) {

	Draw_boundary(&node.Boundary)

	for i := range node.Quadrants {
		if node.Quadrants[i] != nil {
			Draw_tree(node.Quadrants[i])
		}
	}

}

func Draw_boundary(boundary *tree.Boundary_t) {
	p5.Stroke(color.White)

	p5.Line(
		boundary.Min_x, -boundary.Min_y,
		boundary.Max_x, -boundary.Min_y,
	)
	p5.Line(
		boundary.Max_x, -boundary.Min_y,
		boundary.Max_x, -boundary.Max_y,
	)
	p5.Line(
		boundary.Max_x, -boundary.Max_y,
		boundary.Min_x, -boundary.Max_y,
	)
	p5.Line(
		boundary.Min_x, -boundary.Max_y,
		boundary.Min_x, -boundary.Min_y,
	)
}
