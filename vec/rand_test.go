package vec

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	v := Make(5)
	Rand(v, Gaussian)

	fmt.Println("V:", v)
}
