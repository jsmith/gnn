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
	labels vec.Vector
}

// SampleCount SampleCount
func (d DataSet) SampleCount() int {
	return d.labels.Len()
}

// Shuffle Shuffle
func (d DataSet) Shuffle() {
	for i := d.SampleCount() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.data[i], d.data[j] = d.data[j], d.data[i]
		d.labels.Swap(i, j)
	}
}

// Batches Batches
func (d DataSet) Batches(batchSize int) []DataSet {
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

		batches[i] = DataSet{d.data[from:to], vec.Slice(d.labels, from, to)}
	}

	return batches
}
