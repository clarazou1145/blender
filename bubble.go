package blender

type Bubble interface {
	TurnOn()
	TurnOff()
	SetColor(color string) error
}
