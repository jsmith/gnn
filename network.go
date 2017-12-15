package gnn

import "github.com/jacsmith21/gnn/vec"

// Layer a layer in a neural network
type Layer interface {
	Forward(x []vec.Vector) []vec.Vector
	Backward()
}

// Net a neural network
type Net []Layer

// Forward Forward
func (n Net) Forward(d DataSet) {

}
