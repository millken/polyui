package data

type FontFamily struct {
	Name string
	Path string
}

func NewFontFamily(name, path string) *FontFamily { return &FontFamily{Name: name, Path: path} }
