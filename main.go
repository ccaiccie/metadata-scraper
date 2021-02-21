package main

import (
	flag "github.com/spf13/pflag"
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
	bingSearch(domain, fileType)
	googleSearch(domain, fileType)

	// TODO: Run seperately and collect the results in a channel
}