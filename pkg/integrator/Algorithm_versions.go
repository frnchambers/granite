package integrator

import "math"

func Default_algorithm() Algorithm_t {
	return Default_O4_algorithm()
}

func Choose_algorithm(required_error_order int) (algo Algorithm_t) {
	if required_error_order >= 6 {
		algo = Version_3_4_1_v_1()
	} else if required_error_order >= 4 {
		algo = Version_3_2_1_v_2()
	} else {
		algo = Version_3_1_1_v_2()
	}
	return
}

func Default_O2_algorithm() Algorithm_t {
	return Version_3_1_1_v_2()
}

func Default_O4_algorithm() Algorithm_t {
	return Version_3_2_1_v_2()
}

func Default_O6_algorithm() Algorithm_t {
	return Version_3_4_1_v_1()
}

func Velocity_verlet_algorithm() Algorithm_t {
	return Version_3_1_1_v_2()
}

/* --------------------------- 3-stage algorithms --------------------------- */

// func Version_3_1_1_v_1() stepper_params_t {
// 	return stepper_params_t{
// 		is_velocity_first:     true,
// 		stages:                3,
// 		fg_steps:              1,
// 		order:                 2,
// 		velocity_coefficients: []float64{1.0},
// 		force_coefficients:    []float64{0.5},
// 		fg_coefficients:       []float64{-1.0 / 48.0},
// 	}
// }

func Version_3_1_1_v_2() Algorithm_t { // velocity-verlet algorithm
	return Algorithm_t{
		is_velocity_version:          true,
		stages:                       3,
		fg_steps:                     0,
		error_order:                  2,
		unique_velocity_coefficients: []float64{1.0},
		unique_force_coefficients:    []float64{0.5},
		unique_fg_coefficients:       []float64{0.0},
	}
}

func Version_3_1_2_v_1() Algorithm_t {
	return Algorithm_t{
		is_velocity_version:          false,
		stages:                       3,
		fg_steps:                     1,
		error_order:                  2,
		unique_velocity_coefficients: []float64{0.5},
		unique_force_coefficients:    []float64{1.0},
		unique_fg_coefficients:       []float64{-1.0 / 48.0},
	}
}

// func Version_3_1_2_v_2() stepper_params_t {
// 	return stepper_params_t{
// 		is_velocity_first:     false,
// 		stages:                3,
// 		fg_steps:              0,
// 		order:                 2,
// 		velocity_coefficients: []float64{0.5},
// 		force_coefficients:    []float64{1.0},
// 		fg_coefficients:       []float64{0.0},
// 	}
// }

/* --------------------------- 3-stage algorithms --------------------------- */

// func Version_3_2_1_v_1() stepper_params_t {
// 	lambda := 1.0 / 6.0
// 	return stepper_params_t{
// 		is_velocity_first:     true,
// 		stages:                5,
// 		fg_steps:              2,
// 		order:                 4,
// 		velocity_coefficients: []float64{0.5},
// 		force_coefficients:    []float64{lambda, 1.0 - 2.0*lambda},
// 		fg_coefficients:       []float64{-17.0 / 18000.0, 71.0 / 4500},
// 	}
// }

func Version_3_2_1_v_2() Algorithm_t {
	lambda := 1.0 / 6.0
	return Algorithm_t{
		is_velocity_version:          true,
		stages:                       5,
		fg_steps:                     1,
		error_order:                  4,
		unique_velocity_coefficients: []float64{0.5},
		unique_force_coefficients:    []float64{lambda, 1.0 - 2.0*lambda},
		unique_fg_coefficients:       []float64{0.0, 1.0 / 72.0},
	}
}

// func Version_3_2_1_v_3() stepper_params_t {
// 	lambda := 1.0 / 6.0
// 	return stepper_params_t{
// 		is_velocity_first:     true,
// 		stages:                5,
// 		fg_steps:              1,
// 		order:                 4,
// 		velocity_coefficients: []float64{0.5},
// 		force_coefficients:    []float64{lambda, 1.0 - 2.0*lambda},
// 		fg_coefficients:       []float64{1.0 / 144.0, 0.0},
// 	}
// }

// func Version_3_2_1_v_4() stepper_params_t {
// 	lambda := 0.1931833275037836
// 	return stepper_params_t{
// 		is_velocity_first:     true,
// 		stages:                5,
// 		fg_steps:              0,
// 		order:                 2,
// 		velocity_coefficients: []float64{0.5},
// 		force_coefficients:    []float64{lambda, 1.0 - 2.0*lambda},
// 		fg_coefficients:       []float64{0.0, 0.0},
// 	}
// }

