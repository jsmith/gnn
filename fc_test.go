package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestFCForward(t *testing.T) {
	biases := vec.Init(1, 2)

	weights := mat.InitRows(
		vec.Init(1, 4),
		vec.Init(1, 2),
	)

	fc := InitFC(weights, biases)

	data := mat.InitCols(
		vec.Init(1, 1),
	)
	data = fc.Forward(data)

	assert.Equal(t, 6., data.At(0, 0))
	assert.Equal(t, 5., data.At(1, 0))
}
