package checker

import (
	"GoHealthCheck/internal/logger"
	"GoHealthCheck/internal/output"
	"net/http"
	"time"
	"sync"
)

func CheckMultipleUrl(urls []string, options output.CheckOptions, pathJsonResultFile string) {
	var wg sync.WaitGroup
	ch := make(chan output.UrlCheckResult, len(urls))
	
	checkResultList := output.UrlCheckResultList{}
	execution_time := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go func(){
			url := url
			ch <- CheckUrl(url, options, &wg)
		}()
	}
	wg.Wait()

	for i:= 0; i < len(urls); i++ {
		logger.WriteLog("Checking URL: " + urls[i])
		checkResult := <-ch
		checkResultList.Results = append(checkResultList.Results, checkResult)
		checkResultList.TotalUrls++
		//output.PrintResultInTerminal(checkResult)
	}

	checkResultList.ExecutionTimeInSeconds = time.Since(execution_time).Seconds()
	output.SaveResultInJsonFile(checkResultList, pathJsonResultFile)
}

func CheckUrl(url string, options output.CheckOptions, wg *sync.WaitGroup) (output.UrlCheckResult) {
	defer wg.Done()
	result := output.UrlCheckResult{}
	result.URL = url
	execution_time := time.Now()

	if options.CheckStatusCode {
		result.StatusCode = checkStatusCode(url)
	}

	if options.CheckProtocolversion {
		result.ProtocolVersion = checkProtocolVersion(url)
	}

	result.ExecutionTimeInSeconds = time.Since(execution_time).Seconds()

	return result
}

func checkStatusCode(url string) (int) {
	res, err := http.Get(url)
	if err != nil {
		logger.WriteLog(err.Error())
		return 0
	}
	defer res.Body.Close()

	return res.StatusCode
}

func checkProtocolVersion(url string) (string) {
	res, err := http.Get(url)
	if err != nil {
		logger.WriteLog(err.Error())
		return ""
	}
	defer res.Body.Close()

	return res.Proto
}
