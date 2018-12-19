package blender

import "errors"

type container struct {
	capacity  float64
	substance Substance
}

func (c *container) Capacity() float64 {
	return c.capacity
}

func (c *container) PutIn(s Substance) (err error) {
	if s.IsOverflow(c.capacity) {
		err = errors.New("overflow")
	} else {
		c.substance = s
	}
	return
}

func (c *container) PourOut() Substance {
	s := c.substance
	c.substance = nil
	return s
}

func NewContainer(capacity float64) Container {
	return &container{
		capacity: capacity,
	}
}

//go:generate mockgen -package=blender -destination=./mock_container.go github.com/lonegunmanb/blender Container
type Container interface {
	Capacity() float64
	PutIn(s Substance) error
	PourOut() Substance
}
