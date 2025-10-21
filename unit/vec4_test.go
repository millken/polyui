package unit

import (
	"math"
	"testing"
)

func TestVec4Basics(t *testing.T) {
	v := OfVec4(1, 2, 3, 4)
	if v.X() != 1 || v.Y() != 2 || v.W() != 3 || v.H() != 4 {
		t.Fatalf("unexpected components: %v", v)
	}
	if math.Abs(float64(v.Magnitude2()-(1*1+2*2+3*3+4*4))) > 1e-6 {
		t.Fatalf("magnitude2 mismatch: %v", v.Magnitude2())
	}
}

func TestVec4SingleAndConstants(t *testing.T) {
	if Vec4ONE.X() != 1 || Vec4ONE.Y() != 1 || Vec4ONE.W() != 1 || Vec4ONE.H() != 1 {
		t.Fatalf("Vec4ONE mismatch")
	}
	if !Vec4ZERO.IsZero() {
		t.Fatalf("Vec4ZERO should be zero")
	}
}
