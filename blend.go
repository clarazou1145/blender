package blender

import (
	"errors"
	"fmt"
)

type Blender struct {
	Capacity float64
}

func (b *Blender) Blend(substance string, inputLiter float64) (output string, outputLiter float64, err error) {
	s, err := createSubstance(substance, inputLiter)
	if err != nil {
		return
	}

	liquid, err := b.BlendSubstance(s)
	if err != nil {
		return
	}
	output = liquid.Name()
	outputLiter = liquid.Liter()
	return
}

func (b *Blender) BlendSubstance(s Substance) (LiquidPhase, error) {
	if s.IsOverflow(b.Capacity) {
		return nil, errors.New("overflow")
	}
	switch s.(type) {
	case Liquable:
		{
			liquid := (s.(Liquable)).Liquefy()
			return liquid, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("cannot blend %s", s))
	}
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
