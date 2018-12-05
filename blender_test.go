package blender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlendFruit(t *testing.T) {
	output, outputLiter, err := Blend("water melon", 1.0)
	assert.Equal(t, "water melon juice", output)
	assert.Equal(t, 0.8, outputLiter)
	assert.Nil(t, err)
}

func TestBlendWater(t *testing.T) {
	output, outputLiter, err := Blend("water", 1.0)
	assert.Equal(t, "water", output)
	assert.Equal(t, 1.0, outputLiter)
	assert.Nil(t, err)
}

func TestBlendRock(t *testing.T) {
	_, _, err := Blend("rock", 1.0)
	assert.NotNil(t, err)
}
