package gnn

import "github.com/jacsmith21/gnn/vec"

// Cost Cost
type Cost interface {
	Cost(expected, actual vec.Vector) float64
}

// MSE MSE
type MSE struct{}

// Cost Cost
func (mse MSE) Cost(expected, actual vec.Vector) float64 {
	c := expected.Creator()
	diff := c.Sub(expected, actual)
	diff.Pow(2)

	n := diff.Sum()
	f := c.Float64(n)
	f /= float64(diff.Len())

	return f
}
