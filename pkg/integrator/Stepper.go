package integrator

import "github.com/granite/pkg/physics"

type Stepper_t struct {
	algorithm Algorithm_t
	steps     []single_step_t
}

func New_stepper(algorithm Algorithm_t) (stepper Stepper_t) {
	stepper = Stepper_t{
		algorithm: algorithm,
		steps:     make_steps(algorithm),
	}
	return
}

func (stepper *Stepper_t) Run(system *physics.System_t, timestep float64) {
	for i := range stepper.steps {
		stepper.steps[i].run(system, timestep)
	}
}

func (stepper Stepper_t) String() (output string) {
	output = stepper.algorithm.String() + "\nExponential form:\n"
	for _, s := range stepper.steps {
		output += s.String() + " "
	}
	return
}

func make_steps(algorithm Algorithm_t) (steps []single_step_t) {

	v_coeffs := algorithm.velocity_stages()
	f_coeffs := algorithm.force_stages()
	fg_coeffs := algorithm.force_gradient_stages()

	steps = make([]single_step_t, 0, algorithm.stages)

	if !algorithm.is_velocity_version {
		steps = append(steps, velocity_step(v_coeffs[0]))
		v_coeffs = v_coeffs[1:]
	}
	for i := range v_coeffs {
		steps = append(steps, force_or_fg_step(f_coeffs[i], fg_coeffs[i]))
		steps = append(steps, velocity_step(v_coeffs[i]))
	}
	if algorithm.is_velocity_version {
		steps = append(steps, force_or_fg_step(f_coeffs[len(f_coeffs)-1], fg_coeffs[len(fg_coeffs)-1]))
	}

	return
}

func velocity_step(coefficient float64) single_step_t {
	return &velocity_step_t{
		coefficient: coefficient,
	}
}

func force_or_fg_step(c_f, c_fg float64) (step single_step_t) {
	if should_perform_fg_step(c_fg) {
		step = &force_and_force_gradient_step_t{
			f_coefficient:  c_f,
			fg_coefficient: c_fg,
		}
	} else {
		step = &force_step_t{
			coefficient: c_f,
		}
	}
	return
}
