package vec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	v := Make(5)
	assert.Equal(t, 5, v.Len())
}

func TestCreatorSub(t *testing.T) {
	v1 := Vector{[]float64{0, 1, 2, 3, 4}}
	v2 := Vector{[]float64{0, 1, 2, 3, 4}}
	vec := Sub(v1, v2)

	assert.Equal(t, 0., vec.At(0))
	assert.Equal(t, 0., vec.At(3))
}

func TestCopy(t *testing.T) {
	v1 := Vector{[]float64{0, 1, 2, 3, 4}}
	v2 := Copy(v1)

	v1.Set(0, 1.)
	assert.Equal(t, 0., v2.At(0))
	assert.Equal(t, 1., v2.At(1))
}
