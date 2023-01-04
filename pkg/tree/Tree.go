package tree

import (
	"errors"
	"fmt"

	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/vector"
	"gonum.org/v1/gonum/spatial/r2"
)

type Node_t struct {
	Centre_of_mass vector.Vec
	Mass           float64
	Boundary       Boundary_t
	Particle       *physics.Particle_t
	Quadrants      []*Node_t
}

func (node *Node_t) String() string {

	if node.Is_leaf() {
		return fmt.Sprintf("mass: %.2e, centre of mass: (%.2e, %.2e), Partcle: %v",
			node.Mass, node.Centre_of_mass.X, node.Centre_of_mass.Y, &node.Particle)
	}

	return fmt.Sprintf("mass: %.2e, centre of mass: (%.2e, %.2e), N children: %d",
		node.Mass, node.Centre_of_mass.X, node.Centre_of_mass.Y, node.N_children())

}

func (node *Node_t) Has_at_least_one_child() bool {
	for _, p := range node.Quadrants {
		if p != nil {
			return true
		}
	}
	return false
}

func (node *Node_t) N_children() (count int) {
	count = 0
	for _, p := range node.Quadrants {
		if p != nil {
			count += 1
		}
	}
	return
}

func (node *Node_t) Is_leaf() bool {
	if node.Particle != nil {
		return true
	}
	return false
}

func (node *Node_t) Is_valid() bool {
	if node.Has_at_least_one_child() && node.Particle != nil {
		return false
	}
	return true
}

func New_node(boundary Boundary_t) (node Node_t) {
	node = Node_t{
		Boundary: boundary,
	}
	node.Centre_of_mass = vector.New(node.Boundary.Mid_x, node.Boundary.Mid_y)
	node.Quadrants = make([]*Node_t, N_quadrants)
	return
}

func (node *Node_t) Choose_quadrant_safe(position vector.Vec) (Quadrant, error) {
	return node.Boundary.Choose_quadrant_safe(position)
}

func (node *Node_t) Choose_quadrant_(position vector.Vec) Quadrant {
	return node.Boundary.Choose_quadrant(position)
}

func (node *Node_t) Insert(particle *physics.Particle_t) error {

	if !node.Boundary.Contains(particle.Position) {
		return errors.New("Insert: particle not contained in node")
	}

	node.Centre_of_mass = r2.Add(
		r2.Scale(node.Mass/(node.Mass+particle.Mass), node.Centre_of_mass),
		r2.Scale(particle.Mass/(node.Mass+particle.Mass), particle.Position))
	node.Mass += particle.Mass

	if node.Has_at_least_one_child() {
		q := node.Boundary.Choose_quadrant(particle.Position)

		if node.Quadrants[q] != nil {
			node.Quadrants[q].Insert(particle)
			return nil
		}

		n := New_node(node.Boundary.Sub_quadrant(q))
		node.Quadrants[q] = &n
		node.Quadrants[q].Insert(particle)

		return nil
	}

	if node.Particle == nil {
		node.Particle = particle
		return nil
	}

	q_new := node.Boundary.Choose_quadrant(particle.Position)
	q_existing := node.Boundary.Choose_quadrant(node.Particle.Position)

	if q_new == q_existing {
		n := New_node(node.Boundary.Sub_quadrant(q_existing))
		node.Quadrants[q_existing] = &n
		node.Quadrants[q_existing].Insert(node.Particle)
		node.Quadrants[q_existing].Insert(particle)
		node.Particle = nil
		return nil
	}

	new_node := New_node(node.Boundary.Sub_quadrant(q_new))
	node.Quadrants[q_new] = &new_node
	node.Quadrants[q_new].Insert(particle)

	existing_node := New_node(node.Boundary.Sub_quadrant(q_existing))
	node.Quadrants[q_existing] = &existing_node
	node.Quadrants[q_existing].Insert(node.Particle)

	node.Particle = nil
	return nil
}
