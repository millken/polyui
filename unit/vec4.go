package unit

import (
	"fmt"
	"math"
)

// Vec4 represents a 4D vector (x,y,w,h). It mirrors the Kotlin Vec4 API in a Go-friendly way.
type Vec4 interface {
	X() float32
	Y() float32
	W() float32
	H() float32

	IsNegative() bool
	IsZero() bool
	Magnitude2() float32
	Magnitude() float32
	Get(i int) float32
	String() string
	Copy(x, y, w, h float32) Vec4
}

// impl holds two Vec2 values (xy and wh) like the Kotlin Impl class.
type vec4Impl struct {
	xy Vec2
	wh Vec2
}

func (v *vec4Impl) X() float32 { return v.xy.X }
func (v *vec4Impl) Y() float32 { return v.xy.Y }
func (v *vec4Impl) W() float32 { return v.wh.X }
func (v *vec4Impl) H() float32 { return v.wh.Y }

func (v *vec4Impl) IsNegative() bool    { return v.xy.IsNegative() && v.wh.IsNegative() }
func (v *vec4Impl) IsZero() bool        { return v.xy.IsZero() && v.wh.IsZero() }
func (v *vec4Impl) Magnitude2() float32 { return v.xy.Magnitude2() + v.wh.Magnitude2() }
func (v *vec4Impl) Magnitude() float32  { return float32(math.Sqrt(float64(v.Magnitude2()))) }
func (v *vec4Impl) Get(i int) float32 {
	switch i {
	case 0:
		return v.X()
	case 1:
		return v.Y()
	case 2:
		return v.W()
	case 3:
		return v.H()
	}
	panic(fmt.Sprintf("index out of range: %d", i))
}

func (v *vec4Impl) String() string { return fmt.Sprintf("%gx%gx%gx%g", v.X(), v.Y(), v.W(), v.H()) }

func (v *vec4Impl) Copy(x, y, w, h float32) Vec4 { return OfVec4(x, y, w, h) }

// singleVec4 represents the Kotlin Single variant where all components are equal.
type singleVec4 struct {
	x float32
}

func (s *singleVec4) X() float32                   { return s.x }
func (s *singleVec4) Y() float32                   { return s.x }
func (s *singleVec4) W() float32                   { return s.x }
func (s *singleVec4) H() float32                   { return s.x }
func (s *singleVec4) IsNegative() bool             { return s.x < 0 }
func (s *singleVec4) IsZero() bool                 { return s.x == 0 }
func (s *singleVec4) Magnitude2() float32          { return (s.x * s.x) * 4 }
func (s *singleVec4) Magnitude() float32           { return float32(math.Sqrt(float64(s.Magnitude2()))) }
func (s *singleVec4) Get(i int) float32            { return s.x }
func (s *singleVec4) String() string               { return fmt.Sprintf("%g", s.x) }
func (s *singleVec4) Copy(x, y, w, h float32) Vec4 { return OfVec4(x, y, w, h) }

// Constructors and constants
var (
	Vec4ONE  Vec4 = &singleVec4{x: 1}
	Vec4ZERO Vec4 = &singleVec4{x: 0}
)

func OfVec4(x, y, w, h float32) Vec4      { return &vec4Impl{xy: NewVec2(x, y), wh: NewVec2(w, h)} }
func OfVec4FromVec2(xy, wh Vec2) Vec4     { return &vec4Impl{xy: xy, wh: wh} }
func OfVec4FromAtSize(at, size Vec2) Vec4 { return &vec4Impl{xy: at, wh: size} }
