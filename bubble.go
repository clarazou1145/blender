package blender

import "errors"

type Bubble interface {
	TurnOn()
	TurnOff()
	SetColor(color string) error
}

type bubble struct {
	isOn bool
}

func (b *bubble) TurnOn() {
	b.isOn = true
}

func (b *bubble) TurnOff() {
	b.isOn = false
}

func (b *bubble) SetColor(color string) error {
	if !b.isOn {
		return errors.New("Bubble Is TurnOff")
	}
	return nil
}
