package gnn

import (
	"testing"
)

var net Net

func initNet() {
	net = Net{
		NewFC(10, 10),
		ReLU{},
		NewFC(10, 1),
		Sigmoid{},
	}
}

func TestNetForward(t *testing.T) {

}
