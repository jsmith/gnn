package mat

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	v1 := vec.Init(0, 1, 2, 3)
	v2 := vec.Init(0, 1, 2, 3)
	m := Init(v1, v2)

	assert.Equal(t, 2, m.ColCount())
	assert.Equal(t, 4, m.Col(0).Len())
	assert.Equal(t, 2., m.Col(0).At(2))
}
