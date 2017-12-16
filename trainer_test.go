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
		&ReLU{},
	}

	trainer := Trainer{
		Net:    net,
		Cost:   SE{},
		Epochs: 500,
	}

	assert.Panics(t, func() { trainer.Train(xor) })
}

var trainer Trainer

func initDefaultTrainer() {
	weights1 := mat.InitRows(
		vec.Init(0.1, -0.1),
		vec.Init(0.2, -0.4),
		vec.Init(0.05, -0.2),
		vec.Init(0.4, -0.05),
	)
	biases1 := vec.Make(4)

	weights2 := mat.InitRows(
		vec.Init(0.1, -0.1, 0.05, -0.05),
	)
	biases2 := vec.Make(1)

	net := Net{
		InitFC(weights1, biases1),
		&ReLU{},
		InitFC(weights2, biases2),
		&Sigmoid{},
	}

	trainer = Trainer{
		Net:          net,
		Cost:         SE{},
		LearningRate: 0.1,
		Epochs:       10000,
		BatchSize:    4,
	}
}

func TestTrainer(t *testing.T) {
	initXORDataSet()
	initDefaultTrainer()
	trainer.Train(xor)
	predictions := trainer.Predict(xor.Data())
	expected := mat.InitRows(
		vec.Init(0.010810160797699253, 0.9736054703566581, 0.9736054703566581, 0.009612902402324593),
	)
	assert.Equal(t, expected, predictions)
}
