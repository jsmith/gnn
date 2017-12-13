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
