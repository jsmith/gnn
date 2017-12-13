package gnn

import "github.com/jacsmith21/gnn/vec"

// Label Label
type Label vec.Number

// Sample Sample
type Sample struct {
	x vec.Vector
	y Label
}

// DataSet DataSet
type DataSet []Sample

// Len Len
func (d DataSet) Len() int {
	return len(d)
}
