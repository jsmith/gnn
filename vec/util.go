package vec

// Make Make
func Make(size int) Vector {
	slice := make([]float64, size)
	return Vector{slice}
}

// Sub Sub
func Sub(v1, v2 Vector) Vector {
	v := Copy(v1)
	v.Sub(v2)

	return v
}

// Copy Copy
func Copy(from Vector) Vector {
	to := Make(from.Len())
	copy(to.Slice, from.Slice)

	return to
}
