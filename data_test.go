package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestDataSetLen(t *testing.T) {
	x := vec.Init([]float64{1, 1, 1})
	y := vec.Init([]float64{1})
	dataset := DataSet{
		data:   []vec.Vector{x},
		labels: y,
	}

	assert.Equal(t, 1, dataset.SampleCount())
}
