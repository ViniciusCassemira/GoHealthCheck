package checker

import (
	"GoHealthCheck/internal/output"
	"net/http"
)

func CheckMultipleUrl(urls []string, options output.CheckOptions) error {

	for _, url := range urls {
		checkResult, err := CheckUrl(url, options)
		if err != nil {
			return err
		}

		output.PrintResultInTerminal(checkResult)
	}

	return nil
}

func CheckUrl(url string, options output.CheckOptions) (output.CheckResult, error) {
	result := output.CheckResult{}
	result.URL = url

	if options.CheckStatusCode {
		statusCode, err := checkStatusCode(url)
		if err != nil {
			return result, err
		}

		result.StatusCode = statusCode
	}

	if options.CheckProtocolversion {
		protocolVersion, err := checkProtocolVersion(url)
		if err != nil {
			return result, err
		}

		result.ProtocolVersion = protocolVersion
	}

	return result, nil
}

func checkStatusCode(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

func checkProtocolVersion(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return res.Proto, nil
}
