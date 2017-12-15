package mat

import "github.com/jacsmith21/gnn/vec"

// Init initilizes a Matrix with column vectors
func Init(vecs ...vec.Vector) Matrix {
	return Matrix{vecs}
}

// Slice returns a slice of the Matrix
func Slice(m Matrix, from, to int) Matrix {
	return Matrix{m.cols[from:to]}
}
