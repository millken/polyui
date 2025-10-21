package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// RGBA creates a color from the given red, green, and blue integer values, and an alpha value 0..1
func RGBA(r, g, b int32, a float32) Color {
	cmps := RGBToHSB(r, g, b)
	return NewStaticColor(cmps[0], cmps[1], cmps[2], a)
}

// ARGB creates a color from an ARGB integer
func ARGB(argb int32) Color {
	cmps := RGBToHSB((argb>>16)&0xFF, (argb>>8)&0xFF, argb&0xFF)
	return NewStaticColor(cmps[0], cmps[1], cmps[2], float32((argb>>24)&0xFF)/255.0)
}

// HSBA creates a color from HSB values and alpha
func HSBA(h, s, b, a float32) Color {
	return NewStaticColor(h, s, b, a)
}

// Hex creates a color from a hex string
func Hex(hex string) (Color, error) {
	return StringToColor(hex, 1.0)
}

// IntToColor takes an ARGB integer color and returns a PolyColor object
func IntToColor(color int32) Color {
	cmps := RGBToHSB((color>>16)&0xFF, (color>>8)&0xFF, color&0xFF)
	return NewStaticColor(cmps[0], cmps[1], cmps[2], float32((color>>24)&0xFF)/255.0)
}

// StringToColor turns a hex string into a color.
//
// - If it is 8 characters long, it is assumed to be in the format #RRGGBBAA (alpha optional)
// - If there is a leading #, it will be removed.
// - If it is 1 character long, the character is repeated e.g. #f -> #ffffff
// - If it is 2 characters long, the character is repeated e.g. #0f -> #0f0f0f
// - If it is 3 characters long, the character is repeated e.g. #0fe -> #00ffee
func StringToColor(hexColor string, alpha float32) (Color, error) {
	hexColor = strings.TrimPrefix(hexColor, "#")

	switch len(hexColor) {
	case 1:
		hexColor = strings.Repeat(hexColor, 6)
	case 2:
		hexColor = strings.Repeat(hexColor, 3)
	case 3:
		var newHex strings.Builder
		for _, c := range hexColor {
			newHex.WriteString(strings.Repeat(string(c), 2))
		}
		hexColor = newHex.String()
	}

	if len(hexColor) < 6 {
		return nil, fmt.Errorf("invalid hex color format: %s", hexColor)
	}

	r, err := strconv.ParseInt(substringOr(hexColor, 0, 2, "0"), 16, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid red component: %v", err)
	}

	g, err := strconv.ParseInt(substringOr(hexColor, 2, 4, "0"), 16, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid green component: %v", err)
	}

	b, err := strconv.ParseInt(substringOr(hexColor, 4, 6, "0"), 16, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid blue component: %v", err)
	}

	a := alpha
	if len(hexColor) == 8 {
		aVal, err := strconv.ParseInt(substringOr(hexColor, 6, 8, "0"), 16, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid alpha component: %v", err)
		}
		a = float32(aVal) / 255.0
	}

	return RGBA(int32(r), int32(g), int32(b), a), nil
}

// ToHex converts a color to a hex string
func ToHex(c Color, includeAlpha, includeHash bool) string {
	format := "%02X%02X%02X"
	if includeAlpha {
		format = "%02X%02X%02X%02X"
	}

	hex := fmt.Sprintf(format, c.Red(), c.Green(), c.Blue(), c.AlphaInt())
	if includeHash {
		hex = "#" + hex
	}

	if !includeAlpha && len(hex) > 7 {
		hex = hex[:7] // Remove alpha if not needed
	}

	return hex
}

// AsAnimatable converts a color to an animatable color
// AsAnimatable converts a color to an animatable color
func AsAnimatable(c Color) Animatable {
	if anim, ok := c.(Animatable); ok {
		return anim
	}
	return NewAnimatedColor(c.Hue(), c.Saturation(), c.Brightness(), c.Alpha())
}

// AsAnimatableGradient converts a gradient to an animatable gradient
func AsAnimatableGradient(g *Gradient) interface{} { // TODO: proper return type
	// TODO: implement gradient animation
	return g
}

// AsMutable converts a color to a mutable color
// AsMutable converts a color to a mutable color
func AsMutable(c Color) Mut {
	if mut, ok := c.(Mut); ok {
		return mut
	}
	return NewMutableColor(c.Hue(), c.Saturation(), c.Brightness(), c.Alpha())
}

// HSBAToRGB converts HSB components to equivalent RGB values.
// The saturation and brightness components should be floating-point values between zero and one.
// The hue component can be any floating-point number.
func HSBAToRGB(hue, saturation, brightness, alpha float32) int32 {
	var r, g, b int32

	if saturation == 0 {
		r = int32(brightness*255.0 + 0.5)
		g = r
		b = r
	} else {
		h := (hue - float32(math.Floor(float64(hue)))) * 6.0
		f := h - float32(math.Floor(float64(h)))
		p := brightness * (1.0 - saturation)
		q := brightness * (1.0 - saturation*f)
		t := brightness * (1.0 - saturation*(1.0-f))

		switch int(h) {
		case 0:
			r = int32(brightness*255.0 + 0.5)
			g = int32(t*255.0 + 0.5)
			b = int32(p*255.0 + 0.5)
		case 1:
			r = int32(q*255.0 + 0.5)
			g = int32(brightness*255.0 + 0.5)
			b = int32(p*255.0 + 0.5)
		case 2:
			r = int32(p*255.0 + 0.5)
			g = int32(brightness*255.0 + 0.5)
			b = int32(t*255.0 + 0.5)
		case 3:
			r = int32(p*255.0 + 0.5)
			g = int32(q*255.0 + 0.5)
			b = int32(brightness*255.0 + 0.5)
		case 4:
			r = int32(t*255.0 + 0.5)
			g = int32(p*255.0 + 0.5)
			b = int32(brightness*255.0 + 0.5)
		case 5:
			r = int32(brightness*255.0 + 0.5)
			g = int32(p*255.0 + 0.5)
			b = int32(q*255.0 + 0.5)
		default:
			r = 0
			g = 0
			b = 0
		}
	}

	return (int32(alpha*255.0) << 24) | (r << 16) | (g << 8) | b
}

// RGBToHSB converts RGB components to equivalent HSB values.
func RGBToHSB(r, g, b int32) [3]float32 {
	var hue, saturation, brightness float32

	cmax := r
	if g > cmax {
		cmax = g
	}
	if b > cmax {
		cmax = b
	}

	cmin := r
	if g < cmin {
		cmin = g
	}
	if b < cmin {
		cmin = b
	}

	diff := float32(cmax - cmin)
	brightness = float32(cmax) / 255.0

	if cmax != 0 {
		saturation = diff / float32(cmax)
	} else {
		saturation = 0
	}

	if saturation == 0 {
		hue = 0
	} else {
		redc := float32(cmax-r) / diff
		greenc := float32(cmax-g) / diff
		bluec := float32(cmax-b) / diff

		if r == cmax {
			hue = bluec - greenc
		} else if g == cmax {
			hue = 2.0 + redc - bluec
		} else {
			hue = 4.0 + greenc - redc
		}

		hue /= 6.0
		if hue < 0 {
			hue += 1.0
		}
	}

	return [3]float32{hue, saturation, brightness}
}

// Helper function for substring extraction
func substringOr(s string, start, end int, defaultValue string) string {
	if start >= len(s) {
		return defaultValue
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return defaultValue
	}
	return s[start:end]
}
