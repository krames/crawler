package main

import (
	"fmt"

	"github.com/krames/crawler/domain"
)

type Dispatcher struct {
	pool         *WorkerPool
	output       chan domain.Page
	crawledPages map[string]struct{}
}

func NewDispatcher(pool *WorkerPool, outputChan chan domain.Page) Dispatcher {
	d := Dispatcher{
		pool:         pool,
		output:       outputChan,
		crawledPages: make(map[string]struct{}),
	}

	go func() {
		for page := range d.output {
			fmt.Println("Processing: " + page.Source())
			go func() {
				for _, v := range page.Links() {
					if _, ok := d.crawledPages[v]; !ok {
						d.dispatch(v)
					}
				}
			}()
		}
	}()

	return d
}

func (d Dispatcher) dispatch(url string) {
	fmt.Println("DISPATCH: " + url)
	pp := PageProcessor{
		Source: url,
		Output: d.output,
	}
	d.pool.Do(pp)
}

func (d Dispatcher) Start(url string) {
	d.dispatch(url)
}
