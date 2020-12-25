package fourier

import (
	"math"
)

type FFTBox struct {
	n    int
	coef [][]complex128
}

func NewFFT(n int) *FFTBox {
	fft := &FFTBox{
		n:    n,
		coef: make([][]complex128, n),
	}
	for i := 0; i < n; i++ {
		fft.coef[i] = make([]complex128, n)
		for x := 0; x < n; x++ {
			fft.coef[i][x] = FFTCoef(i, x, n)
		}
	}
	return fft
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

func (fft *FFTBox) Coefficient() [][]complex128 {
	return fft.coef
}

func (fft *FFTBox) FFT(f []complex128) []complex128 {
	out := make([]complex128, fft.n)
	for i := 0; i < fft.n; i++ {
		var total complex128
		for x, v := range f {
			total += v * fft.coef[i][x]
		}
		out[i] = complex(real(total)/float64(fft.n), imag(total)/float64(fft.n))
	}
	return out
}

func (fft *FFTBox) IFFT(fs []complex128) []complex128 {
	out := make([]complex128, fft.n)
	for i := 0; i < fft.n; i++ {
		var total complex128
		for x, v := range fs {
			total += v * conjugate(fft.coef[i][x])
		}
		out[i] = total
	}
	return out
}

func FFT(f []complex128) []complex128 {
	fft := NewFFT(len(f))
	return fft.FFT(f)
}

func IFFT(fs []complex128) []complex128 {
	fft := NewFFT(len(fs))
	return fft.IFFT(fs)
}

func IFFTOne(fs []complex128, x int) complex128 {
	n := len(fs)
	var total complex128
	for i, f := range fs {
		total += f * FFTCoef(i, x, n)
	}
	return total
}
