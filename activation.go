package gnn

import (
	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
)

// ReLU ReLU
type ReLU struct{}

// Forward Forward
func (r ReLU) Forward(z mat.Matrix) {
	for i := 0; i < z.ColCount(); i++ {
		z.Col(i).ReLU()
	}
}

// Backward Backward
func (r ReLU) Backward() {

}

// Sigmoid Sigmoid
type Sigmoid struct{}

// Forward Forward
func (s Sigmoid) Forward(v []vec.Vector) {

}

// Backward Backward
func (s Sigmoid) Backward() {

}
