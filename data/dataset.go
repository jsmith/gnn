package data

import (
	"math/rand"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
)

// DataSet a dataset
type DataSet struct {
	data   mat.Matrix
	labels mat.Matrix
}

// SampleCount SampleCount
func (d DataSet) SampleCount() int {
	return d.data.ColCount()
}

// Sample returns the ith sample
func (d DataSet) Sample(i int) Sample {
	return Sample{d.data.Col(i), d.labels.Col(i)}
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

		d.data.SwapCols(i, j)
		d.labels.SwapCols(i, j)
	}
}

// GenerateBatches GenerateBatches
func (d DataSet) GenerateBatches(batchSize int) []DataSet {
	if batchSize == 0 {
		panic("batch size cannot be 0")
	}
	copy := Copy(d)

	sampleCount := copy.SampleCount()
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

		batches[i] = DataSet{mat.Slice(copy.data, from, to), mat.Slice(copy.labels, from, to)}
	}

	return batches
}
