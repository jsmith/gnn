package gnn

import (
	"math/rand"

	"github.com/jacsmith21/gnn/vec"
)

// DataSet DataSet
type DataSet struct {
	data   vec.Matrix
	labels vec.Vector
}

// SampleCount SampleCount
func (d DataSet) SampleCount() int {
	return d.data.ColCount()
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

		d.data.Swap(i, j)
		d.laels.Swap(i, j)
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

		batches[i] = DataSet{copy.data.Slice(from, to), copy.labels.Slice(from, to)}
	}

	return batches
}

// Copy Copy
func Copy(d DataSet) DataSet {
	data := c.data.Copy()
	labels := c.labels.Copy()
	return DataSet{
		data:   data,
		labels: labels,
	}
}
