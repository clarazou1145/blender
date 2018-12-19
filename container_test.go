package blender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOverflow(t *testing.T) {
	overflowWater := NewWater(capacity + 1)
	container := NewContainer(capacity)
	err := container.PutIn(overflowWater)
	assert.NotNil(t, err)
}
