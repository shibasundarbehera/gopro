package utils

import (
	"encoding/json"
	"io"
	"os"
)

// LoadFromFile is a generic function to load and unmarshal JSON into a slice of any type.
func LoadFromFile[T any](filename string) ([]T, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result []T
	err = json.Unmarshal(data, &result)
	return result, err
}

// SaveToFile is a generic function to marshal and save a slice of any type to a JSON file.
func SaveToFile[T any](filename string, data []T) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
