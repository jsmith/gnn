package gnn

import "github.com/jacsmith21/gnn/vec"

// Cost Cost
type Cost interface {
	Cost(expected, actual vec.Vector) float64
}

// MSE MSE
type MSE struct{}

// Cost Cost
func (mse MSE) Cost(expected, actual vec.Vector) float64 {
	c := expected.Creator()
	diff := c.Sub(expected, actual)
	diff.Pow(2)

	n := diff.Sum()
	f := c.Float64(n)
	f /= float64(diff.Len())

	return f
}

// Entropy Entropy
type Entropy struct{}

// Cost Cost
func (e Entropy) Cost(expected, actual vec.Vector) float64 {
	negOne := expected.Creator().Number(-1)
	expectedCopy := expected.Copy()
	actualCopy := actual.Copy()

	actual.Log()
	expected.Mul(actual)
	cost := expected.Sum()

	expected = expectedCopy
	actual = actualCopy

	expected.Scale(negOne)
	expected.AddScalar(negOne)
	actual.Scale(negOne)
	actual.AddScalar(negOne)
	actual.Log()
	expected.Mul(actual)
	cost.Add(expected.Sum())

	return c.Float64(cost)
}
