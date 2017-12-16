package vec

import (
	"fmt"
	"math"
)

// Vector is the vector struct
type Vector struct {
	slice []float64
}

// At returns the ith element
func (v Vector) At(i int) float64 {
	return v.slice[i]
}

// Set sets the ith element to the given float
func (v Vector) Set(i int, f float64) {
	v.slice[i] = f
}

// SetData replaces the data of the vector
func (v Vector) SetData(data []float64) {
	copy(v.slice, data)
}

// Add does the element wise addition with the given vector
func (v Vector) Add(other Vector) {
	for i, n := range other.slice {
		v.slice[i] += n
	}
}

// AddScalar adds the given scalar to each element
func (v Vector) AddScalar(f float64) {
	for i := range v.slice {
		v.slice[i] += f
	}
}

// Sub does the element wise substactions with the given vector
func (v Vector) Sub(other Vector) {
	for i, f := range other.slice {
		v.slice[i] -= f
	}
}

// Mul does the element wise multiplication with the other vector
func (v Vector) Mul(other Vector) {
	for i, f := range other.slice {
		v.slice[i] *= f
	}
}

// Len returns length of the vector
func (v Vector) Len() int {
	return len(v.slice)
}

// Scale scales each element by the given float
func (v Vector) Scale(f float64) {
	for i := range v.slice {
		v.slice[i] *= f
	}
}

// Pow applies the power function to each element with the given exponent
func (v Vector) Pow(f float64) {
	for i := range v.slice {
		v.slice[i] = math.Pow(v.slice[i], f)
	}
}

// Exp applies the exponent function to each element
func (v Vector) Exp() {
	for i := range v.slice {
		v.slice[i] = math.Exp(v.slice[i])
	}
}

// Sigmoid applies the sigmoid activation function to the elements
func (v Vector) Sigmoid() {
	v.Scale(-1)
	v.Exp()
	v.AddScalar(1)
	v.Pow(-1)
}

// SigmoidDer applies the sigmoid derivative function to the elements
func (v Vector) SigmoidDer() {
	v.Sigmoid()
	copy := Copy(v)
	copy.Pow(2)
	v.Sub(copy)
}

// Sum returns the sum of the elements
func (v Vector) Sum() float64 {
	sum := 0.
	for _, n := range v.slice {
		sum += n
	}

	return sum
}

// Ln applies the natural logorithm to the receiver
func (v Vector) Ln() {
	for i, n := range v.slice {
		v.slice[i] = math.Log(n)
	}
}

// Swap swaps the ith and jth elements
func (v Vector) Swap(i, j int) {
	v.slice[i], v.slice[j] = v.slice[j], v.slice[i]
}

// ReLU applies the ReLU activation function to the receiver
func (v Vector) ReLU() {
	for i, n := range v.slice {
		v.slice[i] = math.Max(n, 0)
	}
}

// ReLUDer applies the ReLU derivative to the receiver
func (v Vector) ReLUDer() {
	for i, n := range v.slice {
		if n > 0 {
			v.slice[i] = 1
		} else {
			v.slice[i] = 0
		}
	}
}

// String returns a string representation of the Vector
func (v Vector) String() string {
	return fmt.Sprintf("%v", v.slice)
}
