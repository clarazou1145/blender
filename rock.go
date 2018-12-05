package blender

type Rock struct {
	solid
}

func (*Rock) Name() string {
	return "rock"
}

func NewRock(volume float64) *Rock {
	return &Rock{solid{volume: volume}}
}
