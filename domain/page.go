package domain

type Page struct {
	Source string
	Status int
	links  map[string]struct{}
}

func NewPage(source string) Page {
	return Page{
		Source: source,
		links:  make(map[string]struct{}),
	}
}

func (p *Page) AddLink(link string) {
	p.links[link] = struct{}{}
}

func (p *Page) Links() []string {
	keys := make([]string, len(p.links))
	i := 0
	for k, _ := range p.links {
		keys[i] = k
		i += 1
	}
	return keys
}
