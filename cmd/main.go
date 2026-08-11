package main

import (
	"GoHealthCheck/internal/checker"
	"GoHealthCheck/internal/loader"
	"GoHealthCheck/internal/logger"
	"GoHealthCheck/internal/output"
	"flag"
	"log"
)

func main() {
	
	json_config_file := flag.String("json-config-file", "", "path to the JSON config file")
	json_path := flag.String("json-path", "", "path to the directory where the results will be saved")
	url_file := flag.String("url-file", "", "path to the file containing URLs to check")
	flag.Parse()
	output.PrintMessageInTerminal("Starting...")

	if *json_config_file == "" {
		output.PrintMessageInTerminal("JSON config file path is required, using --json-config-file flag")
		logger.WriteLog("JSON config file path is required, using --json-config-file flag")
		log.Fatal("JSON config file path is required, using --json-config-file flag") 
	}
	if *json_path == "" { 
		output.PrintMessageInTerminal("JSON path is required, using --json-path flag")
		logger.WriteLog("JSON path is required, using --json-path flag")
		log.Fatal("JSON path is required, using --json-path flag") 
	}
	if *url_file == "" { 
		output.PrintMessageInTerminal("URL file path is required, using --url-file flag")
		logger.WriteLog("URL file path is required, using --url-file flag")
		log.Fatal("URL file path is required, using --url-file flag") 
	}

	checkOptions, err := loader.ConvertJsonConfigFileToStruct(*json_config_file)
	if err != nil {
		logger.WriteLog(err.Error())
		log.Fatal(err)
	}

	cfg	:= output.FlagConfig{
		UrlFilePath: *url_file,
		JsonPath: *json_path,
	}

	urls, err := loader.ValidateAndLoadUrls(cfg.UrlFilePath)
	if err != nil {
		logger.WriteLog(err.Error())
		log.Fatal(err)
	}

	checker.CheckMultipleUrl(urls, checkOptions, cfg.JsonPath)
}
