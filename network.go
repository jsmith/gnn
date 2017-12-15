package gnn

import "github.com/jacsmith21/gnn/vec"

// Layer a layer in a neural network
type Layer interface {
	Forward(x []vec.Vector)
	Backward()
}

// Net a neural network
type Net []Layer

// Forward Forward
func (n Net) Forward(input []vec.Vector) []vec.Vector {
	for _, layer := range n {
		layer.Forward(input)
	}

	return input
}
