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
