package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
)

var net Net

func initNet() {
	net = Net{
		NewFC(2, 2),
		ReLU{},
		NewFC(2, 1),
		Sigmoid{},
	}
}

func TestNetForward(t *testing.T) {
	net.Forward(mat.Init(vec.Init(1, 1)))
}
