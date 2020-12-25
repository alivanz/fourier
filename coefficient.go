package fourier

type CoefComplex [][]complex128

type CoefPairComplex struct {
	Transform CoefComplex
	Inverse   CoefComplex
}

func (coef CoefComplex) Do(f []complex128) []complex128 {
	n := len(coef)
	out := make([]complex128, n)
	for i := 0; i < n; i++ {
		var total complex128
		for x, v := range f {
			total += v * coef[i][x]
		}
		out[i] = total
	}
	return out
}
