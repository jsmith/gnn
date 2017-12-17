package data

import (
	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
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
		row := m.Row(i)
		unique := vec.Unique(row)
		rows := mat.Make(unique.Len(), row.Len())
		for k := 0; k < unique.Len(); k++ {
			indices := row.Indices(unique.At(k))
			for j := 0; j < len(indices); j++ {
				rows.Set(k, indices[j], 1)
			}
		}

		for j := 0; j < rows.RowCount(); j++ {
			hot.Append(rows.Row(j), 1)
		}
	}

	return hot
}
