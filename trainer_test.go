package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/neuron"
	"github.com/jacsmith21/gnn/vec"
)

var xor DataSet

func initXORDataSet() {
	xor = DataSet{
		data: []vec.Vector{
			vec.Init([]float64{0, 0}),
			vec.Init([]float64{0, 1}),
			vec.Init([]float64{1, 1}),
			vec.Init([]float64{1, 0}),
		},
		labels: []vec.Vector{
			vec.Init([]float64{0}),
			vec.Init([]float64{1}),
			vec.Init([]float64{1}),
			vec.Init([]float64{0}),
		},
	}
}

func TestTrainer(t *testing.T) {
	initXORDataSet()

	net := Net{
		neuron.NewFC(10, 10),
		ReLU{},
		neuron.NewFC(10, 1),
		Sigmoid{},
	}

	trainer := Trainer{
		Net:          net,
		Cost:         SE{},
		LearningRate: 0.01,
		Epochs:       100,
		BatchSize:    4,
	}

	trainer.Train(xor)

}
