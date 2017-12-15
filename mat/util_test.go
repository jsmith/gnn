package mat

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	mat := Make(2, 5)

	assert.Equal(t, 2, mat.RowCount())
	assert.Equal(t, 5, mat.ColCount())
}

func TestInitCols(t *testing.T) {
	mat := InitCols(
		vec.Init(0, 1, 2, 3),
		vec.Init(3, 2, 1, 0),
	)

	assert.Equal(t, 4, mat.RowCount())
	assert.Equal(t, 2., mat.At(2, 0))
}

func TestInitRows(t *testing.T) {
	mat := InitRows(
		vec.Init(0, 1, 2, 3),
		vec.Init(3, 2, 1, 0),
	)

	assert.Equal(t, 2, mat.RowCount())
	assert.Equal(t, 2, mat.Col(0).Len())
	assert.Equal(t, 2., mat.At(0, 2))
}

func TestCopy(t *testing.T) {
	initMatrix()
	c := Copy(m)
	m.Set(0, 0, 500)
	assert.Equal(t, 4, c.ColCount())
	assert.Equal(t, 2, c.RowCount())
	assert.Equal(t, 0., c.At(0, 0))
}

func TestMul(t *testing.T) {
	m1 := InitRows(
		vec.Init(0, 1, 1),
		vec.Init(1, 0, 1),
	)

	m2 := InitRows(
		vec.Init(2, 2),
		vec.Init(2, 2),
		vec.Init(2, 2),
	)

	m := Mul(m1, m2)
	assert.Equal(t, 2, m.RowCount())
	assert.Equal(t, 2, m.ColCount())
	assert.Equal(t, 4., m.At(0, 0))
	assert.Equal(t, 4., m.At(1, 0))
}
