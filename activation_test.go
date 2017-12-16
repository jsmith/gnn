package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var z mat.Matrix

func initActivationMatrix() {
	z = mat.InitRows(
		vec.Init(-1, 0, 1),
	)
}

func TestReLUForward(t *testing.T) {
	initActivationMatrix()
	r := ReLU{}
	a := r.Forward(z)
	assert.Equal(t, 0., a.At(0, 0))
	assert.Equal(t, 0., a.At(0, 1))
	assert.Equal(t, 1., a.At(0, 2))
}

func TestSigmoidForward(t *testing.T) {
	initActivationMatrix()
	r := Sigmoid{}
	a := r.Forward(z)
	assert.InDelta(t, 1 - 0.731058578630, a.At(0, 0), 0.0001)
	assert.InDelta(t, 0.5, a.At(0, 1), 0.00001)
	assert.InDelta(t, 0.731058578630, a.At(0, 2), 0.00001)
}

func TestReLUBackProp(t *testing.T) {
	initActivationMatrix()
	initDefaultTrainer()
	r := ReLU{}
	r.Forward(z)
	output := r.BackProp(trainer, mat.InitRows(vec.Init(1, 2, 3)))
	assert.Equal(t, 0., output.At(0, 0))
	assert.Equal(t, 0., output.At(0, 1))
	assert.Equal(t, 3., output.At(0, 2))
}

func TestSigmoidBackProp(t *testing.T) {
	initActivationMatrix()
	initDefaultTrainer()
	s := Sigmoid{}
	s.Forward(z)
	output := s.BackProp(trainer, mat.InitRows(vec.Init(1, 2, 3)))
	assert.InDelta(t, 0.1966119332, output.At(0, 0), 0.000001)
	assert.Equal(t, 0.5, output.At(0, 1))
	assert.InDelta(t, 3.*0.1966119332, output.At(0, 2), 0.0000001)
}
