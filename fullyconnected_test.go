package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var fc *FullyConnected

func initFC() {
	biases := vec.Init(1, 2)

	weights := mat.InitRows(
		vec.Init(1, 4),
		vec.Init(1, 2),
	)

	fc = InitFC(weights, biases)
}

func TestFCForward(t *testing.T) {
	initFC()
	data := mat.InitCols(
		vec.Init(1, 1),
	)
	z = fc.Forward(data)

	assert.Equal(t, 6., z.At(0, 0))
	assert.Equal(t, 5., z.At(1, 0))
}

func TestFullyConnectedBackProp(t *testing.T) {
	initFC()
	initDefaultTrainer()

	a := mat.InitRows(
		vec.Init(1),
		vec.Init(1),
	)
	fc.Forward(a)
	fc.BackProp(trainer, mat.InitRows(
		vec.Init(1),
		vec.Init(1),
	))

	assert.Equal(t, 0.9, fc.Weights.At(0, 0))
	assert.Equal(t, 1.9, fc.Weights.At(1, 1))
	assert.Equal(t, 0.9, fc.Biases.At(0))
}


