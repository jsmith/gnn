package vec

// Make makes a vector of the given size
func Make(size int) Vector {
	slice := make([]float64, size)
	return Vector{slice}
}

// Init initializes a vector is the given numbers
func Init(nums ...float64) Vector {
	return Vector{nums}
}

// Sub subtracts v1 from v2 and returns the result
func Sub(v1, v2 Vector) Vector {
	v := Copy(v1)
	v.Sub(v2)

	return v
}

// Copy returns a copy of the given vector
func Copy(from Vector) Vector {
	to := Make(from.Len())
	copy(to.slice, from.slice)

	return to
}

// Slice returns a slice of vector [from, to)
func Slice(v Vector, from, to int) Vector {
	return Vector{v.slice[from:to]}
}

// Unique returns a new vector containing only the unique elements of the given vector
func Unique(v Vector) Vector {
	unique := Make(0)
	for i := range v.slice {
		if !unique.Contains(v.At(i)) {
			unique.Append(v.At(i))
		}
	}

	return unique
}