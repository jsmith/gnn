package data

import "github.com/jacsmith21/gnn/mat"

// Copy Copy
func Copy(d DataSet) DataSet {
	data := mat.Copy(d.data)
	labels := mat.Copy(d.labels)
	return DataSet{
		data:   data,
		labels: labels,
	}
}
