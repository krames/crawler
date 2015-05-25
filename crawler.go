package main

import (
	"fmt"
	"net/http"

	"github.com/krames/crawler/domain"

	"golang.org/x/net/html"
)

func main() {
	page, _ := domain.NewPage("http://reddit.com/")

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

	for _, v := range page.Links() {
		fmt.Println(v)
	}
}
