package vec32

import "github.com/jacsmith21/gnn/math"

// Vec32 implementation of Vector
type Vec32 struct {
	slice []float32
}

func (v Vec32) Creator() math.Creator {
	return Creator32{}
}

// At retuns v[i]
func (v Vec32) At(i int) math.Number {
	return v.slice[i]
}

// Set sets v[i] = n
func (v Vec32) Set(i int, n math.Number) {
	v.slice[i] = n.(float32)
}

// Add adds v[i] += other[i] for i to len(other)
func (v Vec32) Add(other math.Vector) {
	for i, n := range other.(Vec32).slice {
		v.slice[i] += n
	}
}

// Sub adds v[i] -= other[i] for i to len(other)
func (v Vec32) Sub(other math.Vector) {
	for i, n := range other.(Vec32).slice {
		v.slice[i] -= n
	}
}

// Len returns len(v)
func (v Vec32) Len() int {
	return len(v.slice)
}
