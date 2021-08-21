package main

import (
	"github.com/dimiro1/banner"
	"io/ioutil"
	"metadata-scraper/internal/handlers"
	"metadata-scraper/internal/options"
	"os"
	"strings"
)

func init() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	mso := options.GetMetadataScraperOptions()
	handlers.BingSearch(mso.Domain, mso.FileType)
	handlers.GoogleSearch(mso.Domain, mso.FileType)

	// TODO: Run seperately and collect the results in a channel
}
