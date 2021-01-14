package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
)

// Urlset is the outermost XML tag
type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []URL    `xml:"url"`
}

// URL is contained within Urlset
type URL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	Priority   string   `xml:"priority"`
}

func main() {

	parser := argparse.NewParser("", "Grab all URLs from a provided sitemap.xml url. E.g https://example.com/sitemap.xml")
	u := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL to sitemap.xml"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	resp, err := http.Get(*u)
	if err != nil {
		panic("Could not fetch url")
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	var urlset Urlset
	xml.Unmarshal(content, &urlset)

	for i := 0; i < len(urlset.Urls); i++ {
		fmt.Println(urlset.Urls[i].Loc)
	}
}
