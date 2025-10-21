package color

// DarkTheme is the default color set used in PolyUI.
type DarkTheme struct {
	name      string
	page      *Page
	brand     *Brand
	onBrand   *OnBrand
	state     *State
	component *Component
	text      *Text
}

// NewDarkTheme creates a new dark theme with the default colors
func NewDarkTheme() *DarkTheme {
	return &DarkTheme{
		name: "dark",
		page: &Page{
			Bg: &Palette{
				Normal:   RGBA(17, 23, 28, 1.0), // #11171C
				Hovered:  RGBA(21, 28, 34, 1.0), // #151C22
				Pressed:  RGBA(14, 19, 23, 1.0), // #0E1317
				Disabled: RGBA(17, 23, 28, 0.5), // #11171C80
			},
			BgOverlay: RGBA(255, 255, 255, 0.1), // #FFFFFF1A
			Fg: &Palette{
				Normal:   RGBA(17, 23, 28, 1.0), // #11171C
				Hovered:  RGBA(26, 34, 41, 1.0), // #1A2229
				Pressed:  RGBA(14, 19, 23, 1.0), // #0E1317
				Disabled: RGBA(26, 34, 41, 0.5), // #1A222980
			},
			FgOverlay: RGBA(255, 255, 255, 0.1), // #FFFFFF1A
			Border20:  RGBA(255, 255, 255, 0.2),
			Border10:  RGBA(255, 255, 255, 0.1),
			Border5:   RGBA(255, 255, 255, 0.05),
		},
		brand: &Brand{
			Fg: &Palette{
				Normal:   RGBA(43, 75, 255, 1.0), // #2B4BFF
				Hovered:  RGBA(40, 67, 221, 1.0), // #2843DD
				Pressed:  RGBA(57, 87, 255, 1.0), // #3957FF
				Disabled: RGBA(57, 87, 255, 0.5), // #3957FF80
			},
			Accent: &Palette{
				Normal:   RGBA(15, 28, 51, 1.0), // #0F1C33
				Hovered:  RGBA(12, 23, 41, 1.0), // #0C1729
				Pressed:  RGBA(26, 44, 78, 1.0), // #1A2C4E
				Disabled: RGBA(15, 28, 51, 0.5), // #0F1C3380
			},
		},
		onBrand: &OnBrand{
			Fg: &Palette{
				Normal:   RGBA(213, 219, 255, 1.0),  // #D5DBFF
				Hovered:  RGBA(213, 219, 255, 0.85), // #D5DBFFDA
				Pressed:  RGBA(225, 229, 255, 1.0),  // #E1E5FF
				Disabled: RGBA(225, 229, 255, 0.5),  // #E1E5FF80
			},
			Accent: &Palette{
				Normal:   RGBA(63, 124, 228, 1.0),  // #3F7CE4
				Hovered:  RGBA(63, 124, 228, 0.85), // #3F7CE4DA
				Pressed:  RGBA(37, 80, 154, 1.0),   // #25509A
				Disabled: RGBA(63, 124, 228, 0.5),  // #3F7CE480
			},
		},
		state: &State{
			Danger: &Palette{
				Normal:   RGBA(255, 68, 68, 1.0), // #FF4444
				Hovered:  RGBA(214, 52, 52, 1.0), // #D63434
				Pressed:  RGBA(255, 86, 86, 1.0), // #FF5656
				Disabled: RGBA(255, 68, 68, 0.5), // #FF444480
			},
			Warning: &Palette{
				Normal:   RGBA(255, 171, 29, 1.0), // #FFAB1D
				Hovered:  RGBA(233, 156, 27, 1.0), // #E99C1B
				Pressed:  RGBA(255, 178, 49, 1.0), // #FFB231
				Disabled: RGBA(255, 171, 29, 0.5), // #FFAB1D80
			},
			Success: &Palette{
				Normal:   RGBA(35, 154, 96, 1.0),  // #239A60
				Hovered:  RGBA(26, 139, 82, 1.0),  // #1A8B52
				Pressed:  RGBA(44, 172, 110, 1.0), // #2CAC6E
				Disabled: RGBA(35, 154, 96, 0.5),  // #239A6080
			},
		},
		component: &Component{
			Bg: &Palette{
				Normal:   RGBA(26, 34, 41, 1.0),  // #1A2229
				Hovered:  RGBA(23, 31, 37, 0.85), // #171F25DA
				Pressed:  RGBA(34, 44, 53, 1.0),  // #222C35
				Disabled: RGBA(34, 44, 53, 0.5),  // #222C3580
			},
			BgDeselected: TRANSPARENT,
		},
		text: &Text{
			Primary: &Palette{
				Normal:   RGBA(223, 229, 255, 1.0),  // #D5DBFF
				Hovered:  RGBA(223, 229, 255, 0.85), // #D5DBFFDA
				Pressed:  RGBA(235, 239, 255, 1.0),  // #E1E5FF
				Disabled: RGBA(235, 239, 255, 0.5),  // #E1E5FF80
			},
			Secondary: &Palette{
				Normal:   RGBA(120, 119, 141, 1.0), // #78778D
				Hovered:  RGBA(95, 104, 116, 1.0),  // #5F6874
				Pressed:  RGBA(130, 141, 155, 1.0), // #828D9B
				Disabled: RGBA(120, 129, 141, 0.5), // #78818D80
			},
		},
	}
}

// Colors interface implementation
func (t *DarkTheme) Name() string          { return t.name }
func (t *DarkTheme) Page() *Page           { return t.page }
func (t *DarkTheme) Brand() *Brand         { return t.brand }
func (t *DarkTheme) OnBrand() *OnBrand     { return t.onBrand }
func (t *DarkTheme) State() *State         { return t.state }
func (t *DarkTheme) Component() *Component { return t.component }
func (t *DarkTheme) Text() *Text           { return t.text }
