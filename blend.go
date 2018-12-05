package blender

import (
	"errors"
	"fmt"
)

func Blend(substance string, inputLiter float64) (output string, outputLiter float64, err error) {
	s, err := createSubstance(substance, inputLiter)
	if err != nil {
		return
	}
	switch s.(type) {
	case Liquable:
		{
			liquid := (s.(Liquable)).Liquefy()
			output = liquid.Name()
			outputLiter = liquid.Liter()
		}
	default:
		err = errors.New(fmt.Sprintf("cannot blend %s", substance))
	}
	return
}

func createSubstance(substance string, inputLiter float64) (Substance, error) {
	switch substance {
	case "water melon":
		return NewWaterMelon(inputLiter), nil
	case "water":
		return NewWater(inputLiter), nil
	case "water melon juice":
		return NewWaterMelonJuice(inputLiter), nil
	case "rock":
		return NewRock(inputLiter), nil
	default:
		return nil, errors.New("unrecognized substance")
	}
}

type Substance interface {
	Name() string
}

type solid struct {
	volume float64
}

type LiquidPhase interface {
	Substance
	Liquefy() LiquidPhase
	Liter() float64
}

type liquid struct {
	//abstract class
	liter     float64
	liquefied LiquidPhase
}

type Liquable interface {
	Liquefy() LiquidPhase
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

type WaterMelonJuice struct {
	liquid
}

func (*WaterMelonJuice) Name() string {
	return "water melon juice"
}

func NewWaterMelonJuice(liter float64) *WaterMelonJuice {
	w := &WaterMelonJuice{liquid: liquid{liter, nil}}
	w.liquefied = w
	return w
}

type Rock struct {
	solid
}

func (*Rock) Name() string {
	return "rock"
}

func NewRock(volume float64) *Rock {
	return &Rock{solid{volume: volume}}
}

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
