package gnn

import (
	"github.com/jacsmith21/gnn/mat"
)

// ReLU ReLU
type ReLU struct{}

// Forward Forward
func (r ReLU) Forward(z mat.Matrix) mat.Matrix {
	z.ReLU()
	return z
}

// Backward Backward
func (r ReLU) Backward() {

}

// Sigmoid Sigmoid
type Sigmoid struct{}

// Forward Forward
func (s Sigmoid) Forward(z mat.Matrix) mat.Matrix {
	z.Sigmoid()
	return z
}

// Backward Backward
func (s Sigmoid) Backward() {

}
