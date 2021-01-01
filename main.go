package main

import (
	"flag"
	"metadata-searcher/handler"
)

var (
	domain, fileType *string
)

func init() {
	domain = flag.String("domain", "nytimes.com", "domain of the target to search")
	fileType = flag.String("fileType", "docx", "file type to search for")
	flag.Parse()
}

func main() {
	handler.BingSearch(*domain, *fileType)

	// TODO: Run handlers/google.go and handlers/bing.go seperately and collect the results in a channel
}