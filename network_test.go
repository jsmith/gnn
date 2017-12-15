package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/neuron"
)

var net Net

func initNet() {
	net = Net{
		neuron.NewFC(10, 10),
		ReLU{},
		neuron.NewFC(10, 1),
		Sigmoid{},
	}
}

func TestNetForward(t *testing.T) {

}
