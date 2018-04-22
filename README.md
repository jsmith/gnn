[![Build Status](https://travis-ci.org/jacsmith21/gnn.png?branch=master)](https://travis-ci.org/jacsmith21/gnn)
# GNN
Golang Neural Network (GNN) framework! GNN was built for educational purposes. Currently this framework is able to implement simple fully connected networks.

# Example
```go
net := Net{
  NewFC(2, 4),
  &ReLU{},
  InitFC(4, 1),
  &Sigmoid{},
}

trainer = Trainer{
  Net:          net,
  Cost:         SE{}, // Standard Error Loss
  LearningRate: 0.1,
  Epochs:       10000,
  BatchSize:    4,
}

xor = data.Init(
  mat.InitRows(
    vec.Init(0, 0, 1, 1),
    vec.Init(0, 1, 1, 0),
  ),
  mat.InitRows(
    vec.Init(0, 1, 1, 0),
  ),
)


trainer.Train(xor)
predictions := trainer.Predict(xor.Data())
```

See `trainer_test.go:TestTrainer` for a worker example.


*[anynet](https://github.com/unixpickle/anynet) heavily used as reference*
