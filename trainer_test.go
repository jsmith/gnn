package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/data"
	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var xor data.DataSet

func initXORDataSet() {
	xor = data.Init(
		mat.InitRows(
			vec.Init(0, 0, 1, 1),
			vec.Init(0, 1, 1, 0),
		),
		mat.InitRows(
			vec.Init(0, 1, 1, 0),
		),
	)
}

func TestBadTrainer(t *testing.T) {
	net := Net{
		NewFC(1, 1),
		ReLU{},
	}

	trainer := Trainer{
		Net:    net,
		Cost:   SE{},
		Epochs: 500,
	}

	assert.Panics(t, func() { trainer.Train(xor) })
}

func TestTrainer(t *testing.T) {
	initXORDataSet()

	net := Net{
		NewFC(2, 4),
		ReLU{},
		NewFC(4, 1),
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
	assert.Equal(t, trainer, trainer)
}
