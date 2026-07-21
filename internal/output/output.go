package output

import (
	"fmt"
)

func PrintResultInTerminal(result CheckResult) {
	fmt.Printf("URL: %v\n", result.URL)
	fmt.Printf("StatusCode: %v\n", result.StatusCode)
	fmt.Printf("Protocol Version: %v\n", result.ProtocolVersion)
	fmt.Println("===============================================")
}
