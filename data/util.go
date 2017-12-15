package data

import "github.com/jacsmith21/gnn/mat"

// Init initializes a DataSet with the given data and labels
func Init(data, labels mat.Matrix) DataSet {
	return DataSet{data, labels}
}

// Copy Copy
func Copy(d DataSet) DataSet {
	data := mat.Copy(d.data)
	labels := mat.Copy(d.labels)
	return DataSet{
		data:   data,
		labels: labels,
	}
}
