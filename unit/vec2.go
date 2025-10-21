package unit

import (
	"fmt"
	"math"
)

// Vec2 represents a 2D vector or size (X, Y).
type Vec2 struct {
	X, Y float32
}

func NewVec2(x, y float32) Vec2 { return Vec2{X: x, Y: y} }

var Zero = Vec2{0, 0}
var One = Vec2{1, 1}

func Of(x, y float32) Vec2 { return NewVec2(x, y) }

func (v Vec2) IsNegative() bool { return v.X < 0 || v.Y < 0 }
func (v Vec2) IsZero() bool     { return v.X == 0 && v.Y == 0 }
func (v Vec2) IsPositive() bool { return !v.IsNegative() && !v.IsZero() }

func (v Vec2) Magnitude2() float32 { return v.X*v.X + v.Y*v.Y }
func (v Vec2) Magnitude() float32  { return float32(math.Sqrt(float64(v.Magnitude2()))) }

func (v Vec2) Get(i int) float32 {
	if i == 0 {
		return v.X
	}
	if i == 1 {
		return v.Y
	}
	panic(fmt.Sprintf("index out of range: %d", i))
}

// CompareTo compares lexicographically by X then Y.
// returns -1 if v < other, 0 if equal, 1 if v > other
func (v Vec2) CompareTo(other Vec2) int {
	if v.X < other.X {
		return -1
	}
	if v.X > other.X {
		return 1
	}
	if v.Y < other.Y {
		return -1
	}
	if v.Y > other.Y {
		return 1
	}
	return 0
}

func (v Vec2) TimesWithRounding(sx, sy, rounding float32) Vec2 {
	return Vec2{RoundTo(v.X*sx, rounding), RoundTo(v.Y*sy, rounding)}
}

func (v Vec2) Mul(other Vec2) Vec2      { return Vec2{v.X * other.X, v.Y * other.Y} }
func (v Vec2) MulScalar(s float32) Vec2 { return Vec2{v.X * s, v.Y * s} }
func (v Vec2) Div(other Vec2) Vec2      { return Vec2{v.X / other.X, v.Y / other.Y} }
func (v Vec2) DivScalar(s float32) Vec2 { return Vec2{v.X / s, v.Y / s} }
func (v Vec2) Add(other Vec2) Vec2      { return Vec2{v.X + other.X, v.Y + other.Y} }
func (v Vec2) Sub(other Vec2) Vec2      { return Vec2{v.X - other.X, v.Y - other.Y} }
func (v Vec2) Rem(other Vec2) Vec2 {
	return Vec2{float32(math.Mod(float64(v.X), float64(other.X))), float32(math.Mod(float64(v.Y), float64(other.Y)))}
}

func (v Vec2) CoerceAtLeast(other Vec2) Vec2 {
	return Vec2{coerceAtLeastFloat(v.X, other.X), coerceAtLeastFloat(v.Y, other.Y)}
}
func (v Vec2) CoerceAtMost(other Vec2) Vec2 {
	return Vec2{coerceAtMostFloat(v.X, other.X), coerceAtMostFloat(v.Y, other.Y)}
}
func (v Vec2) CoerceIn(min, max Vec2) Vec2 {
	return Vec2{coerceInFloat(v.X, min.X, max.X), coerceInFloat(v.Y, min.Y, max.Y)}
}

func (v Vec2) String() string { return fmt.Sprintf("%gx%g", v.X, v.Y) }

// RoundTo rounds a float32 to the nearest multiple. If multiple <= 0, returns value.
func RoundTo(value, multiple float32) float32 {
	if multiple <= 0 {
		return value
	}
	return float32(math.Round(float64(value)/float64(multiple)) * float64(multiple))
}

func coerceAtLeastFloat(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func coerceAtMostFloat(a, b float32) float32 {
	if a > b {
		return b
	}
	return a
}

func coerceInFloat(v, a, b float32) float32 {
	if a < b {
		if v < a {
			return a
		}
		if v > b {
			return b
		}
		return v
	}
	if v < b {
		return b
	}
	if v > a {
		return a
	}
	return v
}
