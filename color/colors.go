package color

// TODO: PolyUI dependency - need to implement these constants
// import "github.com/polyfrost/polyui"

// TODO: PolyUI dependency - these constants are defined in PolyUI package
// const (
// 	INPUT_NONE     = 0
// 	INPUT_HOVERED  = 1
// 	INPUT_PRESSED  = 2
// )

// Colors is the color storage for PolyUI.
// All PolyUI drawables use these colors.
// or even changed on the fly using the field in the PolyUI instance, bringing theming to PolyUI.
type Colors interface {
	Name() string
	Page() *Page
	Brand() *Brand
	OnBrand() *OnBrand
	State() *State
	Component() *Component
	Text() *Text
}

// Page represents the color palette for page backgrounds and foregrounds
type Page struct {
	Bg        *Palette
	BgOverlay Color
	Fg        *Palette
	FgOverlay Color
	Border20  Color
	Border10  Color
	Border5   Color
}

// Brand represents the brand colors
type Brand struct {
	Fg     *Palette
	Accent *Palette
}

// OnBrand represents colors used on brand backgrounds
type OnBrand struct {
	Fg     *Palette
	Accent *Palette
}

// State represents state colors (danger, warning, success)
type State struct {
	Danger  *Palette
	Warning *Palette
	Success *Palette
}

// Component represents component colors
type Component struct {
	Bg           *Palette
	BgDeselected Color
}

// Text represents text colors
type Text struct {
	Primary   *Palette
	Secondary *Palette
}

// Palette represents a set of four colors representing the four key states that a component can have.
type Palette struct {
	Normal   Color
	Hovered  Color
	Pressed  Color
	Disabled Color
}

// Get returns the color for the given state
func (p *Palette) Get(state byte) Color {
	// TODO: PolyUI dependency - need to implement these constants
	// switch state {
	// case INPUT_NONE:
	// 	return p.Normal
	// case INPUT_HOVERED:
	// 	return p.Hovered
	// case INPUT_PRESSED:
	// 	return p.Pressed
	// default:
	// 	return p.Normal
	// }
	return p.Normal // fallback
}

// GetNewPalette calculates and returns the new Palette for the given currentPalette based on the current and new Colors.
// This is useful for theming, where you want to change the palette of a component based on the current theme.
// If the currentPalette is not found in the current Colors, it will return nil.
func GetNewPalette(currentPalette *Palette, current Colors, newColors Colors) *Palette {
	if currentPalette == nil {
		return nil
	}

	currentPage := current.Page()
	newPage := newColors.Page()

	if currentPalette == currentPage.Bg {
		return newPage.Bg
	} else if currentPalette == currentPage.Fg {
		return newPage.Fg
	}

	currentBrand := current.Brand()
	newBrand := newColors.Brand()
	if currentPalette == currentBrand.Fg {
		return newBrand.Fg
	} else if currentPalette == currentBrand.Accent {
		return newBrand.Accent
	}

	currentOnBrand := current.OnBrand()
	newOnBrand := newColors.OnBrand()
	if currentPalette == currentOnBrand.Fg {
		return newOnBrand.Fg
	} else if currentPalette == currentOnBrand.Accent {
		return newOnBrand.Accent
	}

	currentState := current.State()
	newState := newColors.State()
	if currentPalette == currentState.Danger {
		return newState.Danger
	} else if currentPalette == currentState.Warning {
		return newState.Warning
	} else if currentPalette == currentState.Success {
		return newState.Success
	}

	currentComponent := current.Component()
	newComponent := newColors.Component()
	if currentPalette == currentComponent.Bg {
		return newComponent.Bg
	}

	currentText := current.Text()
	newText := newColors.Text()
	if currentPalette == currentText.Primary {
		return newText.Primary
	} else if currentPalette == currentText.Secondary {
		return newText.Secondary
	}

	return nil
}

// GetNewColor calculates and returns the new Color for the given currentColor based on the current and new Colors.
// This is useful for theming, where you want to change the color of a component based on the current theme.
// If the currentColor is not found in the current Colors, it will return nil.
func GetNewColor(currentColor Color, current Colors, newColors Colors) Color {
	currentPage := current.Page()
	newPage := newColors.Page()

	if currentColor == currentPage.BgOverlay {
		return newPage.BgOverlay
	} else if currentColor == currentPage.FgOverlay {
		return newPage.FgOverlay
	} else if currentColor == currentPage.Border20 {
		return newPage.Border20
	} else if currentColor == currentPage.Border10 {
		return newPage.Border10
	} else if currentColor == currentPage.Border5 {
		return newPage.Border5
	}

	currentComponent := current.Component()
	newComponent := newColors.Component()
	if currentColor == currentComponent.BgDeselected {
		return newComponent.BgDeselected
	}

	return TRANSPARENT
}
