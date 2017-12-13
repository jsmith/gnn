package gnn

import "github.com/jacsmith21/gnn/vec"

// Layer a layer in a neural network
type Layer interface {
	Forward(x vec.Vector)
	Backward(grad Grad)
}

// Net a neural network
type Net []Layer

// Grad a gradient
type Grad float64
