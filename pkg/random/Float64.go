package random

import (
	"math"
	"math/rand"

	"github.com/granite/pkg/vector"
)

func Position(D_max float64) vector.Vec {
	return vector.New(Signed_float64(D_max), Signed_float64(D_max))
}

func Velocity(v_max float64) vector.Vec {
	v_abs := Non_negative_float64(v_max)
	phi := Non_negative_float64(math.Pi * 2.0)
	return vector.Cartesian_position_from_polar(v_abs, phi)
}

func Signed_float64(x_abs_max float64) float64 {
	return sign() * Non_negative_float64(x_abs_max)
}

func Non_negative_float64(x_max float64) float64 {
	return rand.Float64() * x_max
}

func sign() float64 {
	return float64(sign_int())
}

func sign_int() int {
	return 2*rand.Intn(2) - 1
}
