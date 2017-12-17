package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	data := Read([]byte("1,1,1\n2,2,2\n?,3,3"), "?")
	exp := [][]string{
		{"1", "1", "1"},
		{"2", "2", "2"},
	}

	assert.Equal(t, exp, data)
}
