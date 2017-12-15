package gnn

import "github.com/jacsmith21/gnn/vec"

// Cost Cost
type Cost interface {
	Cost(exp, act []vec.Vector) vec.Vector
	Der(exp, act []vec.Vector) vec.Vector
}

// SE SE
type SE struct{}

// Cost Cost
func (mse SE) Cost(exp, act []vec.Vector) vec.Vector {
	cost := vec.Make(len(act))
	for i := range act {
		diff := vec.Sub(act[i], exp[i])
		diff.Pow(2)
		cost.Set(i, diff.Sum())
	}
	return cost
}

// Der Der
func (mse SE) Der(exp, act []vec.Vector) vec.Vector {
	der := vec.Make(len(act))
	for i := range act {
		diff := vec.Sub(act[i], exp[i])
		der.Set(i, diff.Sum()*2)
	}
	return der
}
