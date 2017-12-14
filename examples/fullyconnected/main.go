package test

import (
	"fmt"
	"testing"

	"github.com/jacsmith21/gnn"
	"github.com/jacsmith21/gnn/activation"
	"github.com/jacsmith21/gnn/neuron"
)

// TestFCNN TestFCNN
func TestFCNN(t *testing.T) {

	nn := gnn.Net{
		neuron.NewFC(10, 10),
		activation.ReLU{},
		neuron.NewFC(10, 1),
		activation.Sigmoid{},
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
