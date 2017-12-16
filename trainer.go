package gnn

import (
	"github.com/go-playground/validator"
	"github.com/jacsmith21/gnn/data"
	"github.com/jacsmith21/gnn/log"
	"github.com/jacsmith21/gnn/mat"
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
		panic(err)
	}

	if t.BatchSize == 0 {
		t.BatchSize = d.SampleCount()
	}

	for i := 1; i <= t.Epochs; i++ {
		// Shuffle if we're actually doing mini-batch
		if t.BatchSize < d.SampleCount() {
			log.Method("Train").Debug("shuffling DataSet")
			d.Shuffle(nil)
		}

		for _, batch := range d.GenerateBatches(t.BatchSize) {
			output := t.Net.Forward(batch.Data())
			da := t.Cost.Der(batch.Labels(), output)
			t.Net.BackProp(t, da)
		}
	}
}

func (t Trainer) Predict(m mat.Matrix) mat.Matrix {
	return t.Net.Forward(m)
}
