package vec64

import "github.com/jacsmith21/gnn/vec"

// Creator Creator
type Creator struct{}

// Number Number
func (c Creator) Number(n float64) vec.Number {
	return float32(n)
}

// Make Make
func (c Creator) Make(size int) vec.Vector {
	slice := make([]float64, size)
	return Vec64{slice}
}
