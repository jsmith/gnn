package csv

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	file, err := os.Open("test.csv") // just pass the file name
	if err != nil {
		panic(err)
	}

	data := Read(file)
	exp := [][]string{
		{"1", "1", "1"},
		{"2", "2", "2"},
	}

	assert.Equal(t, exp, data)
}