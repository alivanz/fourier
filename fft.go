package fourier

import (
	"math"
)

func NewFFT(n int) *CoefPairComplex {
	inv := make([][]complex128, n)
	for i := 0; i < n; i++ {
		inv[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			inv[i][j] = Expi(2 * math.Pi * float64(i) * float64(j) / float64(n))
		}
	}
	l := float64(n)
	coef := make([][]complex128, n)
	for i := 0; i < n; i++ {
		coef[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			c := conjugate(inv[i][j])
			coef[i][j] = div(c, l)
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
