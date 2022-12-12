package plot_p5

import (
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/vector"
)

type Pulse_t struct {
	col                            color.Color
	min_size, max_size             float64
	time_in_cycle, length_of_cycle int
	position                       vector.Vec
}

func New_pulse(col color.Color, position vector.Vec, length_of_cycle int,
	min_size, max_size float64) Pulse_t {
	return Pulse_t{
		col:             col,
		min_size:        min_size,
		max_size:        max_size,
		time_in_cycle:   0,
		length_of_cycle: length_of_cycle,
		position:        position,
	}
}

func (pulse *Pulse_t) Update() {
	// pulse.position = position
	pulse.time_in_cycle += 1
	if pulse.time_in_cycle > pulse.length_of_cycle {
		pulse.time_in_cycle -= pulse.length_of_cycle
	}

}

func (pulse *Pulse_t) Plot() {
	p5.Stroke(pulse.col)
	p5.Fill(pulse.col)
	p5.Circle(pulse.position.X, -pulse.position.Y, pulse.diameter())
}

func (pulse *Pulse_t) Position() vector.Vec {
	return pulse.position
}

func (pulse *Pulse_t) diameter() float64 {
	t := float64(pulse.time_in_cycle) / float64(pulse.length_of_cycle)
	return pulse.min_size + (pulse.max_size-pulse.min_size)*math.Pow(math.Sin(2*math.Pi*t), 2)
}
