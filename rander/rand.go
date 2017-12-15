package rander

import (
	"math/rand"

	"github.com/jacsmith21/gnn/mat"
)

// Distribution Distribution
type Distribution int

const (
	// Gaussian Gaussian
	Gaussian Distribution = iota
)

// Rander Rander
type Rander interface {
	// Intn returns a pseudo-random number in [0,n)
	Intn(n int) int
}

// Rand Rand
func Rand(m mat.Matrix, d Distribution) {
	for i := 0; i < m.RowCount(); i++ {
		for j := 0; j < m.ColCount(); j++ {
			switch d {
			case Gaussian:
				m.Set(i, j, rand.NormFloat64())
			default:
				panic("unkown distribution")
			}
		}
	}
}
