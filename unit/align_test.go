package unit

import "testing"

func TestAlignDefault(t *testing.T) {
	a := AlignDefault
	if a.PadBetween.X != 6 || a.PadBetween.Y != 6 {
		t.Fatalf("unexpected default padBetween: %v", a.PadBetween)
	}
}

func TestNewAlignConstructors(t *testing.T) {
	a := NewAlignUniformPad(ContentCenter, ContentEnd, LineEnd, ModeVertical, NewVec2(5, 5), WrapNever)
	if a.PadEdges.X != 5 || a.PadBetween.X != 5 {
		t.Fatalf("pads not set correctly: %v", a)
	}
}
