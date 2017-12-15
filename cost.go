package gnn

import "github.com/jacsmith21/gnn/vec"

// Cost Cost
type Cost interface {
	Cost(expected vec.Vector, actual []vec.Vector) float64
	Der(expected vec.Vector, actual []vec.Vector) float64
}

// SE SE
type SE struct{}

// Cost Cost
func (mse SE) Cost(expected vec.Vector, actual []vec.Vector) float64 {
	for
	diff := vec.Sub(expected, actual)
	diff.Pow(2)
	return diff.Sum()
}

// Der Der
func (mse SE) Der(expected, actual []vec.Vector) float64 {
	diff := vec.Sub(expected, actual)
	return diff.Sum() * 2
}
