package gnn

import "gonum.org/v1/gonum/mat"

// Layer a layer in a neural network
type Layer interface {
	Forward(x mat.Vector)
	Backward(grad Grad)
}

// Net a neural network
type Net []Layer

// Grad a gradient
type Grad float64
