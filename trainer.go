package gnn

import "fmt"

// LearningRate LearningRate
type LearningRate float64

// Trainer Trainer
type Trainer struct {
	Net          Net
	Cost         Cost
	LearningRate LearningRate
	BatchSize    int
	Epochs       int
	Status       func(cost float64)
}

// Train Train
func (t Trainer) Train(d DataSet) {
	for i := 1; i <= t.Epochs; i++ {

		// Shuffle if we're actually doing mini-batch
		if t.BatchSize < d.SampleCount() {
			d.Shuffle(nil)
		}

		for _, batch := range d.GenerateBatches(t.BatchSize) {
			output := batch.data
			for _, layer := range t.Net {
				output = layer.Forward(output)
			}

			derivative := t.Cost.Der(batch.labels, output)
			fmt.Println(derivative)
		}
	}
}
