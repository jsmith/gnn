package data

import (
	"github.com/jacsmith21/gnn/vec"
)

// Sample an individual sample
type Sample struct {
	data  vec.Vector
	label vec.Vector
}

// Label returns the label at the ith index
func (s Sample) Label(i int) float64 {
	return s.label.At(i)
}

// DataValue returns the data value at the ith index
func (s Sample) DataValue(i int) float64 {
	return s.data.At(i)
}
