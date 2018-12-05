package blender

type LiquidPhase interface {
	Substance
	Liquefy() LiquidPhase
	Liter() float64
}
