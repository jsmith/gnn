package gnn

import "github.com/jacsmith21/gnn/vec"

// ReLU ReLU
type ReLU struct{}

// Forward Forward
func (r ReLU) Forward(z []vec.Vector) {
	for i := range z {
		z[i].ReLU()
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
