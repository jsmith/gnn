package csv

import (
	"encoding/csv"
	"bufio"
	"io"
	"os"
)

// Read returns a string matrix with the contents of the given csv file
func Read(file *os.File) [][]string {
	reader := csv.NewReader(bufio.NewReader(file))
	var data [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		data = append(data, line)
	}

	return data
}
