package tree

import (
	"fmt"
	"testing"

	"github.com/granite/pkg/comparison"
	"github.com/granite/pkg/physics"
	"github.com/granite/pkg/vector"
)

func Test_New_boundary(t *testing.T) {

	boundary := New_boundary(-1, 1, -1, 1)

	if !comparison.Float64_equality(boundary.Mid_x, 0.0) ||
		!comparison.Float64_equality(boundary.Mid_y, 0.0) {
		t.Fatalf(
			"Test failed : expect (mid_x, mid_y) = (0.0, 0.0), actual = (%.2e, %.2e)",
			boundary.Mid_x, boundary.Mid_y,
		)
	}

}

func Test_Contains(t *testing.T) {

	node := New_boundary(-1, 1, -1, 1)

	tests := []struct {
		input  vector.Vec
		expect bool
	}{
		{input: vector.New(0.5, 0.5), expect: true},
		{input: vector.New(-0.5, 1.5), expect: false},
		{input: vector.New(1.5, -1.5), expect: false},
		{input: vector.New(-0.5, -0.5), expect: true},
	}

	for i, test := range tests {

		// when
		actual := node.Contains(test.input)

		if actual != test.expect {
			t.Fatalf(
				"Test, i = %v, %v, failed : expect = %v, actual = %v",
				i, test, test.expect, actual,
			)
		}
	}

}

func Test_Choose_quadrant(t *testing.T) {

	boundary := New_boundary(-1, 1, -1, 1)

	tests := []struct {
		input  vector.Vec
		expect Quadrant
	}{
		{input: vector.New(0.5, 0.5), expect: Upper_right},
		{input: vector.New(-0.5, 0.5), expect: Upper_left},
		{input: vector.New(0.5, -0.5), expect: Lower_right},
		{input: vector.New(-0.5, -0.5), expect: Lower_left},
	}

	for i, test := range tests {

		// when
		actual, err := boundary.Choose_quadrant_safe(test.input)

		if err != nil {
			t.Fatalf(
				"Test, i = %v, %v, failed : err = %v",
				i, test, err,
			)
		}

		if actual != test.expect {
			t.Fatalf(
				"Test, i = %v, %v, failed : expect = %v, actual = %v",
				i, test, test.expect, actual,
			)
		}
	}

}

func Test_New_node(t *testing.T) {

	node := New_node(New_boundary(-1, 1, -1, 1))

	if !comparison.Float64_equality(node.Mass, 0.0) {
		t.Fatalf(
			"Test failed : expect mass = 0.0, actual = %.2e",
			node.Mass,
		)
	}

	if !vector.Are_equal(node.Centre_of_mass, vector.Null()) {
		t.Fatalf(
			"Test failed : expect centre_of_mass = (0.0, 0.0), actual = %v",
			node.Centre_of_mass,
		)
	}

	if node.Particle != nil {
		t.Fatalf(
			"Test failed : expect nil for particle but found: %v",
			node.Particle,
		)
	}

	for _, p := range node.Quadrants {
		if p != nil {
			t.Fatalf(
				"Test failed : expect nil for all quadrants but found: %v",
				node.Quadrants,
			)
		}
	}

}

func Test_Fill_tree(t *testing.T) {

	node := New_node(New_boundary(-1, 1, -1, 1))

	tests := []struct {
		input  vector.Vec
		expect Quadrant
	}{
		{input: vector.New(0.5, 0.5), expect: Upper_right},
		{input: vector.New(-0.5, 0.5), expect: Upper_left},
		{input: vector.New(0.5, -0.5), expect: Lower_right},
		{input: vector.New(-0.5, -0.5), expect: Lower_left},
	}

	for i, test := range tests {

		// when
		actual, err := node.Choose_quadrant_safe(test.input)

		if err != nil {
			t.Fatalf(
				"Test, i = %v, %v, failed : err = %v",
				i, test, err,
			)
		}

		if actual != test.expect {
			t.Fatalf(
				"Test, i = %v, %v, failed : expect = %v, actual = %v",
				i, test, test.expect, actual,
			)
		}
	}

}

func Test_Insert_single_particle(t *testing.T) {

	node := New_node(New_boundary(-1, 1, -1, 1))
	particle := physics.New_particle("", 1.0, vector.New(0.5, 0.5), vector.Null())

	// when
	err := node.Insert(&particle)

	if err != nil {
		t.Fatalf("Unexpected failure with error: %v", err)
	}

	if !comparison.Float64_equality(particle.Mass, node.Mass) {
		t.Fatalf(
			"Masses not equal : expect = %.2e, actual = %.2e",
			particle.Mass, node.Mass,
		)
	}

	if !vector.Are_equal(particle.Position, node.Centre_of_mass) {
		t.Fatalf(
			"Positions not equal : expect = %v, actual = %v",
			particle.Position, node.Centre_of_mass,
		)
	}

	if &particle != node.Particle {
		t.Fatalf(
			"Node does not point to particle: expect = %v, actual = %v",
			&particle, node.Particle,
		)
	}

}

