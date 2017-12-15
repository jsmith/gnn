package rander

import (
	"fmt"
	"testing"

	"github.com/jacsmith21/gnn/mat"
)

func TestRand(t *testing.T) {
	m := mat.Make(2, 2)
	Rand(m, Gaussian)

	fmt.Println("matrix:", m)
}
