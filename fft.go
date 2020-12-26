package fourier

import (
	"math"
)

func NewFFT(n int) *CoefPairComplex {
	coef := make([][]complex128, n)
	inv := make([][]complex128, n)
	for i := 0; i < n; i++ {
		coef[i] = make([]complex128, n)
		inv[i] = make([]complex128, n)
	}
	l := float64(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c := Expi(-2 * math.Pi * float64(i) * float64(j) / float64(n))
			coef[i][j] = div(c, l)
			inv[i][j] = conjugate(c)
		}
	}
	return &CoefPairComplex{
		Transform: coef,
		Inverse:   inv,
	}
}

func IFFT(fs []complex128, x float64) complex128 {
	var total complex128
	l := float64(len(fs))
	for i, f := range fs {
		total += f * Expi(2*math.Pi*float64(i)*x/l)
	}
	return total
}
