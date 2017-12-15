package mocks

import "github.com/stretchr/testify/mock"

// Rander Rander
type Rander struct {
	count int
	mock.Mock
}

// Intn Intn
func (m *Rander) Intn(n int) int {
	args := m.Called(n)
	return args.Int(0)
}
