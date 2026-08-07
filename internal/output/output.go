package output

import (
	"encoding/json"
	"os"
	"time"
	"fmt"
)

func SaveResultInJsonFile(result UrlCheckResultList, jsonPath string) error {

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	time := time.Now().Format("2006-01-02_15-04-05")

	err = os.WriteFile(jsonPath+"/"+time+".json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func PrintProgressBarInTerminal(current, total int) {
	percentage := float64(current) / float64(total) * 100
	fmt.Printf("\rProgress: %d/%d (%.2f%%)", current, total, percentage)
}

func PrintMessageInTerminal(message string) {
	fmt.Println(message)
}