func Test_Insert_multiple_particles(t *testing.T) {

	node := New_node(New_boundary(-1, 1, -1, 1))

	particle1 := physics.New_particle("", 1.0, vector.New(0.4, 0.25), vector.Null())
	particle2 := physics.New_particle("", 1.0, vector.New(-0.5, 0.25), vector.Null())
	particle3 := physics.New_particle("", 1.0, vector.New(0.1, -0.5), vector.Null())

	// when
	err := node.Insert(&particle1)
	err = node.Insert(&particle2)
	err = node.Insert(&particle3)

	if err != nil {
		t.Fatalf("Unexpected failure with error: %v", err)
	}

	message := node.assert_equals(3.0, vector.Null(), nil, []Quadrant{Lower_right, Upper_left, Upper_right})

	if message != "" {
		t.Fatalf(message)
	}

}

func Test_Insert_depth_two(t *testing.T) {

	node_level_0 := New_node(New_boundary(-1, 1, -1, 1))

	particle1 := physics.New_particle("", 1.0, vector.New(0.25, 0.25), vector.Null())
	particle2 := physics.New_particle("", 1.0, vector.New(0.75, 0.75), vector.Null())

	// when
	err := node_level_0.Insert(&particle1)
	err = node_level_0.Insert(&particle2)

	if err != nil {
		t.Fatalf("Unexpected failure with error: %v", err)
	}

	message := node_level_0.assert_equals(2.0, vector.New(0.5, 0.5), nil, []Quadrant{Upper_right})
	if message != "" {
		t.Fatalf(message)
	}

	node_level_1_upper_right := node_level_0.Quadrants[Upper_right]
	message = node_level_1_upper_right.assert_equals(2.0, vector.New(0.5, 0.5), nil, []Quadrant{Upper_right, Lower_left})
	if message != "" {
		t.Fatalf(message)
	}

	node_level_2_lower_left := node_level_1_upper_right.Quadrants[Lower_left]
	message = node_level_2_lower_left.assert_equals(1.0, vector.New(0.25, 0.25), &particle1, []Quadrant{})
	if message != "" {
		t.Fatalf(message)
	}

	node_level_2_upper_right := node_level_1_upper_right.Quadrants[Upper_right]
	message = node_level_2_upper_right.assert_equals(1.0, vector.New(0.75, 0.75), &particle2, []Quadrant{})
	if message != "" {
		t.Fatalf(message)
	}

}

func (node *Node_t) assert_equals(
	expected_mass float64,
	expected_com vector.Vec,
	expected_particle_address *physics.Particle_t,
	expected_occupied_quadrants []Quadrant,
) string {

	if !comparison.Float64_equality(expected_mass, node.Mass) {
		return fmt.Sprintf("Masses not equal : expect = %.2e, actual = %.2e",
			expected_mass, node.Mass,
		)
	}

	if !vector.Are_equal(expected_com, node.Centre_of_mass) {
		return fmt.Sprintf(
			"Positions not equal : expect = %v, actual = %v",
			expected_com, node.Centre_of_mass,
		)
	}

	if expected_particle_address != node.Particle {
		return fmt.Sprintf(
			"Node does not point to nil: expect = %v, actual = %v",
			&expected_particle_address, node.Particle,
		)
	}

	for _, q := range expected_occupied_quadrants {
		if node.Quadrants[q] == nil {
			return fmt.Sprintf(
				"Node does not point to nil in q = %d, quadrants = %v",
				q, node.Quadrants,
			)
		}
	}

	expected_unoccupied_quadrants := make_unoccupied_quadrants(expected_occupied_quadrants)
	for _, q := range expected_unoccupied_quadrants {
		if node.Quadrants[q] != nil {
			return fmt.Sprintf(
				"Node point to nil in q = %d, quadrants = %v",
				q, node.Quadrants,
			)
		}
	}

	return ""
}

func make_unoccupied_quadrants(occupied_quadrants []Quadrant) []Quadrant {
	res := []Quadrant{Lower_left, Lower_right, Upper_left, Upper_right}
	for _, q := range occupied_quadrants {
		i_remove := find_q_in_quadrants(q, res)
		res[i_remove] = res[len(res)-1]
		res = res[:len(res)-1]
	}
	return res
}

func find_q_in_quadrants(q Quadrant, quadrants []Quadrant) int {
	for i := range quadrants {
		if quadrants[i] == q {
			return i
		}
	}
	return -1
}
