package unit

import (
	"testing"
)

func TestTimeUnits(t *testing.T) {
	if Seconds(1) != 1_000_000_000 {
		t.Fatalf("seconds helper wrong: %v", Seconds(1))
	}
	if Ms(1) != 1_000_000 {
		t.Fatalf("ms helper wrong: %v", Ms(1))
	}
}

func TestFixDPS(t *testing.T) {
	if FixDPS(1.2345, 2) != 1.23 {
		t.Fatalf("FixDPS failed: %v", FixDPS(1.2345, 2))
	}
}
