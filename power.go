package blender

type PowerSource interface {
	Voltage() int
	Ampere() float32
	Shape() string
}
