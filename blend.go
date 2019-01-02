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

func NewBlender(bubble Bubble) *Blender {
	if bubble == nil {
		panic("bubble is nil")
	}
	return &Blender{Bubble: bubble}
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
			l, err := b.liquefy(s.(Liquable))
			if err != nil {
				return err
			}
			liquid = l
		}
	default:
		return errors.New(fmt.Sprintf("cannot blend %s", s))
	}
	return b.Container.PutIn(liquid)
}

func (b *Blender) liquefy(liquable Liquable) (LiquidPhase, error) {
	var liquid LiquidPhase
	err := operateBubbleAroundAction(b.Bubble, func() {
		liquid = liquable.Liquefy()
	})
	if err != nil {
		return nil, err
	}
	return liquid, nil
}

func operateBubbleAroundAction(bubble Bubble, action func()) error {
	bubble.TurnOn()
	err := bubble.SetColor("red")
	if err != nil {
		return err
	}
	action()
	err = bubble.SetColor("green")
	if err != nil {
		return err
	}
	bubble.TurnOff()
	return nil
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
