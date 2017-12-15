package mat

import (
	"github.com/jacsmith21/gnn/vec"
)

// Make initilizes a matrix with i rows and j cols
func Make(i, j int) Matrix {
	vecs := make([]vec.Vector, j)
	for k := 0; k < j; k++ {
		vecs[k] = vec.Make(i)
	}

	return Matrix{vecs}
}

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

func Mul(m1, m2 Matrix) Matrix {
	res := Make(m1.RowCount(), m2.ColCount())
	for i := 0; i < m1.RowCount(); i++ {
		for j := 0; j < m2.ColCount(); j++ {
			sum := 0.
			for c := 0; c < m1.ColCount(); c++ {
				sum += m1.At(i, c) * m2.At(i, j)
			}

			res.Set(i, j, sum)
		}
	}

	return res
}
