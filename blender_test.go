package blender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const capacity = 10.0

var sut = &Blender{capacity}

func TestBlendFruit(t *testing.T) {
	output, outputLiter, err := sut.Blend("water melon", 1.0)
	assert.Equal(t, "water melon juice", output)
	assert.Equal(t, 0.8, outputLiter)
	assert.Nil(t, err)
}

func TestBlendWater(t *testing.T) {
	output, outputLiter, err := sut.Blend("water", 1.0)
	assert.Equal(t, "water", output)
	assert.Equal(t, 1.0, outputLiter)
	assert.Nil(t, err)
}

func TestBlendRock(t *testing.T) {
	_, _, err := sut.Blend("rock", 1.0)
	assert.NotNil(t, err)
}

func TestOverflow(t *testing.T) {
	overflowWater := NewWater(capacity + 1)
	_, err := sut.BlendSubstance(overflowWater)
	assert.NotNil(t, err)
}
