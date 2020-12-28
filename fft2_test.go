package fourier

import (
	"reflect"
	"testing"
)

func TestFFT2(t *testing.T) {
	fft := NewFFT(5)
	in := [][]complex128{
		[]complex128{1, 2, 3, 4, 5},
		[]complex128{2, 2, 3, 4, 5},
		[]complex128{1, 3, 5, 4, 5},
		[]complex128{1, 2, 3, 4, 5},
		[]complex128{1, 3, 5, 4, 1},
	}
	// transform
	out := make([][]complex128, 5)
	out[0] = fft.Transform.Do(in[0])
	out[1] = fft.Transform.Do(in[1])
	out[2] = fft.Transform.Do(in[2])
	out[3] = fft.Transform.Do(in[3])
	out[4] = fft.Transform.Do(in[4])
	out = Transpose(out)
	out[0] = fft.Transform.Do(in[0])
	out[1] = fft.Transform.Do(in[1])
	out[2] = fft.Transform.Do(in[2])
	out[3] = fft.Transform.Do(in[3])
	out[4] = fft.Transform.Do(in[4])
	// inverse
	inv := make([][]complex128, 5)
	inv[0] = fft.Inverse.Do(out[0])
	inv[1] = fft.Inverse.Do(out[1])
	inv[2] = fft.Inverse.Do(out[2])
	inv[3] = fft.Inverse.Do(out[3])
	inv[4] = fft.Inverse.Do(out[4])
	inv = Transpose(inv)
	inv[0] = fft.Inverse.Do(out[0])
	inv[1] = fft.Inverse.Do(out[1])
	inv[2] = fft.Inverse.Do(out[2])
	inv[3] = fft.Inverse.Do(out[3])
	inv[4] = fft.Inverse.Do(out[4])
	rounds(inv[0])
	rounds(inv[1])
	rounds(inv[2])
	rounds(inv[3])
	rounds(inv[4])
	t.Log(inv)
	// check
	if !reflect.DeepEqual(in, inv) {
		t.Fail()
	}
}
