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
			d.Shuffle()
		}

		for _, batch := range d.Batches(t.BatchSize) {
			for _, layer := range t.Net {
				fmt.Println(batch)
				fmt.Println(layer)
			}
		}
	}
}
