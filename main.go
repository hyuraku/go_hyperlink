package main

import (
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
	"regexp"

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
	responseBody := string(byteArray)
	re := regexp.MustCompile("<title>(.*)</title>")
	title := re.FindStringSubmatch(responseBody)[1]

	results := hyperlink.Hyperlink(url, title)

	for key, value := range results {
		fmt.Println(key, value)
	}
}
