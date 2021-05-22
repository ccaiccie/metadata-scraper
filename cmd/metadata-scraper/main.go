package main

import (
	flag "github.com/spf13/pflag"
	"metadata-scraper/pkg/handlers"
)

var (
	domain, fileType string
)

func init() {
	flag.StringVar(&domain, "domain", "nytimes.com", "domain of the target to search")
	flag.StringVar(&fileType, "fileType", "docx", "file type to search for")
	flag.Parse()
}

func main() {
	handlers.BingSearch(domain, fileType)
	handlers.GoogleSearch(domain, fileType)

	// TODO: Run seperately and collect the results in a channel
}