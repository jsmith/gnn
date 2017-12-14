package gnn

import "github.com/jacsmith21/gnn/vec"

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
