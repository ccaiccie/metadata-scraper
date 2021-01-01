package utility

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"metadata-searcher/metadata"
	"net/http"
)

func Handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, url)
	res, err := http.Get(url)
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

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf("%21s %s - %s %s\n", cp.Creator, cp.LastModifiedBy, ap.Application, ap.GetMajorVersion())
}