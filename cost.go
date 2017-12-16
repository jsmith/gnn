package gnn

import (
	"github.com/jacsmith21/gnn/mat"
)

// Cost is the cost interface
type Cost interface {
	Cost(exp, act mat.Matrix) mat.Matrix
	Der(exp, act mat.Matrix) mat.Matrix
}

// SE is the squared error struct
type SE struct{}

// Cost computes the squared error cost
func (s SE) Cost(exp, act mat.Matrix) mat.Matrix {
	diff := mat.Sub(act, exp)
	diff.Pow(2)
	return diff
}

// Der computes the squared error derivatives
func (s SE) Der(exp, act mat.Matrix) mat.Matrix {
	diff := mat.Sub(act, exp)
	diff.Scale(2)
	return diff
}
