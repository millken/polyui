package data

type Size struct{ X, Y float32 }

type ImageType int

const (
	ImageRaster ImageType = iota
	ImageVector
	ImageUnknown
)

type PolyImage struct {
	Resource *Resource
	Type     ImageType
	Size     Size
}

func NewPolyImage(path string, t ImageType) *PolyImage {
	return &PolyImage{Resource: NewResource(path, true), Type: t}
}
