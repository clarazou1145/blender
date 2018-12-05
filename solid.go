package blender

type solid struct {
	Substance
	volume float64
}

func (*WaterMelonJuice) Name() string {
	return "water melon juice"
}

func (s *solid) IsOverflow(capacity float64) bool {
	return s.volume > capacity
}

func NewWaterMelonJuice(liter float64) *WaterMelonJuice {
	w := &WaterMelonJuice{liquid: liquid{liter, nil}}
	w.liquefied = w
	return w
}
