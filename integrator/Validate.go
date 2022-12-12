package integrator

func (algorithm *Algorithm_t) validate() (is_valid bool, reason string) {

	is_valid = true
	reason = ""

	if len(algorithm.unique_force_coefficients) != len(algorithm.unique_fg_coefficients) {
		is_valid = false
		reason = add_invalid_reason(reason, "mismatch between length of force steps and force-gradient steps")
	}

	if (algorithm.is_velocity_version && len(algorithm.unique_velocity_coefficients) != len(algorithm.unique_force_coefficients)-1) ||
		(!algorithm.is_velocity_version && len(algorithm.unique_force_coefficients) != len(algorithm.unique_velocity_coefficients)-1) {
		is_valid = false
		reason = add_invalid_reason(reason, "mismatch between length of velocity steps and force steps")
	}

	if len(algorithm.unique_velocity_coefficients)+len(algorithm.unique_force_coefficients) != algorithm.stages {
		is_valid = false
		reason = add_invalid_reason(reason, "mismatch between number of stages and coefficients given to position and velocity steps")
	}

	manual_count_fg_steps := 0
	for _, c_fg := range algorithm.unique_fg_coefficients {
		if should_perform_fg_step(c_fg) {
			manual_count_fg_steps += 1
		}
	}
	if manual_count_fg_steps != algorithm.fg_steps {
		is_valid = false
		reason = add_invalid_reason(reason, "mismatch between number of fg_steps and non-zero coefficients")
	}

	return
}

func add_invalid_reason(original_reason, extra_reason string) (output_reason string) {
	if original_reason == "" {
		output_reason = extra_reason
	} else {
		output_reason = original_reason + ", " + extra_reason
	}
	return
}
