package vec

import "math/rand"

// Distribution Distribution
type Distribution int

const (
	// Gaussian Gaussian
	Gaussian Distribution = iota
)

// Rand Rand
func Rand(v Vector, d Distribution) {
	data := make([]float64, v.Len())

	for i := range data {
		switch d {
		case Gaussian:
			data[i] = rand.NormFloat64()
		default:
			panic("unkown distribution")
		}
	}

	v.SetData(data)
}
