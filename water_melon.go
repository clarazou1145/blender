package blender

type WaterMelon struct {
	solid
}

func (w *WaterMelon) Liquefy() LiquidPhase {
	return NewWaterMelonJuice(w.volume * 0.8)
}

func NewWaterMelon(volume float64) *WaterMelon {
	w := &WaterMelon{solid{volume: volume}}
	return w
}

func (*WaterMelon) Name() string {
	return "water melon"
}