// func Version_3_2_2_v_1() stepper_params_t {
// 	lambda := 0.5 * (1.0 - 1.0/math.Sqrt(3.0))
// 	xi := (2.0 - math.Sqrt(3.0)) / 48.0
// 	return stepper_params_t{
// 		is_velocity_first:     false,
// 		stages:                5,
// 		fg_steps:              2,
// 		order:                 4,
// 		velocity_coefficients: []float64{lambda, 1.0 - 2.0*lambda},
// 		force_coefficients:    []float64{0.5},
// 		fg_coefficients:       []float64{xi},
// 	}
// }

// func Version_3_2_2_v_2() stepper_params_t {
// 	lambda := 0.5 * (1.0 + 1.0/math.Sqrt(3.0))
// 	xi := (2.0 + math.Sqrt(3.0)) / 48.0
// 	return stepper_params_t{
// 		is_velocity_first:     false,
// 		stages:                5,
// 		fg_steps:              2,
// 		order:                 4,
// 		velocity_coefficients: []float64{lambda, 1.0 - 2.0*lambda},
// 		force_coefficients:    []float64{0.5},
// 		fg_coefficients:       []float64{xi},
// 	}
// }

/* --------------------------- 9-stage algorithms --------------------------- */

func Version_3_4_1_v_1() Algorithm_t {
	cube_root := math.Cbrt(675.0 + 75*math.Sqrt(6.0))
	theta := 0.5 + cube_root/30.0 + 5.0/(2.0*cube_root)
	phi := theta / 3.0
	lambda := -5.0 / 3.0 * theta * (theta - 1.0)
	xi := -5.0*theta*theta/144.0 + theta/36.0 - 1.0/288.0
	chi := 1.0/144.0 - theta/36.0*(0.5*theta+1.0)
	mu := 0.0
	return Algorithm_t{
		is_velocity_version:          true,
		stages:                       9,
		fg_steps:                     3,
		error_order:                  6,
		unique_velocity_coefficients: []float64{theta, 0.5 - theta},
		unique_force_coefficients:    []float64{phi, lambda, 1.0 - 2.0*(lambda+phi)},
		unique_fg_coefficients:       []float64{mu, xi, chi},
	}
}

func Version_3_4_1_v_2() Algorithm_t {
	theta := 0.1705755127786631
	phi := 0.4775180236616381e-1
	lambda := 0.2739456420927671
	xi := 0.2464531166166595e-2
	mu := -0.6175944713542174e-3
	chi := 0.0
	return Algorithm_t{
		is_velocity_version:          true,
		stages:                       9,
		fg_steps:                     3,
		error_order:                  4,
		unique_velocity_coefficients: []float64{theta, 0.5 - theta},
		unique_force_coefficients:    []float64{phi, lambda, 1.0 - 2.0*(lambda+phi)},
		unique_fg_coefficients:       []float64{mu, xi, chi},
	}
}

/* -------------------------- 11-stage algorithms --------------------------- */

func Version_3_5_1_v_2() Algorithm_t {
	root_five := math.Sqrt(5.0)
	nest_root := math.Sqrt(50.0 + 22.0*root_five)
	xi_chi_inc := nest_root * (1.0/2880.0 + root_five/1152.0)
	rho := 0.5 * (1.0 + 1.0/root_five)
	theta := -1.0 / root_five
	phi := 1.0 / 12.0
	lambda := 5.0/12.0 - nest_root/24.0
	xi := (15.0+5.0*math.Sqrt(5.0))/1152.0 - xi_chi_inc
	chi := -(11.0+5.0*math.Sqrt(5.0))/1152.0 + xi_chi_inc
	mu := 0.0
	return Algorithm_t{
		is_velocity_version:          true,
		stages:                       11,
		fg_steps:                     4,
		error_order:                  6,
		unique_velocity_coefficients: []float64{rho, theta, 1.0 - 2.0*(rho+theta)},
		unique_force_coefficients:    []float64{phi, lambda, 0.5 - (lambda + phi)},
		unique_fg_coefficients:       []float64{mu, xi, chi},
	}
}

func Version_3_5_1_v_3() Algorithm_t {
	rho := 0.2742082240034209
	theta := 0.4812780570021632
	phi := 8.350330494925359e-2
	lambda := 4.474919773539384e-1
	xi := 0.0
	chi := 3.435650653755542e-2
	mu := -2.544189176362832e4
	return Algorithm_t{
		is_velocity_version:          true,
		stages:                       11,
		fg_steps:                     3,
		error_order:                  6,
		unique_velocity_coefficients: []float64{rho, theta, 1.0 - 2.0*(rho+theta)},
		unique_force_coefficients:    []float64{phi, lambda, 0.5 - (lambda + phi)},
		unique_fg_coefficients:       []float64{mu, xi, chi},
	}
}
