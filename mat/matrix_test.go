package mat

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var m Matrix

func initMatrix() {
	m = InitRows(
		vec.Init(0, 1, 2, 3),
		vec.Init(3, 2, 1, 0),
	)
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

func TestReLU(t *testing.T) {
	initMatrix()
	m.Set(0, 0, -1)
	m.ReLU()

	assert.Equal(t, 0., m.At(0, 0))
	assert.Equal(t, 1., m.At(0, 1))
}

func TestReLUDer(t *testing.T) {
	initMatrix()
	m.Set(0, 0, -1)
	m.ReLUDer()

	assert.Equal(t, 0., m.At(0, 0))
	assert.Equal(t, 1., m.At(0, 1))
	assert.Equal(t, 1., m.At(0, 2))
}

func TestSigmoid(t *testing.T) {
	initMatrix()
	m.Sigmoid()

	assert.Equal(t, 0.5, m.At(0, 0))
	assert.InDelta(t, 0.731058578630004, m.At(0, 1), 0.00000001)
}

func TestSigmoidDer(t *testing.T) {
	initMatrix()
	m.SigmoidDer()

	assert.Equal(t, 0.25, m.At(0, 0))
	assert.InDelta(t, 0.196611933241481, m.At(0, 1), 0.00000001)
}

func TestString(t *testing.T) {
	initMatrix()
	mat := InitRows(
		vec.Init(1, 5),
		vec.Init(1, 4),
	)

	str := mat.String()
	assert.Equal(t, str, "[ 1 5 ]\n[ 1 4 ]\n")
}

func TestMul(t *testing.T) {
	initMatrix()
	m2 := InitRows(
		vec.Init(1, 1, 0, 0),
		vec.Init(1, 1, 0, 0),
	)
	m.Mul(m2)

	assert.Equal(t, 1., m.At(0, 1))
	assert.Equal(t, 0., m.At(0, 2))
	assert.Equal(t, 0., m.At(1, 2))
}

func TestSub(t *testing.T) {
	initMatrix()
	m2 := InitRows(
		vec.Init(1, 1, 0, 0),
		vec.Init(1, 1, 0, 0),
	)
	m.Sub(m2)

	assert.Equal(t, -1., m.At(0, 0))
	assert.Equal(t, 2., m.At(1, 0))
}

func TestPow(t *testing.T) {
	initMatrix()
	m.Pow(2)

	assert.Equal(t, 0., m.At(0, 0))
	assert.Equal(t, 9., m.At(1, 0))
}

func TestAppend(t *testing.T) {
	m := InitRows(
		vec.Init(1, 1),
	)

	m.Append(vec.Init(2, 2), 1)
	assert.Equal(t, InitRows(
		vec.Init(1, 1),
		vec.Init(2, 2),
	), m)

	m.Append(vec.Init(3, 3), 2)
	assert.Equal(t, InitRows(
		vec.Init(1, 1, 3),
		vec.Init(2, 2, 3),
	), m)

	assert.Panics(t, func() { m.Append(vec.Make(2), 3)})
}

func TestRow(t *testing.T) {
	initMatrix()
	assert.Equal(t, vec.Init(3, 2, 1, 0), m.Row(1))
}
