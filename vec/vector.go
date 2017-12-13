package vec

// Number Number
type Number interface{}

// NumberList NumberList
type NumberList interface{}

// Vector a vector
type Vector interface {
	Creator() Creator
	At(i int) Number
	Set(i int, n Number)
	SetData(data NumberList)
	Scale(n Number)
	Exp()
	Pow(n Number)
	Sub(other Vector)
	Add(other Vector)
	AddScalar(n Number)
	Len() int
	Sigmoid()
}

// Creator provides context for the vector
type Creator interface {
	Make(i int) Vector
	Number(n float64) Number
	List(data []float64) NumberList
}
