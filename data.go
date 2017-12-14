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
type DataSet []Sample

// Len Len
func (d DataSet) Len() int {
	return len(d)
}

// Shuffle Shuffle
func (d DataSet) Shuffle() {
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

// Batches Batches
func (d DataSet) Batches(batchSize int) []DataSet {
	sampleCount := d.Len()
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

		batches[i] = d[from:to]
	}

	return batches
}
