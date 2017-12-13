package activation

import (
	"github.com/jacsmith21/gnn"
	"github.com/jacsmith21/gnn/vec"
)

// ReLU ReLU
type ReLU struct{}

// Forward Forward
func (r ReLU) Forward(v vec.Vector) {

}

// Backward Backward
func (r ReLU) Backward(g gnn.Grad) {

}
