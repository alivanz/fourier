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
		out[i] = Dot(f, coef[i])
	}
	return out
}

func Dot(v1, v2 []complex128) complex128 {
	var total complex128
	for i := range v1 {
		total += v1[i] * v2[i]
	}
	return total
}
