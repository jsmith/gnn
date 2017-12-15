package mat

import "github.com/jacsmith21/gnn/vec"

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

// Col returns the ith column
func (m Matrix) Col(i int) vec.Vector {
	return m.cols[i]
}

// SwapCols swaps the ith and jth columns
func (m Matrix) SwapCols(i, j int) {
	m.cols[i], m.cols[j] = m.cols[j], m.cols[i]
}
