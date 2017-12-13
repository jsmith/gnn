package vec64

import "github.com/jacsmith21/gnn/vec"

// Creator Creator
type Creator struct{}

// Number Number
func (c Creator) Number(n float64) vec.Number {
	return float64(n)
}

// List List
func (c Creator) List(data []float64) vec.NumberList {
	list := make([]float64, len(data))
	for i, n := range data {
		list[i] = float64(n)
	}

	return list
}

// Make Make
func (c Creator) Make(size int) vec.Vector {
	slice := make([]float64, size)
	return Vec64{slice}
}
