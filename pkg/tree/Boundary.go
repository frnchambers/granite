package tree

import (
	"errors"
	"fmt"

	"github.com/granite/pkg/vector"
)

type Quadrant int8

const (
	//  2|3
	// --+--
	//  0|1
	Lower_left Quadrant = iota
	Lower_right
	Upper_left
	Upper_right
)

const (
	N_quadrants = 4
)

type Boundary_t struct {
	Min_x, Mid_x, Max_x,
	Min_y, Mid_y, Max_y float64
}

func New_boundary(min_x, max_x, min_y, max_y float64) (boundary Boundary_t) {
	boundary = Boundary_t{
		Min_x: min_x, Max_x: max_x,
		Min_y: min_y, Max_y: max_y,
	}
	boundary.Mid_x = 0.5 * (min_x + max_x)
	boundary.Mid_y = 0.5 * (min_y + max_y)
	return
}

func Null_boundary() Boundary_t {
	return New_boundary(0.0, 0.0, 0.0, 0.0)
}

func (boundary *Boundary_t) Sub_quadrant(q Quadrant) (sub Boundary_t) {
	switch q {
	case Lower_left:
		return New_boundary(boundary.Min_x, boundary.Mid_x, boundary.Min_y, boundary.Mid_y)
	case Lower_right:
		return New_boundary(boundary.Mid_x, boundary.Max_x, boundary.Min_y, boundary.Mid_y)
	case Upper_left:
		return New_boundary(boundary.Min_x, boundary.Mid_x, boundary.Mid_y, boundary.Max_y)
	case Upper_right:
		return New_boundary(boundary.Mid_x, boundary.Max_x, boundary.Mid_y, boundary.Max_y)
	}
	return Null_boundary()
}

func (boundary *Boundary_t) Contains(position vector.Vec) bool {
	if position.X < boundary.Min_x || position.X >= boundary.Max_x ||
		position.Y < boundary.Min_y || position.Y >= boundary.Max_y {
		return false
	}
	return true
}

func (boundary *Boundary_t) Choose_quadrant(position vector.Vec) (q Quadrant) {
	q = 0
	if position.X >= boundary.Mid_x {
		q += 1
	}
	if position.Y >= boundary.Mid_y {
		q += 2
	}
	return
}

func (boundary *Boundary_t) Choose_quadrant_safe(position vector.Vec) (q Quadrant, err error) {

	err = nil
	if !boundary.Contains(position) {
		message := fmt.Sprintf("Given position, {%.2e, %.2e}, is outside the limits of the boundary range: X: [%.2e, %.2e), Y: [%.2e, %.2e)",
			position.X, position.Y, boundary.Min_x, boundary.Max_x, boundary.Min_y, boundary.Max_y)
		err = errors.New(message)
		return
	}

	q = boundary.Choose_quadrant(position)

	return
}
