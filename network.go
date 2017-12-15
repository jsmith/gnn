package gnn

import "gonum.org/v1/gonum/mat"

// Layer a layer in a neural network
type Layer interface {
	Forward(x mat.Matrix)
	Backward()
}

// Net a neural network
type Net []Layer

// Forward Forward
func (n Net) Forward(input mat.Matrix) mat.Matrix {
	for _, layer := range n {
		layer.Forward(input)
	}

	return input
}
