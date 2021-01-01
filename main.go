package main

import (
	"flag"
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
	bingSearch(*domain, *fileType)
	googleSearch(*domain, *fileType)

	// TODO: Run seperately and collect the results in a channel
}