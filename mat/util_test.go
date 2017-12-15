package mat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	initMatrix()

	assert.Equal(t, 2, m.ColCount())
	assert.Equal(t, 4, m.Col(0).Len())
	assert.Equal(t, 2., m.Col(0).At(2))
}

func TestCopy(t *testing.T) {
	initMatrix()
	c := Copy(m)
	m.Set(0, 0, 500)
	assert.Equal(t, 2, c.ColCount())
	assert.Equal(t, 4, c.Col(0).Len())
	assert.Equal(t, 0., c.At(0, 0))
}
