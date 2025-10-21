package unit

import (
	"math"
	"testing"
)

func floatEq(a, b float32) bool { return math.Abs(float64(a-b)) < 1e-6 }

func TestVec2Basics(t *testing.T) {
	a := NewVec2(3, 4)
	if !floatEq(a.Magnitude(), 5) {
		t.Fatalf("expected magnitude 5 got %v", a.Magnitude())
	}

	b := NewVec2(1, 2)
	c := a.Sub(b)
	if c.X != 2 || c.Y != 2 {
		t.Fatalf("unexpected sub result: %v", c)
	}

	d := a.MulScalar(2)
	if d.X != 6 || d.Y != 8 {
		t.Fatalf("unexpected mul scalar: %v", d)
	}
}

func TestCompareToAndRounding(t *testing.T) {
	a := NewVec2(1, 1)
	b := NewVec2(2, 0)
	if a.CompareTo(b) != -1 {
		t.Fatalf("expected a < b")
	}
	if b.CompareTo(a) != 1 {
		t.Fatalf("expected b > a")
	}

	v := NewVec2(1.2345, 2.3456)
	r := v.TimesWithRounding(1, 1, 0.1)
	if !floatEq(r.X, RoundTo(1.2345, 0.1)) {
		t.Fatalf("round mismatch x: %v vs %v", r.X, RoundTo(1.2345, 0.1))
	}
}

func TestCoerce(t *testing.T) {
	v := NewVec2(5, -3)
	min := NewVec2(0, 0)
	max := NewVec2(4, 4)
	c := v.CoerceIn(min, max)
	if c.X != 4 || c.Y != 0 {
		t.Fatalf("unexpected coerce result: %v", c)
	}
}
