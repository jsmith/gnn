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

// BackProp BackProp
func (r ReLU) BackProp(dz mat.Matrix) mat.Matrix {
	dz.ReLUDer()
	return dz
}

// Sigmoid Sigmoid
type Sigmoid struct{}

// Forward applies the sigmoid function to the given matrix and returns the output
func (s Sigmoid) Forward(z mat.Matrix) mat.Matrix {
	z.Sigmoid()
	return z
}

// BackProp applies the sigmoid derivative to the given vector and returns the output
func (s Sigmoid) BackProp(dz mat.Matrix) mat.Matrix {
	dz.SigmoidDer()
	return dz
}
