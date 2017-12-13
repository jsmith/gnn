package math

// Number a number
type Number interface{}

// Vector a vector
type Vector interface {
	Creator() Creator
	At(i int) Number
	Set(i int, n Number)
	Sub(other Vector)
	Add(other Vector)
	Len() int
}

// Creator provides context for the vector
type Creator interface {
	Make(i int) Vector
	Number(n float64) Number
}
