package gnn

import (
	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
)

// Cost Cost
type Cost interface {
	Cost(exp, act mat.Matrix) vec.Vector
	Der(exp, act mat.Matrix) vec.Vector
}

// SE SE
type SE struct{}

// Cost Cost
func (mse SE) Cost(exp, act mat.Matrix) vec.Vector {
	cost := vec.Make(act.ColCount())
	for i := 0; i < act.ColCount(); i++ {
		diff := vec.Sub(act.Col(i), exp.Col(i))
		diff.Pow(2)
		cost.Set(i, diff.Sum())
	}
	return cost
}

// Der Der
func (mse SE) Der(exp, act mat.Matrix) vec.Vector {
	der := vec.Make(act.ColCount())
	for i := 0; i < act.ColCount(); i++ {
		diff := vec.Sub(act.Col(i), exp.Col(i))
		der.Set(i, diff.Sum()*2)
	}
	return der
}
