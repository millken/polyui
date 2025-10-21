package unit

// Align and related enums ported from Kotlin align.kt

// Content describes main/cross alignment behavior.
type Content int

const (
	ContentStart Content = iota
	ContentCenter
	ContentEnd
	ContentSpaceBetween
	ContentSpaceEvenly
)

// Line describes vertical line alignment.
type Line int

const (
	LineStart Line = iota
	LineCenter
	LineEnd
)

// Mode is the layout direction.
type Mode int

const (
	ModeHorizontal Mode = iota
	ModeVertical
)

// Wrap mode for autolayout.
type Wrap int

const (
	WrapNever Wrap = iota
	WrapAlways
	WrapAuto
)

// SpawnPos maps the Kotlin SpawnPos enum.
type SpawnPos int

const (
	SpawnPosAboveMouse SpawnPos = iota
	SpawnPosAtMouse
	SpawnPosBelowMouse
)

// Align defines layout alignment and paddings.
type Align struct {
	Main       Content
	Cross      Content
	Line       Line
	Mode       Mode
	PadBetween Vec2
	PadEdges   Vec2
	Wrap       Wrap
}

// NewAlign creates an Align with explicit padBetween and padEdges.
func NewAlign(main Content, cross Content, line Line, mode Mode, padBetween, padEdges Vec2, wrap Wrap) Align {
	return Align{Main: main, Cross: cross, Line: line, Mode: mode, PadBetween: padBetween, PadEdges: padEdges, Wrap: wrap}
}

// NewAlignUniformPad creates Align using the same pad for between and edges.
func NewAlignUniformPad(main Content, cross Content, line Line, mode Mode, pad Vec2, wrap Wrap) Align {
	return NewAlign(main, cross, line, mode, pad, pad, wrap)
}

// Convenience constructor with px,py for both padBetween and padEdges
func NewAlignPX(main Content, cross Content, line Line, mode Mode, px, py float32) Align {
	p := NewVec2(px, py)
	return NewAlign(main, cross, line, mode, p, p, WrapAuto)
}

// AlignDefault is the default Align value
var AlignDefault = Align{Main: ContentStart, Cross: ContentStart, Line: LineCenter, Mode: ModeHorizontal, PadBetween: NewVec2(6, 6), PadEdges: NewVec2(6, 6), Wrap: WrapAuto}
