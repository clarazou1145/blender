package blender

type Substance interface {
	Name() string
	IsOverflow(capacity float64) bool
}
