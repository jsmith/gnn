package vec

import (
	"math"
)

// Vector Vector
type Vector struct {
	slice []float64
}

// At At
func (v Vector) At(i int) float64 {
	return v.slice[i]
}

// Set Set
func (v Vector) Set(i int, f float64) {
	v.slice[i] = f
}

// SetData SetData
func (v Vector) SetData(data []float64) {
	copy(v.slice, data)
}

// Add Add
func (v Vector) Add(other Vector) {
	for i, n := range other.slice {
		v.slice[i] += n
	}
}

// AddScalar AddScalar
func (v Vector) AddScalar(f float64) {
	for i := range v.slice {
		v.slice[i] += f
	}
}

// Sub Sum
func (v Vector) Sub(other Vector) {
	for i, f := range other.slice {
		v.slice[i] -= f
	}
}

// Mul Mul
func (v Vector) Mul(other Vector) {
	for i, f := range other.slice {
		v.slice[i] *= f
	}
}

// Len returns len(v)
func (v Vector) Len() int {
	return len(v.slice)
}

// Scale Scale
func (v Vector) Scale(f float64) {
	for i := range v.slice {
		v.slice[i] *= f
	}
}

// Pow Pow
func (v Vector) Pow(f float64) {
	for i := range v.slice {
		v.slice[i] = math.Pow(v.slice[i], f)
	}
}

// Exp Exp
func (v Vector) Exp() {
	for i := range v.slice {
		v.slice[i] = math.Exp(v.slice[i])
	}
}

// Sigmoid Sigmoid
func (v Vector) Sigmoid() {
	v.Scale(-1)
	v.Exp()
	v.AddScalar(1)
	v.Pow(-1)
}

// Sum Sum
func (v Vector) Sum() float64 {
	sum := 0.
	for _, n := range v.slice {
		sum += n
	}

	return sum
}

// Ln Ln
func (v Vector) Ln() {
	for i, n := range v.slice {
		v.slice[i] = math.Log(n)
	}
}

// Swap Swap
func (v Vector) Swap(i, j int) {
	v.slice[i], v.slice[j] = v.slice[j], v.slice[i]
}

// ReLU ReLU
func (v Vector) ReLU() {
	for i, n := range v.slice {
		v.slice[i] = math.Max(n, 0)
	}
}
