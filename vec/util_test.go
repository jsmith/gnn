package vec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	v := Make(5)
	assert.Equal(t, 5, v.Len())
}

func TestInit(t *testing.T) {
	v := Init(0, 1, 2, 3, 4)

	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 3., v.At(3))
}

func TestCreatorSub(t *testing.T) {
	v1 := Init(0, 1, 2, 3, 4)
	v2 := Init(0, 1, 2, 3, 4)
	vec := Sub(v1, v2)

	assert.Equal(t, 0., vec.At(0))
	assert.Equal(t, 0., vec.At(3))
}

func TestCopy(t *testing.T) {
	v1 := Init(0, 1, 2, 3, 4)
	v2 := Copy(v1)

	v1.Set(0, 1.)
	assert.Equal(t, 0., v2.At(0))
	assert.Equal(t, 1., v2.At(1))
}

func TestSlice(t *testing.T) {
	initVector()
	v := Slice(v, 0, 3)
	assert.Equal(t, 3, v.Len())
	assert.Equal(t, 0., v.At(0))
	assert.Equal(t, 1., v.At(1))
}
