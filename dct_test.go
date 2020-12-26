package fourier

import (
	"reflect"
	"testing"
)

func TestDCT(t *testing.T) {
	in := []complex128{
		complex(1, 0),
		complex(2, 0),
		complex(3, 0),
		complex(4, 0),
	}
	t.Log(in)
	dct := NewDCT(len(in))
	out := dct.Transform.Do(in)
	t.Log(out)
	inv := dct.Inverse.Do(out)
	rounds(inv)
	t.Log(inv)
	if !reflect.DeepEqual(in, inv) {
		t.Fail()
	}
}
