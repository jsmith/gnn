package data

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

var s Sample

func initSample() {
	s = Sample{vec.Init(1, 2, 3, 4), vec.Init(1)}
}

func TestSampleLabel(t *testing.T) {
	initSample()
	assert.Equal(t, 1., s.Label(0))
}

func TestSampleData(t *testing.T) {
	initSample()
	assert.Equal(t, 3., s.DataValue(2))
}
