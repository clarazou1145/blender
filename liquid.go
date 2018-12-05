package blender

type liquid struct {
	//abstract class
	liter     float64
	liquefied LiquidPhase
}

func (*liquid) Name() string {
	panic("implement me") //abstract class cannot be created
}

func (l *liquid) Liquefy() LiquidPhase {
	return l.liquefied
}

func (l *liquid) Liter() float64 {
	return l.liter
}

func (l *liquid) IsOverflow(capacity float64) bool {
	return l.liter > capacity
}
