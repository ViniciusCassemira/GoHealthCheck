package loader

import (
	"os"
	"strings"
	"encoding/json"
	"GoHealthCheck/internal/output"
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

func ConvertJsonConfigFileToStruct(pathFile string) (output.CheckOptions, error) {
	data, err := LoadJsonConfigFile(pathFile)
	if err != nil {
		return output.CheckOptions{}, err
	}

	var checkOptions output.CheckOptions

	err = json.Unmarshal(data, &checkOptions)
	if err != nil {
		return output.CheckOptions{}, err
	}

	return checkOptions, nil
}

func LoadJsonConfigFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
