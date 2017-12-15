package mat

import "github.com/jacsmith21/gnn/vec"

// Init initilizes a Matrix with column vectors
func Init(vecs ...vec.Vector) Matrix {
	return Matrix{vecs}
}

// Copy returns a copy of the given Matrix
func Copy(from Matrix) Matrix {
	cols := make([]vec.Vector, len(from.cols))
	for i, c := range from.cols {
		cols[i] = vec.Copy(c)
	}

	return Matrix{cols}
}

// Slice returns a slice of the Matrix
func Slice(m Matrix, from, to int) Matrix {
	return Matrix{m.cols[from:to]}
}
