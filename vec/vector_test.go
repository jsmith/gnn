package vec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var v Vector

func initVector() {
	v = Vector{[]float64{0, 1, 2, 3, 4}}
}

func TestAt(t *testing.T) {
	initVector()
	assert.Equal(t, 3., v.At(3))
}

func TestSet(t *testing.T) {
	initVector()
	v.Set(3, 6)
	assert.Equal(t, 6., v.At(3))
}

func TestSetData(t *testing.T) {
	initVector()
	list := []float64{4, 3, 2, 1, 0}
	v.SetData(list)
	assert.Equal(t, 4., v.At(0))
	assert.Equal(t, 1., v.At(3))
}

func TestScale(t *testing.T) {
	initVector()
	v.Scale(2)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 6., v.At(3))
}

func TestExp(t *testing.T) {
	initVector()
	v.Exp()
	assert.Equal(t, 1., v.At(0))
	assert.InDelta(t, 2.71828183, v.At(1), 0.00001)
}

func TestPow(t *testing.T) {
	initVector()
	v.Pow(2)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 9., v.At(3))
}

func TestSub(t *testing.T) {
	initVector()
	other := Vector{[]float64{0, 1, 2, 3, 4}}
	v.Sub(other)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 0., v.At(4))
}

func TestAdd(t *testing.T) {
	initVector()
	other := Vector{[]float64{0, 1, 2, 3, 4}}
	v.Add(other)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 6., v.At(3))
}

func TestAddScalar(t *testing.T) {
	initVector()
	v.AddScalar(5)
	assert.Equal(t, 5., v.At(0))
	assert.Equal(t, 8., v.At(3))
}

func TestLen(t *testing.T) {
	initVector()
	assert.Equal(t, 5, v.Len())
}

func TestSigmoid(t *testing.T) {
	initVector()
	v.Sigmoid()
	assert.Equal(t, 0.5, v.At(0))
	assert.InDelta(t, 0.952574126822, v.At(3), 0.000001)
}

func TestSum(t *testing.T) {
	initVector()
	f := v.Sum()
	assert.InDelta(t, 10., f, 0.00000001)
}

func TestMul(t *testing.T) {
	initVector()
	v.Mul(Vector{[]float64{0, 1, 2, 3, 4}})
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 9., v.At(3))
}

func TestLn(t *testing.T) {
	initVector()
	v.Slice[0] = 1 // ln(0) is undefined
	v.Ln()
	assert.Equal(t, 0., v.At(0))
	assert.InDelta(t, 1.09861228, v.At(3), 0.00001)
}
