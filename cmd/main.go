package main

import (
	"GoHealthCheck/internal/checker"
	"GoHealthCheck/internal/loader"
	"GoHealthCheck/internal/output"
	"flag"
	"log"
)

func main() {

	checkOptions := output.CheckOptions{
		CheckStatusCode:      true,
		CheckProtocolversion: true,
	}

	url_file := flag.String("url-file", "", "path to the file containing URLs to check")
	json_path := flag.String("json-path", "", "path to the directory where the results will be saved")
	flag.Parse()
	if *url_file == "" { log.Fatal("URL file path is required, using -url-file flag") }
	if *json_path == "" { log.Fatal("JSON path is required, using -json-path flag") }

	cfg	:= output.FlagConfig{
		UrlFilePath: *url_file,
		JsonPath: *json_path,
	}

	urls, err := loader.LoadUrls(cfg.UrlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = checker.CheckMultipleUrl(urls, checkOptions, cfg.JsonPath)
	if err != nil {
		log.Fatal(err)
	}
}
