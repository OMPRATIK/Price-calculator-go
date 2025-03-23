package fileManager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filePath string) (*[] string, error) {
	file, error := os.Open(filePath)

	if error != nil {
		file.Close()
		return nil, errors.New("Failed to open the file!")
	}
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	error = scanner.Err()

	if error != nil {
		file.Close()
		return nil, errors.New("Failed to read line in file.")
	}

	file.Close()
	return &lines, nil
}