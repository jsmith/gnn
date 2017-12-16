package mat

import (
	"fmt"

	"github.com/jacsmith21/gnn/vec"
)

// Matrix Matrix
type Matrix struct {
	cols []vec.Vector
}

// At returns the element at the ith row and jth column
func (m Matrix) At(i, j int) float64 {
	return m.cols[j].At(i)
}

// Set sets the number at the ith row & jth column to the given number
func (m Matrix) Set(i, j int, f float64) {
	m.cols[j].Set(i, f)
}

// ColCount returns the Matrix column count
func (m Matrix) ColCount() int {
	return len(m.cols)
}

// RowCount returns the Matrix row count
func (m Matrix) RowCount() int {
	if len(m.cols) == 0 {
		return 0
	}

	return m.cols[0].Len()
}

// Col returns the ith column
func (m Matrix) Col(i int) vec.Vector {
	return m.cols[i]
}

// SwapCols swaps the ith and jth columns
func (m Matrix) SwapCols(i, j int) {
	m.cols[i], m.cols[j] = m.cols[j], m.cols[i]
}

// Scale scales every element by the given float
func (m Matrix) Scale(f float64) {
	for i := range m.cols {
		m.cols[i].Scale(f)
	}
}

// Transpose transposes the matrix
func (m *Matrix) Transpose() {
	t := Make(m.ColCount(), m.RowCount())
	for i := 0; i < m.RowCount(); i++ {
		for j := 0; j < m.ColCount(); j++ {
			t.Set(j, i, m.At(i, j))
		}
	}
	m.cols = t.cols
}

// AddCol adds the given vector to each column vector
func (m Matrix) AddCol(col vec.Vector) {
	for _, c := range m.cols {
		c.Add(col)
	}
}

// Sigmoid applies the sigmoid function to each element
func (m Matrix) Sigmoid() {
	for _, c := range m.cols {
		c.Sigmoid()
	}
}

// SigmoidDer applies the sigmoid derivative to each element
func (m Matrix) SigmoidDer() {
	for _, c := range m.cols {
		c.SigmoidDer()
	}
}

// ReLU applies the ReLU function to each element
func (m Matrix) ReLU() {
	for _, c := range m.cols {
		c.ReLU()
	}
}

// ReLUDer applies the ReLU derivative to each element
func (m Matrix) ReLUDer() {
	for _, c := range m.cols {
		c.ReLUDer()
	}
}

func (m Matrix) String() string {
	str := ""
	for i := 0; i < m.RowCount(); i++ {
		str += "[ "
		for _, c := range m.cols {
			str += fmt.Sprintf("%v ", c.At(i))
		}
		str += "]\n"
	}

	return str
}
