package main

import (
	"fmt"
	"image/color"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/plot_p5"
	"github.com/granite/pkg/tree"
	"github.com/granite/pkg/vector"
)

var (
	root           tree.Node_t
	particles      []physics.Particle_t
	background_col = color.Black
	dots           []plot_p5.Dot_t
)

func main() {

	root = tree.New_node(tree.New_boundary(-1, 1, -1, 1))

	default_mass := 1.0
	particles = []physics.Particle_t{
		physics.New_particle("", default_mass, vector.New(0.52, 0.42), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.55, 0.39), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.6, 0.4), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.7, 0.8), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.7, 0.2), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.6, 0.3), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.85, 0.3), vector.Null()),
		physics.New_particle("", default_mass, vector.New(0.7, -0.2), vector.Null()),
		physics.New_particle("", default_mass, vector.New(-0.5, -0.5), vector.Null()),
	}

	for i := range particles {
		root.Insert(&particles[i])
	}

	fmt.Println(root)

	dot_size := 1.0e-2
	dot_color := color.RGBA{R: 246, G: 244, B: 129, A: 255}
	dots = plot_p5.Dots_from_system(particles, dot_color, dot_size)
	plot_p5.Update_dots(dots, particles)

	for i := range dots {
		fmt.Println(dots[i])
	}

	p5.Run(setup, draw_frame)
}

func setup() {

	dimensions := plot_p5.New_dimensions(1000, 1000, 2.1, 0.5, 0.5)
	fmt.Println(dimensions)

	p5.PhysCanvas(dimensions.Pixels_x, dimensions.Pixels_y,
		dimensions.X_min, dimensions.X_max, dimensions.Y_min, dimensions.Y_max)
	p5.Background(background_col)
}

func draw_frame() {
	plot_p5.Draw_tree(&root)
	for i := range dots {
		dots[i].Plot()
	}
}
