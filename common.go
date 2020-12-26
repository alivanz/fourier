package fourier

import (
	"math"
)

func Expi(x float64) complex128 {
	return complex(
		math.Cos(x),
		math.Sin(x),
	)
}

func conjugate(c complex128) complex128 {
	return complex(real(c), -imag(c))
}

func div(c complex128, d float64) complex128 {
	return complex(real(c)/d, imag(c)/d)
}

func rounds(vs []complex128) {
	for i, v := range vs {
		vs[i] = round(v)
	}
}
func round(v complex128) complex128 {
	return complex(
		roundFloat(real(v)),
		roundFloat(imag(v)),
	)
}
func roundFloat(v float64) float64 {
	return math.Round(v*1000000) / 1000000
}
