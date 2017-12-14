package gnn

// LearningRate LearningRate
type LearningRate float64

// Trainer Trainer
type Trainer struct {
	Net          Net
	Cost         Cost
	LearningRate LearningRate
	BatchSize    int
	Epochs       int
	Status       func(cost float64)
}

// Train Train
func (t Trainer) Train(d DataSet) {
	for i := 1; i <= t.Epochs; i++ {

	}
}
