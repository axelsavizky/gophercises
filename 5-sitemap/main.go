package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	linkparser "gophercises/4-linkparser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	var rootpage string
	flag.StringVar(&rootpage, "page", "https://courses.calhoun.io/", "a page to generate sitemap")

	flag.Parse()

	visited := make(map[string]bool)
	visited[rootpage] = true

	toVisit := Queue{rootpage}

	rootURL, _ := url.Parse(rootpage)

	var totalLinks []string

	var currentURL string
	var newURL *url.URL
	for !toVisit.IsEmpty() {
		currentURL = toVisit.Pop()
		resp, err := http.Get(currentURL)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		links := linkparser.GetLinksFromPage(body)

		for _, link := range links {
			newURL, _ = rootURL.Parse(link.Href)
			if newURL.Hostname() != rootURL.Hostname() {
				continue
			}

			if _, ok := visited[newURL.String()]; ok {
				continue
			}

			toVisit.Push(newURL.String())
			visited[newURL.String()] = true
			totalLinks = append(totalLinks, newURL.String())
		}
	}

	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, link := range totalLinks {
		toXml.Urls = append(toXml.Urls, loc{link})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}

type Queue []string

func (q *Queue) Push(x string) {
	*q = append(*q, x)
}

func (q *Queue) Pop() string {
	first := (*q)[0]
	*q = (*q)[1:]

	return first
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
