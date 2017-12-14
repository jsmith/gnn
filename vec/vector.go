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
	Mul(other Vector)
	Exp()
	Pow(n Number)
	Sub(other Vector)
	Add(other Vector)
	AddScalar(n Number)
	Len() int
	Sigmoid()
	Sum() Number
	Copy() Vector
	Log()
}

// Creator provides context for the vector
type Creator interface {
	Make(i int) Vector
	Number(n float64) Number
	List(data []float64) NumberList
	Float64(n Number) float64
	Sub(v1, v2 Vector) Vector
	Copy(v Vector) Vector
}

type
