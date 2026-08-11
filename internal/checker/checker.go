package checker

import (
	"GoHealthCheck/internal/logger"
	"GoHealthCheck/internal/output"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func CheckMultipleUrl(urls []string, options output.CheckOptions, pathJsonResultFile string) {
    var wg sync.WaitGroup
    ch := make(chan output.UrlCheckResult, len(urls))
    checkResultList := output.UrlCheckResultList{CheckOptions: options}
    executionTime := time.Now()

    for _, url := range urls {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
            ch <- CheckUrl(url, options)
			output.PrintProgressBarInTerminal(checkResultList.TotalUrls+1, len(urls))
        }(url)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for result := range ch {
        checkResultList.Results = append(checkResultList.Results, result)
        checkResultList.TotalUrls++
    }

    checkResultList.ExecutionTime = time.Since(executionTime).Seconds()
    output.SaveResultInJsonFile(checkResultList, pathJsonResultFile)
}

func CheckUrl(url string, options output.CheckOptions) output.UrlCheckResult {
    result := output.UrlCheckResult{URL: url}
    executionTime := time.Now()

    if options.CheckHttpInfo {
        result.HttpInfo = checkHttpInfo(url)
    }
    if options.CheckDnsInfo {
        result.DnsInfo = checkDnsInfo(url)
    }

    result.ExecutionTime = time.Since(executionTime).Seconds()
    return result
}

func checkHttpInfo(url string) (output.HttpInfo) {
	res, err := http.Get(url)
	if err != nil {
		logger.WriteLog(err.Error())
		return output.HttpInfo{}
	}
	defer res.Body.Close()

	return output.HttpInfo{
		ProtocolVersion: res.Proto,
		StatusCode:      res.StatusCode,
		ContentType:     res.Header.Get("Content-Type"),
	}
}

func checkDnsInfo(rawUrl string) (output.DnsInfo) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		logger.WriteLog(err.Error())
		return output.DnsInfo{}
	}
	hostname := parsedUrl.Hostname()

	aRecords, aaaaRecords := lookupAddressRecords(hostname)

	return output.DnsInfo{
		ARecords:     aRecords,
		AAAARecords:  aaaaRecords,
		CnameRecords: lookupCnameRecords(hostname),
		MxRecords:    lookupMxRecords(hostname),
		NsRecords:    lookupNsRecords(hostname),
		TxtRecords:   lookupTxtRecords(hostname),
	}
}

func lookupAddressRecords(hostname string) (aRecords, aaaaRecords []string) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		logger.WriteLog(err.Error())
		return nil, nil
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			aRecords = append(aRecords, ip.String())
		} else {
			aaaaRecords = append(aaaaRecords, ip.String())
		}
	}

	return aRecords, aaaaRecords
}

func lookupCnameRecords(hostname string) []string {
	cnameRecord, err := net.LookupCNAME(hostname)
	if err != nil {
		logger.WriteLog(err.Error())
		return nil
	}

	return []string{cnameRecord}
}

func lookupMxRecords(hostname string) []string {
	mxRecords, err := net.LookupMX(hostname)
	if err != nil {
		logger.WriteLog(err.Error())
		return nil
	}

	records := make([]string, 0, len(mxRecords))
	for _, mx := range mxRecords {
		records = append(records, mx.Host)
	}

	return records
}

func lookupNsRecords(hostname string) []string {
	nsRecords, err := net.LookupNS(hostname)
	if err != nil {
		logger.WriteLog(err.Error())
		return nil
	}

	records := make([]string, 0, len(nsRecords))
	for _, ns := range nsRecords {
		records = append(records, ns.Host)
	}

	return records
}

func lookupTxtRecords(hostname string) []string {
	txtRecords, err := net.LookupTXT(hostname)
	if err != nil {
		logger.WriteLog(err.Error())
		return nil
	}

	records := make([]string, 0, len(txtRecords))
	for _, txt := range txtRecords {
		records = append(records, txt)
	}

	return records
}