package main

import (
	"golang.org/x/net/html"

	"fmt"
	"net/http"
)

func main() {
	links := map[string]struct{}{}

	resp, err := http.Get("http://bbq.kylerames.com/")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if _, ok := links[attr.Val]; !ok {
						links[attr.Val] = struct{}{}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for v := range links {
		fmt.Println(v)
	}
}
