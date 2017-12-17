package gnn

import (
	"github.com/jacsmith21/gnn/mat"
)

// ReLU ReLU
type ReLU struct {
	z mat.Matrix
}

// Forward Forward
func (r *ReLU) Forward(z mat.Matrix) mat.Matrix {
	r.z = mat.Copy(z)
	z.ReLU()
	return z
}

// BackProp BackProp
func (r *ReLU) BackProp(t Trainer, dz mat.Matrix) mat.Matrix {
	r.z.ReLUDer()
	dz.Mul(r.z)
	return dz
}

// Sigmoid is the sigmoid activation struct
type Sigmoid struct {
	z mat.Matrix
}

// Forward applies the sigmoid function to the given matrix and returns the output
func (s *Sigmoid) Forward(z mat.Matrix) mat.Matrix {
	s.z = mat.Copy(z)
	z.Sigmoid()
	return z
}

// BackProp applies the sigmoid derivative to the given vector and returns the output
func (s *Sigmoid) BackProp(t Trainer, dz mat.Matrix) mat.Matrix {
	s.z.SigmoidDer()
	dz.Mul(s.z)
	return dz
}

// Softmax is the softmax activation struct
type Softmax struct {
	z mat.Matrix
}

// Forward applies the softmax function to each column
func (s *Softmax) Forward(z mat.Matrix) mat.Matrix {
	s.z = mat.Copy(z)
	for i := 0; i < z.ColCount(); i++ {
		z.Col(i).Softmax()
	}

	return z
}

// BackProp applies the softmax derivative to the given matrix and returns the output
func (s *Softmax) BackProp(t Trainer, dz mat.Matrix) mat.Matrix {
	s.z.SigmoidDer()
	dz.Mul(s.z)
	return dz
}
