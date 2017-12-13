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

// Sub Sub
func (c Creator) Sub(v1, v2 vec.Vector) vec.Vector {
	v := c.Make(v1.Len()).(Vec64)
	copy(v.slice, v1.(Vec64).slice)
	v.Sub(v2)

	return v
}

// Sub Sub
func (c Creator) Copy(v vec.Vector) vec.Vector {
	new := c.Make(v.Len()).(Vec64)
	copy(new.slice, v.(Vec64).slice)

	return new
}

// Float64 Float64
func (c Creator) Float64(n vec.Number) float64 {
	return n.(float64)
}
