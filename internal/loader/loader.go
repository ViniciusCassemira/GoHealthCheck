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

	// Remove empty strings
	urls_filtered := []string{}
	for _, url := range urls {
		url = strings.TrimSpace(url)
		if url != "" {
			urls_filtered = append(urls_filtered, url)
		}
	}

	return urls_filtered, nil
}
