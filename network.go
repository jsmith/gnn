package gnn

import "github.com/jacsmith21/gnn/mat"

// Layer a layer in a neural network
type Layer interface {
	Forward(x mat.Matrix) mat.Matrix
	Backward()
}

// Net a neural network
type Net []Layer

// Forward Forward
func (n Net) Forward(input mat.Matrix) {
	output := input
	for _, layer := range n {
		output = layer.Forward(output)
	}
}
