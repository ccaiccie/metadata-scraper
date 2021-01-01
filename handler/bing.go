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
	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(query))
	res, err := http.Get(search)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	// TODO: Fix, not working
	s := "html body div#b_content ol#b_results li.b_algo h2"
	doc.Find(s).Each(utility.Handler)
}