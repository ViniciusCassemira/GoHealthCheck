package loader

import (
	"os"
	"strings"
)

func LoadUrls(path string) ([]string, error) {

	// Read content file
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	// Create Slice with file content
	urls := strings.Split(string(data), "\n")

	return urls, nil
}
