package plot_p5

import (
	"image/color"
	"math"

	"github.com/go-p5/p5"
	"github.com/granite/pkg/vector"
)

type Pulse_t struct {
	col                  color.Color
	min_size, delta_size float64
	time, period         float64
	jerk                 float64
	position             vector.Vec
}

func New_pulse(col color.Color, period, min_size, max_size float64) Pulse_t {
	return Pulse_t{
		col:        col,
		min_size:   min_size,
		delta_size: max_size - min_size,
		time:       period / 2,
		period:     period,
		jerk:       10.0,
	}
}

func (pulse *Pulse_t) Reset_time(new_time float64) {
	pulse.time = pulse.period/2.0 + new_time
}

func (pulse *Pulse_t) Update(step_size float64, new_position vector.Vec) {
	pulse.Update_time(step_size)
	pulse.Update_position(new_position)
}

func (pulse *Pulse_t) Update_position(new_position vector.Vec) {
	pulse.position = new_position
}

func (pulse *Pulse_t) Update_time(step_size float64) {
	pulse.time += step_size
	if pulse.time > pulse.period {
		pulse.time -= pulse.period
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
	return pulse.min_size + pulse.delta_size*math.Pow(math.Sin(math.Pi*pulse.time/pulse.period), pulse.jerk)
}
