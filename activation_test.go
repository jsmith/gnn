package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var activationMat mat.Matrix

func initActivationMatrix() {
	activationMat = mat.InitRows(
		vec.Init(-1, 0, 1),
	)
}

func TestReLU(t *testing.T) {
	initActivationMatrix()
	r := ReLU{}
	activationMat = r.Forward(activationMat)
	assert.Equal(t, 0., activationMat.At(0, 0))
	assert.Equal(t, 0., activationMat.At(0, 1))
	assert.Equal(t, 1., activationMat.At(0, 2))
}

func TestSigmoid(t *testing.T) {
	initActivationMatrix()
	r := Sigmoid{}
	activationMat = r.Forward(activationMat)
	assert.InDelta(t, (1 - 0.731058578630), activationMat.At(0, 0), 0.0001)
	assert.InDelta(t, 0.5, activationMat.At(0, 1), 0.00001)
	assert.InDelta(t, 0.731058578630, activationMat.At(0, 2), 0.00001)
}
