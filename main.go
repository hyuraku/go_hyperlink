package main

import (
	"bytes"
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"

	"example.com/hyperlink"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		panic("write the url")
	}
	resp, err := http.Get(url)
  if err != nil {
		panic("check the url you write")
	}
	defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)

	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(byteArray)

	bReader := bytes.NewReader(byteArray)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	title := doc.Find("title").Text()

	results := hyperlink.Hyperlink(url, title)

	for _, value := range results {
		fmt.Println(value)
	}
}
