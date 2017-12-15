package gnn

import (
	"testing"

	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestSECost(t *testing.T) {
	se := SE{}
	exp := []vec.Vector{vec.Init([]float64{0, 1, 2})}
	act := []vec.Vector{vec.Init([]float64{0, 0, 0})}
	cost := se.Cost(exp, act)

	assert.Equal(t, 1, cost.Len())
	assert.Equal(t, (0. + 1. + 4.), cost.At(0))
}

func TestSEDef(t *testing.T) {
	mse := SE{}
	exp := []vec.Vector{vec.Init([]float64{0, 1, 2})}
	act := []vec.Vector{vec.Init([]float64{0, 0, 0})}
	cost := mse.Der(exp, act)

	assert.Equal(t, 1, cost.Len())
	assert.Equal(t, -(0.+1.+2.)*2, cost.At(0))
}
