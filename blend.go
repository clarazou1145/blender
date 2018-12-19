package blender

import (
	"errors"
	"fmt"
)

type Blender struct {
	Container   Container
	Bubble      Bubble
	PowerSource PowerSource
	Voltage     int
	Ampere      float32
	PlugShape   string
}

func (b *Blender) SetContainer(c Container) {
	b.Container = c
}

func (b *Blender) SetBubble(bubble Bubble) {
	b.Bubble = bubble
}

func (b *Blender) Blend() error {
	s := b.Container.PourOut()
	var liquid LiquidPhase
	switch s.(type) {
	case Liquable:
		{
			liquid = (s.(Liquable)).Liquefy()
		}
	default:
		return errors.New(fmt.Sprintf("cannot blend %s", s))
	}
	return b.Container.PutIn(liquid)
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
