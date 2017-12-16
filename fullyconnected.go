package gnn

import (
	"math"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/rander"
	"github.com/jacsmith21/gnn/vec"
)

// FullyConnected a fully connected network layer
type FullyConnected struct {
	Weights     mat.Matrix
	Biases      vec.Vector
	activations mat.Matrix
	sampleCount int
}

// NewFC creates a net fully connected layer and initializes the weights using a Gaussian distribution
func NewFC(in, out int) *FullyConnected {
	fc := FullyConnected{
		Weights: mat.Make(out, in),
		Biases:  vec.Make(out),
	}

	rander.Rand(fc.Weights, rander.Gaussian)
	fc.Weights.Scale(1 / math.Sqrt(float64(in)))

	return &fc
}

// InitFC initializes a FC layer with the given weights and biases
func InitFC(weights mat.Matrix, biases vec.Vector) *FullyConnected {
	if weights.ColCount() == 0 || biases.Len() == 0 {
		panic("the weights and biases cannot be empty")
	}

	return &FullyConnected{
		Weights: weights,
		Biases:  biases,
	}
}

// Forward applies the forward operation of a fully connected layer
func (f FullyConnected) Forward(a mat.Matrix) mat.Matrix {
	f.activations = mat.Copy(a) // caching for back propogation
	a = mat.Mul(f.Weights, a)
	a.AddCol(f.Biases)
	return a
}

// BackProp applies the backward operation of a fully connected layer
func (f FullyConnected) BackProp(t Trainer, dz mat.Matrix) mat.Matrix {
	da := mat.Mul(dz, f.Weights)
	f.updateParams(t, dz)
	return da
}

func (f FullyConnected) updateParams(t Trainer, dz mat.Matrix) {
	f.activations.Transpose()
	dw := mat.Mul(dz, f.activations)
	db := mat.SumCols(dz)

	scale := (1. / float64(f.sampleCount)) * float64(t.LearningRate)
	dw.Scale(scale)
	db.Scale(scale)
	f.Weights.Sub(dw)
	f.Biases.Sub(db)
}
