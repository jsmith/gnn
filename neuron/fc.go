package neuron

import (
	"math"

	"github.com/jacsmith21/gnn"
	"github.com/jacsmith21/gnn/vec"
)

// FC FC
type FC struct {
	In      int
	Out     int
	Weights vec.Vector
	Biases  vec.Vector
}

// NewFC NewFC
func NewFC(c vec.Creator, in, out int) *FC {
	fc := FC{
		In:      in,
		Out:     out,
		Weights: c.Make(in * out),
		Biases:  c.Make(out),
	}

	vec.Rand(fc.Weights, vec.Gaussian)
	fc.Weights.Scale(c.Number(1 / math.Sqrt(float64(in))))

	return &fc
}

// Forward Forward
func (f FC) Forward(v vec.Vector) {
	for i := 0; i < f.Out; i++ {

	}
}

// Backward Backward
func (f FC) Backward(g gnn.Grad) {

}
