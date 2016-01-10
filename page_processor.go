package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type PageProcessor struct {
	Source string
	Output chan Page
	depth  int
}

func (pp PageProcessor) Process() {
	page, _ := NewPage(pp.Source, pp.depth)

	resp, err := http.Get(page.Source())
	if err != nil {
		fmt.Println(err)
		return
	}
	page.Status = resp.StatusCode

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					page.AddLink(attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	pp.Output <- *page
}
