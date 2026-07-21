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

	path := flag.String("path", "", "path to the file containing URLs to check")
	flag.Parse()

	if *path == "" {
		log.Fatal("path is required")
	}

	urls, err := loader.LoadUrls(*path)
	if err != nil {
		log.Fatal(err)
	}

	err = checker.CheckMultipleUrl(urls, checkOptions)
	if err != nil {
		log.Fatal(err)
	}
}
