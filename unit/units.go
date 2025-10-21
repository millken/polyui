package unit

// Time unit helpers following Kotlin semantics (smallest unit: nanosecond)
func Nanoseconds(n int64) int64    { return n }
func Microseconds(n float64) int64 { return int64(n * 1_000.0) }
func Milliseconds(n float64) int64 { return int64(n * 1_000_000.0) }
func Ms(n float64) int64           { return Milliseconds(n) }
func Seconds(n float64) int64      { return int64(n * 1_000_000_000.0) }
func Secs(n float64) int64         { return Seconds(n) }
func Minutes(n float64) int64      { return int64(n * 60_000_000_000.0) }
func Hours(n float64) int64        { return int64(n * 3_600_000_000_000.0) }

// Number to Vec2 convenience: in Go we'll accept float32
func Vec(n float32) Vec2 { return NewVec2(n, n) }

// Fix returns integer part as float32
func Fix(f float32) float32 { return float32(int(f)) }

// Fix with decimal places (dps)
func FixDPS(f float32, dps int) float32 {
	if dps < 0 {
		return f
	}
	if dps == 0 {
		return float32(int(f))
	}
	pow := float32(1)
	for i := 0; i < dps; i++ {
		pow *= 10
	}
	return float32(int(f*pow)) / pow
}

// Infix-by equivalents: we use plain functions in Go
func ByFloatFloat(a, b float32) Vec2          { return NewVec2(a, b) }
func ByFloatInt(a float32, b int) Vec2        { return NewVec2(a, float32(b)) }
func ByFloatDouble(a float32, b float64) Vec2 { return NewVec2(a, float32(b)) }
func ByIntFloat(a int, b float32) Vec2        { return NewVec2(float32(a), b) }
func ByIntInt(a, b int) Vec2                  { return NewVec2(float32(a), float32(b)) }
func ByIntDouble(a int, b float64) Vec2       { return NewVec2(float32(a), float32(b)) }
func ByDoubleFloat(a float64, b float32) Vec2 { return NewVec2(float32(a), b) }
func ByDoubleInt(a float64, b int) Vec2       { return NewVec2(float32(a), float32(b)) }
func ByDoubleDouble(a, b float64) Vec2        { return NewVec2(float32(a), float32(b)) }

// Vec2 to Vec4 helper
func (v Vec2) ToVec4() Vec4 { return OfVec4FromVec2(v, v) }
