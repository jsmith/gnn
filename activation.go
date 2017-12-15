package gnn

import "github.com/jacsmith21/gnn/vec"

// ReLU ReLU
type ReLU struct{}

// Forward Forward
func (r ReLU) Forward(v []vec.Vector) []vec.Vector {
	return nil
}

// Backward Backward
func (r ReLU) Backward() {

}

// Sigmoid Sigmoid
type Sigmoid struct{}

// Forward Forward
func (s Sigmoid) Forward(v []vec.Vector) []vec.Vector {
	return nil
}

// Backward Backward
func (s Sigmoid) Backward() {

}
