package fourier

import (
	"math"
	"math/rand"
	"testing"
)

func TestFFT(t *testing.T) {
	in := []complex128{
		complex(rand.Float64(), rand.Float64()),
		complex(rand.Float64(), rand.Float64()),
		complex(rand.Float64(), rand.Float64()),
		complex(rand.Float64(), rand.Float64()),
		complex(rand.Float64(), rand.Float64()),
	}
	t.Log(in)
	out := FFT(in)
	t.Log(out)
	inv := IFFT(out)
	t.Log(inv)
	for i, v := range inv {
		diff := round(v - in[i])
		if math.Abs(real(diff)) > 0 {
			t.Fail()
		}
		if math.Abs(imag(diff)) > 0 {
			t.Fail()
		}
	}
}
