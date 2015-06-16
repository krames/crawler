package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/krames/crawler/domain"
	"golang.org/x/net/html"
)

type Worker interface {
	Process()
}

type WorkerPool struct {
	workers chan Worker
	wg      sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	p := WorkerPool{
		workers: make(chan Worker),
	}

	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			for w := range p.workers {
				w.Process()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (wp *WorkerPool) Do(w Worker) {
	wp.workers <- w
}

func (p *WorkerPool) Shutdown() {
	close(p.workers)
	p.wg.Wait()
}

type PageProcessor struct {
	Source string
	Output chan domain.Page
}

func (pp PageProcessor) Process() {
	page, _ := domain.NewPage(pp.Source)

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
