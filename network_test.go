package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var net Net
var netM mat.Matrix

func initNetDataSet() {
	netM = mat.InitRows(
		vec.Init(1, 1),
		vec.Init(3, 3),
	)
}

func initNet() {
	weights := mat.InitRows(
		vec.Init(1, 1),
		vec.Init(1, 1),
	)

	biases := vec.Init(
		1,
		1,
	)

	net = Net{
		InitFC(weights, biases),
		ReLU{},
		InitFC(weights, biases),
		Sigmoid{},
	}
}

func TestNetForward(t *testing.T) {
	initNet()
	initNetDataSet()
	output := net.Forward(netM)
	assert.Equal(t, 2, output.RowCount())
	assert.Equal(t, 2, output.ColCount())
	assert.InDelta(t, 0.999983298578152, output.At(0, 0), 0.00001)
}
