package gnn

import (
	"math/rand"

	"github.com/jacsmith21/gnn/vec"
)

// Label Label
type Label float64

// Sample Sample
type Sample struct {
	Input vec.Vector
	Label Label
}

// DataSet DataSet
type DataSet struct {
	data   []vec.Vector
	labels []vec.Vector
}

// SampleCount SampleCount
func (d DataSet) SampleCount() int {
	return len(d.data)
}

// Shuffle Shuffle
func (d DataSet) Shuffle(r vec.Rander) {
	for i := d.SampleCount() - 1; i > 0; i-- {
		var j int
		if r == nil {
			j = rand.Intn(i + 1)
		} else {
			j = r.Intn(i + 1)
		}

		d.data[i], d.data[j] = d.data[j], d.data[i]
		d.labels[i], d.labels[j] = d.labels[j], d.labels[i]
	}
}

// GenerateBatches GenerateBatches
func (d DataSet) GenerateBatches(batchSize int) []DataSet {
	if batchSize == 0 {
		panic("batch size cannot be 0")
	}

	sampleCount := d.SampleCount()
	batchCount := sampleCount / batchSize
	if sampleCount%batchSize != 0 {
		batchCount++
	}

	batches := make([]DataSet, batchCount)
	for i := 0; i < batchCount; i++ {
		from := i * batchSize
		to := from + batchSize
		if to > sampleCount {
			to = sampleCount
		}

		batches[i] = DataSet{d.data[from:to], d.labels[from:to]}
	}

	return batches
}
