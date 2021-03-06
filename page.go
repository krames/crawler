package main

import (
	"net/url"
	"strings"
)

type Page struct {
	source *url.URL
	Status int
	links  map[string]struct{}
	Depth  int
}

func NewPage(source string, depth int) (*Page, error) {
	u, err := url.Parse(source)
	if err != nil {
		return nil, err
	}

	return &Page{
		source: u,
		links:  make(map[string]struct{}),
		Depth:  depth,
	}, nil
}

func (p Page) Source() string {
	return p.source.String()
}

func (p *Page) AddLink(link string) {
	u, err := url.Parse(link)
	if err != nil {
		return
	}

	//Don't include anchor links on source page
	if u.Path == "" && (u.Host == "" || u.Host == p.source.Host) {
		return
	}

	//Don't include mobile versions
	parts := strings.Split(u.Host, ".")
	if len(parts) > 0 && strings.EqualFold(parts[0], "m") {
		return
	}

	if u.Host == "" {
		u.Host = p.source.Host
	}

	if u.Scheme == "" {
		u.Scheme = p.source.Scheme
	}

	p.links[u.String()] = struct{}{}
}

func (p Page) Links() []string {
	keys := make([]string, len(p.links))
	i := 0
	for k, _ := range p.links {
		keys[i] = k
		i += 1
	}
	return keys
}
