package vec64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var v Vec64

func initVec64() {
	v = Vec64{[]float64{0, 1, 2, 3, 4}}
}

func TestAt(t *testing.T) {
	initVec64()
	assert.Equal(t, 3., v.At(3))
}

func TestSet(t *testing.T) {
	initVec64()
	v.Set(3, v.Creator().Number(6))
	assert.Equal(t, 6., v.At(3))
}

func TestSetData(t *testing.T) {
	initVec64()
	list := v.Creator().List([]float64{4, 3, 2, 1, 0})
	v.SetData(list)
	assert.Equal(t, 4., v.At(0))
	assert.Equal(t, 1., v.At(3))
}

func TestScale(t *testing.T) {
	initVec64()
	v.Scale(v.Creator().Number(2))
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 6., v.At(3))
}

func TestExp(t *testing.T) {
	initVec64()
	v.Exp()
	assert.Equal(t, 1., v.At(0))
	assert.InDelta(t, 2.71828183, v.At(1), 0.00001)
}

func TestPow(t *testing.T) {
	initVec64()
	v.Pow(v.Creator().Number(2))
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 9., v.At(3))
}

func TestSub(t *testing.T) {
	initVec64()
	other := Vec64{[]float64{0, 1, 2, 3, 4}}
	v.Sub(other)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 0., v.At(4))
}

func TestAdd(t *testing.T) {
	initVec64()
	other := Vec64{[]float64{0, 1, 2, 3, 4}}
	v.Add(other)
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 6., v.At(3))
}

func TestAddScalar(t *testing.T) {
	initVec64()
	v.AddScalar(v.Creator().Number(5))
	assert.Equal(t, 5., v.At(0))
	assert.Equal(t, 8., v.At(3))
}

func TestLen(t *testing.T) {
	initVec64()
	assert.Equal(t, 5, v.Len())
}

func TestSigmoid(t *testing.T) {
	initVec64()
	v.Sigmoid()
	assert.Equal(t, 0.5, v.At(0))
	assert.InDelta(t, 0.952574126822, v.At(3), 0.000001)
}

func TestSum(t *testing.T) {
	initVec64()
	n := v.Sum()
	f := v.Creator().Float64(n)
	assert.InDelta(t, 10., f, 0.00000001)
}

func TestMul(t *testing.T) {
	initVec64()
	v.Mul(Vec64{[]float64{0, 1, 2, 3, 4}})
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 9., v.At(3))
}
