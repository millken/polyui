package data

type FontWeight int

const (
	WeightThin FontWeight = iota
	WeightExtraLight
	WeightLight
	WeightRegular
	WeightMedium
	WeightSemiBold
	WeightBold
	WeightExtraBold
	WeightBlack
)

type Font struct {
	Resource *Resource
	Family   *FontFamily
	Italic   bool
	Weight   FontWeight
}

func NewFont(path string, family *FontFamily, italic bool, weight FontWeight) *Font {
	return &Font{Resource: NewResource(path, true), Family: family, Italic: italic, Weight: weight}
}
