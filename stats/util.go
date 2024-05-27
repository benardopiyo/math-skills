package stats

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func DataReader(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64

	// reads each line from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line) // trims leading and trailing whitespace
		if line == "" {
			continue // skip empty lines
		}
		num, err := strconv.ParseFloat(line, 64) // parses the trimmed line as a float64 value
		if err != nil {
			return nil, err
		}
		data = append(data, num) // appends the parsed value to the data slice
	}

	if err := scanner.Err(); err != nil { // handle parsing errors
		return nil, err
	}

	return data, nil
}
