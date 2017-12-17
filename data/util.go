package data

import (
	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"os"
	"encoding/csv"
	"bufio"
	"io"
)

// Init initializes a data set with the given data and labels
func Init(data, labels mat.Matrix) DataSet {
	return DataSet{data, labels}
}

// Copy returns a copy of the given data set
func Copy(d DataSet) DataSet {
	data := mat.Copy(d.data)
	labels := mat.Copy(d.labels)
	return DataSet{
		data:   data,
		labels: labels,
	}
}

// OneHot returns the one-hot encoding of the given matrix
func OneHot(m mat.Matrix) mat.Matrix {
	hot := mat.Make(0,m.ColCount())
	for i := 0; i < m.RowCount(); i++ {
		col := m.Row(i)
		unique := vec.Unique(col)
		cols := mat.Make(unique.Len(), col.Len())
		for j := 0; j < col.Len(); j++ {
			index := unique.Index(col.At(i))
			cols.Set(index, i, 1)
		}

		for j := 0; j < cols.ColCount(); j++ {
			hot.Append(cols.Col(i), 1)
		}
	}

	return hot
}
