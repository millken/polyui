package color

import "math"

// Color is the exported interface representing a color's behaviour.
type Color interface {
	Hue() float32
	Saturation() float32
	Brightness() float32
	Alpha() float32
	ARGB() int32
	RGBA() int32
	Transparent() bool
	Red() int32
	Green() int32
	Blue() int32
	AlphaInt() int32
}

// PolyColor is a concrete base struct that implementations can embed to
// provide common functionality. It is intentionally minimal; concrete types
// like StaticColor, MutableColor embed it and implement the Color interface.
type PolyColor struct{}

// Dynamic represents a color which has some kind of operation that updates it based on time, for example, an animation.
type Dynamic interface {
	Update(deltaTimeNanos int64) bool
}

// Mut represents a mutable color that can be recolored
type Mut interface {
	Recolor(to Color) Mut
}

// Animatable represents a color that can be animated
type Animatable interface {
	Mut
	RecolorWithAnimation(to Color, animation interface{}) Animatable // TODO: Animation dependency
}

// StaticColor is a static color implementation
type StaticColor struct {
	hue        float32
	saturation float32
	brightness float32
	alpha      float32
	argb       int32
}

// NewStaticColor creates a new static color
func NewStaticColor(hue, saturation, brightness, alpha float32) *StaticColor {
	c := &StaticColor{
		hue:        float32(math.Mod(float64(hue), 360.0)),
		saturation: clampf32(saturation, 0, 1),
		brightness: clampf32(brightness, 0, 1),
		alpha:      clampf32(alpha, 0, 1),
	}
	c.argb = HSBAToRGB(c.hue, c.saturation, c.brightness, c.alpha)
	return c
}

func (c *StaticColor) Hue() float32        { return c.hue }
func (c *StaticColor) Saturation() float32 { return c.saturation }
func (c *StaticColor) Brightness() float32 { return c.brightness }
func (c *StaticColor) Alpha() float32      { return c.alpha }
func (c *StaticColor) ARGB() int32         { return c.argb }
func (c *StaticColor) RGBA() int32         { return (c.argb << 8) | (c.argb>>24)&0xFF }
func (c *StaticColor) Transparent() bool   { return c.alpha == 0 }
func (c *StaticColor) Red() int32          { return (c.argb >> 16) & 0xFF }
func (c *StaticColor) Green() int32        { return (c.argb >> 8) & 0xFF }
func (c *StaticColor) Blue() int32         { return c.argb & 0xFF }
func (c *StaticColor) AlphaInt() int32     { return (c.argb >> 24) & 0xFF }

// MutableColor is a mutable color implementation
type MutableColor struct {
	hue        float32
	saturation float32
	brightness float32
	alpha      float32
	argb       int32
	dirty      bool
}

// NewMutableColor creates a new mutable color
func NewMutableColor(hue, saturation, brightness, alpha float32) *MutableColor {
	c := &MutableColor{
		hue:        float32(math.Mod(float64(hue), 360.0)),
		saturation: clampf32(saturation, 0, 1),
		brightness: clampf32(brightness, 0, 1),
		alpha:      clampf32(alpha, 0, 1),
		dirty:      true,
	}
	c.argb = HSBAToRGB(c.hue, c.saturation, c.brightness, c.alpha)
	c.dirty = false
	return c
}

func (c *MutableColor) Hue() float32 { return c.hue }

func (c *MutableColor) SetHue(value float32) {
	c.dirty = true
	c.hue = float32(math.Mod(float64(value), 360.0))
}

func (c *MutableColor) Saturation() float32 { return c.saturation }

func (c *MutableColor) SetSaturation(value float32) {
	c.dirty = true
	c.saturation = clampf32(value, 0, 1)
}

func (c *MutableColor) Brightness() float32 { return c.brightness }

func (c *MutableColor) SetBrightness(value float32) {
	c.dirty = true
	c.brightness = clampf32(value, 0, 1)
}

func (c *MutableColor) Alpha() float32 { return c.alpha }

func (c *MutableColor) SetAlpha(value float32) {
	c.dirty = true
	c.alpha = clampf32(value, 0, 1)
}

func (c *MutableColor) ARGB() int32 {
	if c.dirty {
		c.argb = HSBAToRGB(c.hue, c.saturation, c.brightness, c.alpha)
		c.dirty = false
	}
	return c.argb
}

func (c *MutableColor) RGBA() int32       { return (c.ARGB() << 8) | (c.ARGB()>>24)&0xFF }
func (c *MutableColor) Transparent() bool { return c.alpha == 0 }
func (c *MutableColor) Red() int32        { return (c.ARGB() >> 16) & 0xFF }
func (c *MutableColor) Green() int32      { return (c.ARGB() >> 8) & 0xFF }
func (c *MutableColor) Blue() int32       { return c.ARGB() & 0xFF }
func (c *MutableColor) AlphaInt() int32   { return (c.ARGB() >> 24) & 0xFF }

func (c *MutableColor) SetRed(value int32) {
	o := RGBToHSB(value, c.Green(), c.Blue())
	c.SetHue(o[0])
	c.SetSaturation(o[1])
	c.SetBrightness(o[2])
	c.dirty = true
}

