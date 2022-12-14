package plot_p5

import (
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/vector"
)

type Pulse_t struct {
	col                         color.Color
	min_size, delta_size        float64
	time_in_cycle, cycle_length int
	position                    vector.Vec
}

func New_pulse(col color.Color, length_of_cycle int,
	min_size, max_size float64) Pulse_t {
	return Pulse_t{
		col:           col,
		min_size:      min_size,
		delta_size:    max_size - min_size,
		time_in_cycle: int(length_of_cycle / 2),
		cycle_length:  length_of_cycle,
	}
}

func (pulse *Pulse_t) Reset_time(new_time int) {
	pulse.time_in_cycle = int(pulse.cycle_length/2) + new_time
}

func (pulse *Pulse_t) Update(new_position vector.Vec) {
	pulse.Update_time()
	pulse.Update_position(new_position)
}

func (pulse *Pulse_t) Update_position(new_position vector.Vec) {
	pulse.position = new_position
}

func (pulse *Pulse_t) Update_time() {
	pulse.time_in_cycle += 1
	if pulse.time_in_cycle > pulse.cycle_length {
		pulse.time_in_cycle -= pulse.cycle_length
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
	t := float64(pulse.time_in_cycle) / float64(pulse.cycle_length)
	return pulse.min_size + pulse.delta_size*math.Pow(math.Sin(math.Pi*t), 10)
}
