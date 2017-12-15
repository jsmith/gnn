package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
