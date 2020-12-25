package fourier

import (
	"math"
)

func NewFFT(n int) *CoefPairComplex {
	coef := make([][]complex128, n)
	for i := 0; i < n; i++ {
		coef[i] = FFTCoefs(i, n)
	}
	inv := make([][]complex128, n)
	for i := 0; i < n; i++ {
		inv[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			inv[i][j] = conjugate(coef[i][j])
		}
	}
	return &CoefPairComplex{
		Transform: coef,
		Inverse:   inv,
	}
}

func Expi(x float64) complex128 {
	return complex(
		math.Cos(x),
		math.Sin(x),
	)
}

func FFTCoef(i, x, n int) complex128 {
	return Expi(-2 * math.Pi * float64(i) * float64(x) / float64(n))
}

func FFTCoefs(x, n int) []complex128 {
	out := make([]complex128, n)
	for i := 0; i < n; i++ {
		out[i] = FFTCoef(i, x, n)
	}
	return out
}

func FFT(f []complex128) []complex128 {
	fft := NewFFT(len(f))
	return fft.Transform.Do(f)
}

func IFFT(fs []complex128) []complex128 {
	fft := NewFFT(len(fs))
	return fft.Inverse.Do(fs)
}

func IFFTOne(fs []complex128, x int) complex128 {
	n := len(fs)
	var total complex128
	for i, c := range FFTCoefs(x, n) {
		total += fs[i] * c
	}
	return total
}
