package input

// Input package: Keys, Mouse and Modifiers ported from Kotlin input.kt

import "fmt"

// Keys are mapped non-printable or special keys. Use int for value to
// interoperate with existing code.
type Key int

const (
	KeyUnknown Key = -1
	KeyF1      Key = 1
	KeyF2      Key = 2
	KeyF3      Key = 3
	KeyF4      Key = 4
	KeyF5      Key = 5
	KeyF6      Key = 6
	KeyF7      Key = 7
	KeyF8      Key = 8
	KeyF9      Key = 9
	KeyF10     Key = 10
	KeyF11     Key = 11
	KeyF12     Key = 12

	KeyEscape Key = 20

	KeyEnter     Key = 21
	KeyTab       Key = 22
	KeyBackspace Key = 23
	KeyInsert    Key = 24
	KeyDelete    Key = 25
	KeyPageUp    Key = 26
	KeyPageDown  Key = 27
	KeyHome      Key = 28
	KeyEnd       Key = 29

	KeyRight Key = 30
	KeyLeft  Key = 31
	KeyDown  Key = 32
	KeyUp    Key = 33
)

// Mouse buttons
type MouseButton int

const (
	MouseUnknown MouseButton = -1
	MouseLeft    MouseButton = 0
	MouseRight   MouseButton = 1
	MouseMiddle  MouseButton = 2
)

// Modifiers is a bitmask for key modifiers (byte-sized to match kotlin)
type Modifiers byte

const (
	ModLShift Modifiers = 1 << iota
	ModRShift
	ModLPrimary
	ModRPrimary
	ModLSecondary
	ModRSecondary
	ModLMeta
	ModRMeta
)

// Combined helpers
const (
	ModShift     Modifiers = ModLShift | ModRShift
	ModPrimary   Modifiers = ModLPrimary | ModRPrimary
	ModSecondary Modifiers = ModLSecondary | ModRSecondary
	ModMeta      Modifiers = ModLMeta | ModRMeta
)

func (m Modifiers) IsEmpty() bool    { return m == 0 }
func (m Modifiers) HasShift() bool   { return m&ModShift != 0 }
func (m Modifiers) HasPrimary() bool { return m&ModPrimary != 0 }
func (m Modifiers) HasAlt() bool     { return m&ModSecondary != 0 }
func (m Modifiers) HasMeta() bool    { return m&ModMeta != 0 }

func (m Modifiers) String() string {
	if m == 0 {
		return ""
	}
	out := ""
	if m&ModLShift != 0 {
		out += "LSHIFT+"
	}
	if m&ModRShift != 0 {
		out += "RSHIFT+"
	}
	if m&ModLPrimary != 0 {
		out += "LPRIMARY+"
	}
	if m&ModRPrimary != 0 {
		out += "RPRIMARY+"
	}
	if m&ModLSecondary != 0 {
		out += "LSECONDARY+"
	}
	if m&ModRSecondary != 0 {
		out += "RSECONDARY+"
	}
	if m&ModLMeta != 0 {
		out += "LMETA+"
	}
	if m&ModRMeta != 0 {
		out += "RMETA+"
	}
	if len(out) > 0 {
		out = out[:len(out)-1]
	}
	return out
}

// Helper to format key+modifiers like Kotlin's toString/toStringPretty
func ToStringKey(k Key, m Modifiers) string {
	if m.IsEmpty() {
		return fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("%s+%v", m.String(), k)
}

func ToStringMouse(b MouseButton, m Modifiers) string {
	if m.IsEmpty() {
		return fmt.Sprintf("%v", b)
	}
	return fmt.Sprintf("%s+%v", m.String(), b)
}
