package gnn

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jacsmith21/gnn/data"
)

// LearningRate LearningRate
type LearningRate float64

// Trainer Trainer
type Trainer struct {
	Net          Net          `validate:"required"`
	Cost         Cost         `validate:"required"`
	LearningRate LearningRate `validate:"required"`
	BatchSize    int          `validate:"required"`
	Epochs       int          `validate:"required"`
	Status       func(cost float64)
}

// Train Train
func (t Trainer) Train(d data.DataSet) {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		panic("trainer not set up correctly")
	}

	for i := 1; i <= t.Epochs; i++ {
		// Shuffle if we're actually doing mini-batch
		if t.BatchSize < d.SampleCount() {
			d.Shuffle(nil)
		}

		for _, batch := range d.GenerateBatches(t.BatchSize) {
			output := t.Net.Forward(batch.Data())
			derivative := t.Cost.Der(batch.Labels(), output)

			fmt.Println(derivative)
		}
	}
}
