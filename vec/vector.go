package vec

import (
	"math"
)

// Vector Vector
type Vector struct {
	Slice []float64
}

// At At
func (v Vector) At(i int) float64 {
	return v.Slice[i]
}

// Set Set
func (v Vector) Set(i int, f float64) {
	v.Slice[i] = f
}

// SetData SetData
func (v Vector) SetData(data []float64) {
	copy(v.Slice, data)
}

// Add Add
func (v Vector) Add(other Vector) {
	for i, n := range other.Slice {
		v.Slice[i] += n
	}
}

// AddScalar AddScalar
func (v Vector) AddScalar(f float64) {
	for i := range v.Slice {
		v.Slice[i] += f
	}
}

// Sub Sum
func (v Vector) Sub(other Vector) {
	for i, f := range other.Slice {
		v.Slice[i] -= f
	}
}

// Mul Mul
func (v Vector) Mul(other Vector) {
	for i, f := range other.Slice {
		v.Slice[i] *= f
	}
}

// Len returns len(v)
func (v Vector) Len() int {
	return len(v.Slice)
}

// Scale Scale
func (v Vector) Scale(f float64) {
	for i := range v.Slice {
		v.Slice[i] *= f
	}
}

// Pow Pow
func (v Vector) Pow(f float64) {
	for i := range v.Slice {
		v.Slice[i] = math.Pow(v.Slice[i], f)
	}
}

// Exp Exp
func (v Vector) Exp() {
	for i := range v.Slice {
		v.Slice[i] = math.Exp(v.Slice[i])
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
	for _, n := range v.Slice {
		sum += n
	}

	return sum
}

// Ln Ln
func (v Vector) Ln() {
	for i, n := range v.Slice {
		v.Slice[i] = math.Log(n)
	}
}
