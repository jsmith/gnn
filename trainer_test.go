package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/data"
	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
)

var xor data.DataSet

func initXORDataSet() {
	xor = data.Init(
		mat.Init(
			vec.Init(0, 0),
			vec.Init(0, 1),
			vec.Init(1, 1),
			vec.Init(1, 0),
		),
		mat.Init(
			vec.Init(0),
			vec.Init(1),
			vec.Init(1),
			vec.Init(0),
		),
	)
}

func TestTrainer(t *testing.T) {
	initXORDataSet()

	net := Net{
		NewFC(10, 10),
		ReLU{},
		NewFC(10, 1),
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
