package fourier

import (
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
	fft := NewFFT(len(in))
	out := fft.Transform.Do(in)
	t.Log(out)
	inv := fft.Inverse.Do(out)
	t.Log(inv)
	for i, v := range inv {
		diff := round(v - in[i])
		if diff != 0 {
			t.Fail()
		}
	}
	// Test Cont
	for i := 0; i < len(in)*2; i++ {
		v := IFFT(out, float64(i))
		diff := round(v - in[i%len(in)])
		if diff != 0 {
			t.Fail()
		}
	}
}
