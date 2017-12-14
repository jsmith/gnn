package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestDataSetLen(t *testing.T) {
	v := vec.Vector{Slice: []float64{1, 1, 1}}
	dataset := DataSet{
		{v, 1},
	}

	assert.Equal(t, 1, dataset.Len())
}
