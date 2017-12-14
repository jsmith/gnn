package gnn

import "github.com/jacsmith21/gnn/vec"

// ReLU ReLU
type ReLU struct{}

// Forward Forward
func (r ReLU) Forward(v vec.Vector) {

}

// Backward Backward
func (r ReLU) Backward(g Grad) {

}

// Sigmoid Sigmoid
type Sigmoid struct{}

// Forward Forward
func (s Sigmoid) Forward(v vec.Vector) {

}

// Backward Backward
func (s Sigmoid) Backward(g Grad) {

}
