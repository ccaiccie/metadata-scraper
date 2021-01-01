package handler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"metadata-searcher/utility"
	"net/http"
	"net/url"
)

func BingSearch(domain, fileType string) {
	query := fmt.Sprintf("site:%s && filetype:%s && instreamset:(url title):%s", domain, fileType, fileType)
	search := fmt.Sprintf("https://www.bing.com/search?q=%s", url.QueryEscape(query))
	log.Printf("full search url for bing = %s\n", search)

	client := &http.Client{}
	req, err := http.NewRequest("GET", search, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:84.0) Gecko/20100101 Firefox/84.0")


	res, err := client.Do(req)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	s := "html body #b_content main ol#b_results li.b_algo h2"
	// doc.Find(s).Each(utility.Handler)
	log.Printf("found %d items with specified search parameters!\n", len(s))
	doc.Find(s).Each(utility.Handler)
}