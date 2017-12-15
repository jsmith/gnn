package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestFCInit(t *testing.T) {
	biases := vec.Init(
		1,
		2,
	)

	weights := mat.Init(
		vec.Init(
			1,
			1,
		),
		vec.Init(
			4,
			2,
		),
	)

	fc := InitFC(weights, biases)

	data := mat.Init(vec.Init(
		1,
		1,
	))
	fc.Forward(data)

	assert.Equal(t, 500., data.At(0, 0))
	assert.Equal(t, 500., data.At(1, 0))
}
