package data

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	data := mat.Init(vec.Init(1, 2, 3))
	labels := mat.Init(vec.Init(1, 2, 3))
	dataset := Init(data, labels)
	assert.Equal(t, 1, dataset.labels.ColCount())
	assert.Equal(t, 1, dataset.data.ColCount())
}

func TestCopy(t *testing.T) {
	initDataSet()
	c := Copy(d)
	d.labels.Col(0).Set(0, 500)
	d.data.Col(0).Set(0, 500)

	assert.Equal(t, 4, c.labels.ColCount())
	assert.Equal(t, 1, c.labels.Col(0).Len())
	assert.Equal(t, 0., c.labels.Col(0).At(0))

	assert.Equal(t, 4, c.data.ColCount())
	assert.Equal(t, 3, c.data.Col(0).Len())
	assert.Equal(t, 0., c.data.Col(0).At(0))
}
