package vec64

import (
	"math"

	"github.com/jacsmith21/gnn/vec"
)

// Vec64 Vec64
type Vec64 struct {
	slice []float64
}

// Creator Creator
func (v Vec64) Creator() vec.Creator {
	return Creator{}
}

// At At
func (v Vec64) At(i int) vec.Number {
	return v.slice[i]
}

// Set Set
func (v Vec64) Set(i int, n vec.Number) {
	v.slice[i] = n.(float64)
}

// Add Add
func (v Vec64) Add(other vec.Vector) {
	for i, n := range other.(Vec64).slice {
		v.slice[i] += n
	}
}

// AddScalar AddScalar
func (v Vec64) AddScalar(n vec.Number) {
	f := n.(float64)
	for i := range v.slice {
		v.slice[i] += f
	}
}

// Sub Sum
func (v Vec64) Sub(other vec.Vector) {
	for i, n := range other.(Vec64).slice {
		v.slice[i] -= n
	}
}

// Len returns len(v)
func (v Vec64) Len() int {
	return len(v.slice)
}

// Mul Mul
func (v Vec64) Mul(n vec.Number) {
	f := n.(float64)
	for i := range v.slice {
		v.slice[i] *= f
	}
}

// Pow Pow
func (v Vec64) Pow(n vec.Number) {
	f := n.(float64)
	for i := range v.slice {
		v.slice[i] = math.Pow(v.slice[i], f)
	}
}

// Exp Exp
func (v Vec64) Exp() {
	for i := range v.slice {
		v.slice[i] = math.Exp(v.slice[i])
	}
}

// Sigmoid Sigmoid
func (v Vec64) Sigmoid() {
	c := v.Creator()
	v.Mul(c.Number(-1))
	v.AddScalar(c.Number(1))
	v.Pow(c.Number(-1))
}
