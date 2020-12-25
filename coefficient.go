package fourier

type CoefComplex [][]complex128

func (coef CoefComplex) Transform(f []complex128) []complex128 {
	n := len(coef)
	out := make([]complex128, n)
	for i := 0; i < n; i++ {
		var total complex128
		for x, v := range f {
			total += v * coef[i][x]
		}
		out[i] = complex(real(total)/float64(n), imag(total)/float64(n))
	}
	return out
}

func (coef CoefComplex) Inverse(fs []complex128) []complex128 {
	n := len(coef)
	out := make([]complex128, n)
	for i := 0; i < n; i++ {
		var total complex128
		for x, v := range fs {
			total += v * conjugate(coef[i][x])
		}
		out[i] = total
	}
	return out
}
