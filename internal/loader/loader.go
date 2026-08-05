package loader

import (
	"GoHealthCheck/internal/logger"
	"GoHealthCheck/internal/output"
	"encoding/json"
	"net/url"
	"os"
	"strings"
	
)

func ValidateAndLoadUrls(path string) ([]string, error) {

	// Read content file
	data, err := readJFileContent(path)
	if err != nil {
		return []string{}, err
	}

	// Create Slice with file content
	urls := strings.Split(string(data), "\n")

	// Remove empty strings and invalid urls
	urls_filtered := []string{}
	for _, current_url := range urls {
		url_clean := strings.TrimSpace(current_url)

		if url_clean != "" && isValidHttpUrl(url_clean) {
			urls_filtered = append(urls_filtered, url_clean)
		}
	}

	if len(urls_filtered) == 0 {
		logger.WriteLog("No valid URL found")
	}

	return urls_filtered, nil
}

func ConvertJsonConfigFileToStruct(pathFile string) (output.CheckOptions, error) {
	data, err := readJFileContent(pathFile)
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

func readJFileContent(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func isValidHttpUrl(rawUrl string) bool {
	u, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return false
	}

	return u.Scheme != "" && u.Host != "" && (u.Scheme == "http" || u.Scheme == "https")
}