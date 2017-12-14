package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestSECost(t *testing.T) {
	mse := SE{}
	v1 := vec.Init([]float64{0, 1, 2})
	v2 := vec.Init([]float64{0, 0, 0})
	cost := mse.Cost(v1, v2)

	assert.Equal(t, (0. + 1. + 4.), cost)
}

func TestSEDef(t *testing.T) {
	mse := SE{}
	v1 := vec.Init([]float64{0, 1, 2})
	v2 := vec.Init([]float64{0, 0, 0})
	cost := mse.Der(v1, v2)

	assert.Equal(t, (0.+1.+2.)*2, cost)
}
