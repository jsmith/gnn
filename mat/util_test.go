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

func TestMatMul(t *testing.T) {
	m1 := InitRows(
		vec.Init(0, 1, 1),
		vec.Init(1, 0, 1),
	)

	m2 := InitRows(
		vec.Init(2, 2),
		vec.Init(2, 2),
		vec.Init(2, 2),
	)

	mat := Mul(m1, m2)
	assert.Equal(t, 2, mat.RowCount())
	assert.Equal(t, 2, mat.ColCount())
	assert.Equal(t, 4., mat.At(0, 0))
	assert.Equal(t, 4., mat.At(1, 0))
}

func TestSumCols(t *testing.T) {
	initMatrix()
	sums := SumCols(m)
	assert.Equal(t, 6., sums.At(0))
	assert.Equal(t, 2, sums.Len())
}

func TestSubTwoMatrices(t *testing.T) {
	initMatrix()
	m2 := InitRows(
		vec.Init(1, 1, 0, 0),
		vec.Init(1, 1, 0, 0),
	)
	m2 = Sub(m, m2)
	assert.Equal(t, -1., m2.At(0, 0))
	assert.Equal(t, 2., m2.At(1, 0))
}

func TestUtilTranspose(t *testing.T) {
	initMatrix()
	trans := Transpose(m)

	assert.Equal(t, 4, trans.RowCount())
	assert.Equal(t, 2, trans.ColCount())
	assert.Equal(t, 2, m.RowCount())
	assert.Equal(t, 4, m.ColCount())
}
