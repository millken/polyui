package color

// LightTheme is the default light color set used in PolyUI.
type LightTheme struct {
	name      string
	page      *Page
	brand     *Brand
	onBrand   *OnBrand
	state     *State
	component *Component
	text      *Text
}

// NewLightTheme creates a new light theme with the default colors
func NewLightTheme() *LightTheme {
	return &LightTheme{
		name: "light",
		page: &Page{
			Bg: &Palette{
				Normal:   RGBA(232, 237, 255, 1.0), // #E8EDFF
				Hovered:  RGBA(222, 228, 252, 1.0), // #DEE4FC
				Pressed:  RGBA(239, 243, 255, 1.0), // #EFF3FF
				Disabled: RGBA(232, 237, 255, 0.5), // #E8EDFF80
			},
			BgOverlay: RGBA(0, 0, 0, 0.25), // #00000040
			Fg: &Palette{
				Normal:   RGBA(17, 23, 28, 1.0), // #11171C
				Hovered:  RGBA(26, 34, 41, 1.0), // #1A2229
				Pressed:  RGBA(14, 19, 23, 1.0), // #0E1317
				Disabled: RGBA(17, 23, 28, 0.5), // #11171C80
			},
			FgOverlay: RGBA(255, 255, 255, 0.1), // #FFFFFF1A
			Border20:  RGBA(0, 0, 0, 0.2),
			Border10:  RGBA(0, 0, 0, 0.1),
			Border5:   RGBA(0, 0, 0, 0.05),
		},
		brand: &Brand{
			Fg: &Palette{
				Normal:   RGBA(64, 93, 255, 1.0), // #405DFF
				Hovered:  RGBA(40, 67, 221, 1.0), // #2843DD
				Pressed:  RGBA(57, 87, 255, 1.0), // #3957FF
				Disabled: RGBA(64, 93, 255, 0.5), // #405DFF80
			},
			Accent: &Palette{
				Normal:   RGBA(223, 236, 253, 1.0), // #DFECFD
				Hovered:  RGBA(183, 208, 251, 1.0), // #B7D0FB
				Pressed:  RGBA(177, 206, 255, 1.0), // #B1CEFF
				Disabled: RGBA(15, 28, 51, 0.5),    // #0F1C3380
			},
		},
		onBrand: &OnBrand{
			Fg: &Palette{
				Normal:   RGBA(213, 219, 255, 1.0), // #D5DBFF
				Hovered:  RGBA(215, 220, 251, 1.0), // #D7DCFB
				Pressed:  RGBA(225, 229, 255, 1.0), // #E1E5FF
				Disabled: RGBA(213, 219, 255, 0.5), // #D5DBFF80
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
				Hovered:  RGBA(26, 135, 82, 1.0),  // #1A8752
				Pressed:  RGBA(44, 172, 110, 1.0), // #2CAC6E
				Disabled: RGBA(35, 154, 96, 0.5),  // #239A6080
			},
		},
		component: &Component{
			Bg: &Palette{
				Normal:   RGBA(208, 215, 243, 1.0), // #D0D7F3
				Hovered:  RGBA(213, 219, 243, 1.0), // #D5DBF3
				Pressed:  RGBA(238, 241, 255, 1.0), // #EEF1FF
				Disabled: RGBA(222, 228, 252, 0.5), // #DEE4FC80
			},
			BgDeselected: TRANSPARENT,
		},
		text: &Text{
			Primary: &Palette{
				Normal:   RGBA(2, 3, 7, 1.0),    // #020307
				Hovered:  RGBA(11, 15, 33, 1.0), // #0B0F21
				Pressed:  RGBA(2, 5, 15, 1.0),   // #02050F
				Disabled: RGBA(2, 3, 7, 0.5),    // #02030780
			},
			Secondary: &Palette{
				Normal:   RGBA(117, 120, 131, 1.0), // #757883
				Hovered:  RGBA(101, 104, 116, 1.0), // #656874
				Pressed:  RGBA(136, 139, 150, 1.0), // #888B96
				Disabled: RGBA(117, 120, 131, 0.5), // #75788380
			},
		},
	}
}

// Colors interface implementation
func (t *LightTheme) Name() string          { return t.name }
func (t *LightTheme) Page() *Page           { return t.page }
func (t *LightTheme) Brand() *Brand         { return t.brand }
func (t *LightTheme) OnBrand() *OnBrand     { return t.onBrand }
func (t *LightTheme) State() *State         { return t.state }
func (t *LightTheme) Component() *Component { return t.component }
func (t *LightTheme) Text() *Text           { return t.text }
