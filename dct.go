package fourier

import (
	"math"
)

func NewDCT(n int) *CoefPairComplex {
	trans := make(CoefComplex, n)
	inv := make(CoefComplex, n)
	for i := 0; i < n; i++ {
		trans[i] = make([]complex128, n)
		inv[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				trans[i][j] = complex(
					1.0/float64(n),
					0,
				)
			} else {
				trans[i][j] = complex(
					2/float64(n)*math.Cos(math.Pi*float64(i)*(float64(j)+0.5)/float64(n)),
					0,
				)
			}
			if j == 0 {
				inv[i][j] = 1
			} else {
				inv[i][j] = complex(
					math.Cos(math.Pi*(float64(i)+0.5)*float64(j)/float64(n)),
					0,
				)
			}
		}
	}
	return &CoefPairComplex{
		Transform: trans,
		Inverse:   inv,
	}
}

func IDCT(fs []complex128, x float64) complex128 {
	var total complex128
	l := float64(len(fs))
	for i, f := range fs {
		total += f * complex(
			math.Cos(math.Pi*(float64(x)+0.5)*float64(i)/l),
			0,
		)
	}
	return total
}
