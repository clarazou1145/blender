package blender

type Water struct {
	liquid
}

func (*Water) Name() string {
	return "water"
}

func NewWater(liter float64) *Water {
	w := &Water{liquid: liquid{liter, nil}}
	w.liquefied = w
	return w
}
