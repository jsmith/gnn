package test

import (
	"fmt"
	"testing"

	"github.com/jacsmith21/gnn"
	"github.com/jacsmith21/gnn/neuron"
)

// FCNN FCNN
func FCNN(t *testing.T) {

	nn := gnn.Net{
		neuron.NewFC(10, 10),
		gnn.ReLU{},
		neuron.NewFC(10, 1),
		gnn.Sigmoid{},
	}

	trainer := gnn.Trainer{
		LearningRate: 0.01,
	}

	epochs := 100
	for i := 1; i <= epochs; i++ {
		fmt.Print(nn)
		fmt.Print(trainer)
	}
}
