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
