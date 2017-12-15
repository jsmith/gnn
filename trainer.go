package gnn

import (
	"fmt"

	"github.com/go-playground/validator"
)

// LearningRate LearningRate
type LearningRate float64

// Trainer Trainer
type Trainer struct {
	Net          Net                `validate:"required"`
	Cost         Cost               `validate:"required"`
	LearningRate LearningRate       `validate:"required"`
	BatchSize    int                `validate:"required"`
	Epochs       int                `validate:"required"`
	Status       func(cost float64) `validate:"required"`
}

// Train Train
func (t Trainer) Train(d DataSet) {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		panic("trainer not set up correctly")
	}

	for i := 1; i <= t.Epochs; i++ {
		fmt.Println(d)

		// Shuffle if we're actually doing mini-batch
		if t.BatchSize < d.SampleCount() {
			d.Shuffle(nil)
		}

		for _, batch := range d.GenerateBatches(t.BatchSize) {
			output := t.Net.Forward(batch.data)

			derivative := t.Cost.Der(batch.labels, output)
			fmt.Println("hey")
			fmt.Println(derivative)
		}
	}
}
