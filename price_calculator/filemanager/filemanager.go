package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	// first create an encoder for the file, and then encode the data for it.

	encoder := json.NewEncoder(file) // directly writes data on hard drive or web network
	// json.marshal() uses ram to process data and then writes to file, use this if you need to store json data in a variable
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert file to json")
	}

	file.Close()
	return nil
}

func New(inputpath string, outputpath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputpath,
		OutputFilePath: outputpath,
	}
}
