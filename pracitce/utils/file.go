package utils

import (
	"encoding/json"
	"os"
	"strings"
)

func LoadFromFile[T any](filename string, dest *[]T) {
	data, err := os.ReadFile(filename)
	if err == nil {
		_ = json.Unmarshal(data, dest)
	}
}

func SaveToFile[T any](filename string, data []T) {
	bytes, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(filename, bytes, 0644)
}

func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))

}
