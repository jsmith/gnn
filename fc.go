package gnn

import (
	"math"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/rander"
	"github.com/jacsmith21/gnn/vec"
)

// FC FC
type FC struct {
	In      int
	Out     int
	Weights mat.Matrix
	Biases  vec.Vector
}

// NewFC NewFC
func NewFC(in, out int) *FC {
	fc := FC{
		In:      in,
		Out:     out,
		Weights: mat.Make(out, in),
		Biases:  vec.Make(out),
	}

	rander.Rand(fc.Weights, rander.Gaussian)
	fc.Weights.Scale(1 / math.Sqrt(float64(in)))

	return &fc
}

// InitFC initializes a FC layer with the given weights and biases
func InitFC(weights mat.Matrix, biases vec.Vector) *FC {
	if weights.ColCount() == 0 || biases.Len() == 0 {
		panic("the weights and biases cannot be empty")
	}

	in := weights.ColCount()
	out := biases.Len()
	return &FC{
		In:      in,
		Out:     out,
		Weights: weights,
		Biases:  biases,
	}
}

// Forward Forward
func (f FC) Forward(a mat.Matrix) {
	a = mat.Mul(f.Weights, a)
	a.AddCol(f.Biases)
}

// Backward Backward
func (f FC) Backward() {

}
