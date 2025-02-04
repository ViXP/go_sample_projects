package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (manager *FileManager) ReadLines() ([]string, error) {
	var unparsedStrings []string
	file, err := os.Open(manager.InputPath)

	if err != nil {
		file.Close()
		return nil, fmt.Errorf("file %v is not found", manager.InputPath)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		readLine := scanner.Text()

		if scanner.Err() != nil {
			file.Close()
			return nil, fmt.Errorf("unparsable value in %v", manager.InputPath)
		}

		unparsedStrings = append(unparsedStrings, readLine)
	}
	file.Close()
	return unparsedStrings, nil
}

func (manager *FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(manager.OutputPath)

	if err != nil {
		file.Close()
		return fmt.Errorf("failed to create %v", manager.OutputPath)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return fmt.Errorf("data is not JSON-able")
	}

	file.Close()
	return nil
}
