package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

// findLinks recursively finds all <a> tags in the HTML document and collects their href attributes.
func findLinks(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var link Link
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link.Href = attr.Val
				break
			}
		}
		link.Text = extractText(n)
		*links = append(*links, link)
	}
	// Recursively traverse child nodes.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinks(c, links)
	}
}

// extractText extracts all text content within a node.
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}
	var result string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += extractText(c) + " "
	}
	return strings.TrimSpace(result)
}

func main() {
	_, b, _, _ := runtime.Caller(0) // Get the path of this file
	basePath := filepath.Dir(b)     // Get the directory of this file

	// Open the JSON file
	file, err := os.ReadFile(filepath.Join(basePath, "ex4.html"))
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(strings.NewReader(string(file)))
	if err != nil {
		panic(err)
	}

	var links []Link
	findLinks(doc, &links)

	fmt.Println(links)
}
