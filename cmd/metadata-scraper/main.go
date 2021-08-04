package main

import (
	"metadata-scraper/pkg/handlers"
	"metadata-scraper/pkg/options"
)

func main() {
	mso := options.GetMetadataScraperOptions()
	handlers.BingSearch(mso.Domain, mso.FileType)
	handlers.GoogleSearch(mso.Domain, mso.FileType)

	// TODO: Run seperately and collect the results in a channel
}
