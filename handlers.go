package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var googleCounter, bingCounter int

func handle(i int, s *goquery.Selection) {
	href, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, href)
	res, err := http.Get(href)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := newProperties(r)
	if err != nil {
		return
	}

	log.Printf("\nCreator = %s\nLast Modified By = %s\nApplication = %s %s\n\n\n", cp.Creator, cp.LastModifiedBy,
		ap.Application, ap.getMajorVersion())
}

func googleSearch(domain, fileType string) {
	query := fmt.Sprintf("as_sitesearch=%s&as_filetype=%s", domain, fileType)
	search := fmt.Sprintf("https://www.google.com/search?%s", query)
	log.Printf("full search url for google = %s\n", search)

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

	s := "html body#gsr div#main div#cnt div.lEXIrb div#rcnt div.D6j0vc div#center_col div#res div#search div div#rso div.g div.rc div.yuRUbf"
	doc.Find(s).Each(func(i int, selection *goquery.Selection) {
		googleCounter++
		handle(i, selection)
	})
	log.Printf("found %d items with specified search parameters on Google search engine!\n", googleCounter)
}

func bingSearch(domain, fileType string) {
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
	doc.Find(s).Each(func(i int, selection *goquery.Selection) {
		bingCounter++
		handle(i, selection)
	})
	log.Printf("found %d items with specified search parameters on Bing search engine!\n", bingCounter)
}