func (c *MutableColor) SetGreen(value int32) {
	o := RGBToHSB(c.Red(), value, c.Blue())
	c.SetHue(o[0])
	c.SetSaturation(o[1])
	c.SetBrightness(o[2])
	c.dirty = true
}

func (c *MutableColor) SetBlue(value int32) {
	o := RGBToHSB(c.Red(), c.Green(), value)
	c.SetHue(o[0])
	c.SetSaturation(o[1])
	c.SetBrightness(o[2])
	c.dirty = true
}

func (c *MutableColor) SetAlphaInt(value int32) {
	c.dirty = true
	c.alpha = clampf32(float32(value)/255.0, 0, 1)
}

func (c *MutableColor) Recolor(to Color) Mut {
	c.SetHue(to.Hue())
	c.SetSaturation(to.Saturation())
	c.SetBrightness(to.Brightness())
	c.SetAlpha(to.Alpha())
	return c
}

// ChromaColor is a color that animates through hues
type ChromaColor struct {
	*MutableColor
	speedNanos int64
	time       int64
}

// NewChromaColor creates a new chroma color
func NewChromaColor(hue, saturation, brightness, alpha float32, speedNanos int64) *ChromaColor {
	if speedNanos <= 0 {
		panic("speedNanos must be greater than 0!")
	}

	c := &ChromaColor{
		MutableColor: NewMutableColor(hue, saturation, brightness, alpha),
		speedNanos:   speedNanos,
		time:         int64(float32(hue) * float32(speedNanos)),
	}
	return c
}

func (c *ChromaColor) SetSpeedNanos(value int64) {
	if value <= 0 {
		panic("speedNanos must be greater than 0!")
	}
	c.speedNanos = value
}

func (c *ChromaColor) Update(deltaTimeNanos int64) bool {
	c.time += deltaTimeNanos
	c.MutableColor.SetHue(float32(c.time%c.speedNanos) / float32(c.speedNanos))
	return false
}

// AnimatedColor is a color that can be animated
type AnimatedColor struct {
	*MutableColor
	animation interface{} // TODO: Animation dependency
	cdata     []float32
}

// NewAnimatedColor creates a new animated color
func NewAnimatedColor(hue, saturation, brightness, alpha float32) *AnimatedColor {
	return &AnimatedColor{
		MutableColor: NewMutableColor(hue, saturation, brightness, alpha),
	}
}

func (c *AnimatedColor) RecolorWithAnimation(to Color, animation interface{}) Animatable {
	// TODO: Implement animation support. For now, fall back to immediate recolor.
	c.Recolor(to)
	return c
}

func (c *AnimatedColor) Recolor(to Color) Mut {
	c.MutableColor.Recolor(to)
	return c
}

func (c *AnimatedColor) Update(deltaTimeNanos int64) bool {
	// TODO: Animation dependency implementation
	return true
}

// Gradient represents a gradient color
type Gradient struct {
	color1 Color
	color2 Color
	typ    GradientType
}

// GradientType represents the type of gradient
type GradientType interface{}

// NewGradient creates a new gradient
func NewGradient(color1, color2 Color, typ GradientType) *Gradient {
	return &Gradient{
		color1: color1,
		color2: color2,
		typ:    typ,
	}
}

func (g *Gradient) Hue() float32        { return g.color1.Hue() }
func (g *Gradient) Saturation() float32 { return g.color1.Saturation() }
func (g *Gradient) Brightness() float32 { return g.color1.Brightness() }
func (g *Gradient) Alpha() float32      { return g.color1.Alpha() }
func (g *Gradient) ARGB() int32         { return g.color1.ARGB() }
func (g *Gradient) RGBA() int32         { return g.color1.RGBA() }
func (g *Gradient) Transparent() bool   { return g.color1.Transparent() && g.color2.Transparent() }
func (g *Gradient) Red() int32          { return g.color1.Red() }
func (g *Gradient) Green() int32        { return g.color1.Green() }
func (g *Gradient) Blue() int32         { return g.color1.Blue() }
func (g *Gradient) AlphaInt() int32     { return g.color1.AlphaInt() }

func (g *Gradient) Hue2() float32        { return g.color2.Hue() }
func (g *Gradient) Saturation2() float32 { return g.color2.Saturation() }
func (g *Gradient) Brightness2() float32 { return g.color2.Brightness() }
func (g *Gradient) Alpha2() float32      { return g.color2.Alpha() }
func (g *Gradient) ARGB2() int32         { return g.color2.ARGB() }
func (g *Gradient) RGBA2() int32         { return g.color2.RGBA() }
func (g *Gradient) Red2() int32          { return g.color2.Red() }
func (g *Gradient) Green2() int32        { return g.color2.Green() }
func (g *Gradient) Blue2() int32         { return g.color2.Blue() }
func (g *Gradient) AlphaInt2() int32     { return g.color2.AlphaInt() }

func (g *Gradient) Get(index int) Color {
	switch index {
	case 0:
		return g
	case 1:
		return g.color2
	default:
		panic("Invalid index: must be 0 or 1")
	}
}

// Helper functions
func clampf32(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Constants
var (
	// Transparent color
	TRANSPARENT = NewStaticColor(0, 0, 0, 0)
	// White color
	WHITE = NewStaticColor(0, 0, 1, 1)
	// Black color
	BLACK = NewStaticColor(0, 0, 0, 1)
)
