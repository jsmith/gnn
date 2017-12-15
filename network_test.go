package gnn

import (
	"testing"
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

}
