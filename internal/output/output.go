package output

import (
	"fmt"
	"encoding/json"
	"os"
	"time"
)

func PrintResultInTerminal(result UrlCheckResult) {
	fmt.Printf("URL: %v\n", result.URL)
	fmt.Printf("StatusCode: %v\n", result.StatusCode)
	fmt.Printf("Protocol Version: %v\n", result.ProtocolVersion)
	fmt.Println("===============================================")
}

func SaveResultInJsonFile(result UrlCheckResultList, jsonPath string) error {

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	time := time.Now().Format("2006-01-02_15-04-05")

	err = os.WriteFile(jsonPath+"/"+time+".json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}