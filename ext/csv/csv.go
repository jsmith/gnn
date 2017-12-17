package csv

import (
	"encoding/csv"
	"io"
	"strings"
)

// Read returns a string matrix with the contents of the given csv file and ignores lines that contain the given values
func Read(bytes []byte, ignore ...string) [][]string {
	m := make(map[string]bool)
	for _, value := range ignore {
		m[value] = true
	}

	r := csv.NewReader(strings.NewReader(string(bytes)))
	var data [][]string
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		foundIgnore := false
		for _, item := range line {
			if m[item] {
				foundIgnore = true
				break
			}
		}
		if foundIgnore {
			continue
		}

		data = append(data, line)
	}

	return data
}
