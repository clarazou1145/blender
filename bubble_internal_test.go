package blender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetColorAfterTurnOn(t *testing.T) {
	bubbleSut := &bubble{}
	bubbleSut.TurnOn()
	err := bubbleSut.SetColor("red")
	assert.Nil(t, err)
}

func TestSetColorBeforeTurnOn(t *testing.T) {
	bubbleSut := &bubble{}
	err := bubbleSut.SetColor("red")
	assert.Equal(t, "Bubble Is TurnOff", err.Error())
}
