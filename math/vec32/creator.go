package vec32

import "github.com/jacsmith21/gnn/math"

type Creator32 struct{}

// Make makes a Vec32 of length size
func (c Creator32) Number(n float64) math.Number {
	return float32(n)
}

// Make makes a Vec32 of length size
func (c Creator32) Make(size int) math.Vector {
	slice := make([]float32, size)
	return Vec32{slice}
}
