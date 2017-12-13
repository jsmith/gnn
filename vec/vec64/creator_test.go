package vec64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var c Creator

func initCreator() {
	c = Creator{}
}

func TestNumber(t *testing.T) {
	initCreator()
	n := c.Number(1)
	assert.Equal(t, 1., n.(float64))
}

func TestMake(t *testing.T) {
	initCreator()
	v := c.Make(5)
	assert.Equal(t, 5, v.Len())
}

func TestCreatorSub(t *testing.T) {
	initCreator()
	v1 := Vec64{[]float64{0, 1, 2, 3, 4}}
	v2 := Vec64{[]float64{0, 1, 2, 3, 4}}
	vec := c.Sub(v1, v2)

	assert.Equal(t, 0., vec.At(0))
	assert.Equal(t, 0., vec.At(3))
}

func TestFloat64(t *testing.T) {
	initCreator()
	orig := 1.23456789
	n := c.Number(orig)

	after := c.Float64(n)
	assert.InDelta(t, orig, after, 0.00001)
}

func TestCopy(t *testing.T) {
	initCreator()
	v1 := Vec64{[]float64{0, 1, 2, 3, 4}}
	v2 := c.Copy(v1)

	v1.Set(0, 1.)
	assert.Equal(t, 0., v2.At(0))
	assert.Equal(t, 1., v2.At(1))
}
