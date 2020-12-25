package fourier

import (
	"math"
)

func FFT(f []complex128) []complex128 {
	l := len(f)
	out := make([]complex128, l)
	for i := 0; i < l; i++ {
		var re, im float64
		for x, v := range f {
			angle := -2 * math.Pi * float64(i) * float64(x) / float64(l)
			re += real(v)*math.Cos(angle) - imag(v)*math.Sin(angle)
			im += real(v)*math.Sin(angle) + imag(v)*math.Cos(angle)
		}
		out[i] = complex(re/float64(l), im/float64(l))
	}
	return out
}

func IFFT(fs []complex128) []complex128 {
	l := len(fs)
	out := make([]complex128, l)
	for x := range out {
		out[x] = IFFTOne(fs, x)
	}
	return out
}

func IFFTOne(fs []complex128, x int) complex128 {
	l := len(fs)
	var re, im float64
	for i, f := range fs {
		angle := 2 * math.Pi * float64(i) * float64(x) / float64(l)
		re += real(f)*math.Cos(angle) - imag(f)*math.Sin(angle)
		im += real(f)*math.Sin(angle) + imag(f)*math.Cos(angle)
	}
	return complex(re, im)
}
