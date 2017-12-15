package mat

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var m Matrix

func initMatrix() {
	v1 := vec.Init(0, 1, 2, 3)
	v2 := vec.Init(3, 2, 1, 0)
	m = InitRows(v1, v2)
}

func TestAt(t *testing.T) {
	initMatrix()
	assert.Equal(t, 3., m.At(0, 3))
	assert.Equal(t, 0., m.At(1, 3))
}

func TestScale(t *testing.T) {
	initMatrix()
	m.Scale(2)
	assert.Equal(t, 0., m.At(0, 0))
	assert.Equal(t, 2., m.At(0, 1))
}

func TestSet(t *testing.T) {
	initMatrix()
	m.Set(0, 0, 500)
	assert.Equal(t, 500., m.At(0, 0))
}

func TestColCount(t *testing.T) {
	initMatrix()
	assert.Equal(t, 4, m.ColCount())
}

func TestRowCount(t *testing.T) {
	initMatrix()
	assert.Equal(t, 2, m.RowCount())
}

func TestCol(t *testing.T) {
	initMatrix()
	v := m.Col(1)
	assert.Equal(t, 2, v.Len())
	assert.Equal(t, 1., v.At(0))
}

func TestSwapCols(t *testing.T) {
	initMatrix()
	m.SwapCols(0, 1)
	assert.Equal(t, 1., m.Col(0).At(0))
	assert.Equal(t, 0., m.Col(1).At(0))
}

func TestSlice(t *testing.T) {
	initMatrix()
	m = Slice(m, 0, 1)

	assert.Equal(t, 1, m.ColCount())
	assert.Equal(t, 3., m.Col(0).At(1))
}

func TestTranspose(t *testing.T) {
	initMatrix()
	m.Transpose()
	assert.Equal(t, 4, m.RowCount())
	assert.Equal(t, 2, m.ColCount())
}

func TestAddCol(t *testing.T) {
	initMatrix()
	v := vec.Init(1, 2)
	m.AddCol(v)

	assert.Equal(t, 1., m.At(0, 0))
	assert.Equal(t, 2., m.At(1, 3))
}
