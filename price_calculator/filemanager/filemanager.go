package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("could not open file")
		fmt.Print(err)
		return nil, errors.New("failed to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("could not read file:")
		fmt.Print(err)
		file.Close()
		return nil, errors.New("failed to read file")
	}

	return lines, nil
}
