package data

import (
	"testing"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Rander is a mock implementation of the Rander interface
type Rander struct {
	mock.Mock
}

// Intn is the mock implementation of Intn
func (m *Rander) Intn(n int) int {
	args := m.Called(n)
	return args.Int(0)
}

var d DataSet

func initDataSet() {
	d = DataSet{
		data: mat.InitCols(
			vec.Init(0, 0, 0),
			vec.Init(1, 1, 1),
			vec.Init(2, 2, 2),
			vec.Init(3, 3, 3),
		),
		labels: mat.InitRows(
			vec.Init(0, 1, 2, 3),
		),
	}
}

func TestSampleCount(t *testing.T) {
	initDataSet()
	assert.Equal(t, 4, d.SampleCount())
}

func TestShuffle(t *testing.T) {
	initDataSet()

	rander := new(Rander)
	rander.On("Intn", 2).Return(2)
	rander.On("Intn", 3).Return(3)
	rander.On("Intn", 4).Return(0)
	d.Shuffle(rander)

	assert.Equal(t, 2., d.Sample(3).Label(0))
	assert.Equal(t, 2., d.Sample(3).DataValue(0))
	assert.Equal(t, 3., d.Sample(0).Label(0))
	assert.Equal(t, 3., d.Sample(0).DataValue(0))
	rander.AssertExpectations(t)
}

func TestGenerateBatches(t *testing.T) {
	initDataSet()
	batches := d.GenerateBatches(3)

	assert.Equal(t, 3, batches[0].SampleCount())
	assert.Equal(t, 0., batches[0].Sample(0).Label(0))
	assert.Equal(t, 1., batches[0].Sample(1).Label(0))

	assert.Equal(t, 1, batches[1].SampleCount())
	assert.Equal(t, 3., batches[1].Sample(0).Label(0))
}

func TestGenerateEmptyBatches(t *testing.T) {
	initDataSet()
	assert.Panics(t, func() { d.GenerateBatches(0) })
}

func TestLabels(t *testing.T) {
	labels := d.Labels()
	assert.Equal(t, 4, labels.ColCount())
}

func TestData(t *testing.T) {
	data := d.Data()
	assert.Equal(t, 4, data.ColCount())
}

func TestAttributeCount(t *testing.T) {
	initDataSet()
	assert.Equal(t, 4, d.AttributeCount())
}
