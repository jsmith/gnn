package test

import (
	"testing"

	"github.com/jacsmith21/gnn"
	"github.com/jacsmith21/gnn/activation"
	"github.com/jacsmith21/gnn/neuron"
	"github.com/jacsmith21/gnn/vec/vec64"
)

func TestFCNN(t *testing.T) {
	c := vec64.Creator{}

	nn := gnn.Net{
		neuron.NewFC(c, 10, 10),
		activation.ReLU{},
		neuron.NewFC(c, 10, 1),
		activation.Sigmoid{},
	}

  trainer =

  epochs := 100
  for i := 1; i <= epochs; i++ {

  }
}
