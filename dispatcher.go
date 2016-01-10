package main

import (
	"fmt"
)

type Dispatcher struct {
	pool            *WorkerPool
	output          chan Page
	crawledPages    map[string]struct{}
	dispatchedPages map[string]struct{}
	depth           int
}

func NewDispatcher(pool *WorkerPool, outputChan chan Page, depth int) Dispatcher {
	d := Dispatcher{
		pool:            pool,
		output:          outputChan,
		crawledPages:    make(map[string]struct{}),
		dispatchedPages: make(map[string]struct{}),
		depth:           depth,
	}

	go func() {
		for page := range d.output {
			//			fmt.Println("Processing: " + page.Source())
			d.printStats()

			// delete incoming page from dispatched list
			delete(d.dispatchedPages, page.Source())

			// add incoming page to crawled list
			d.crawledPages[page.Source()] = struct{}{}

			// If we have crawled to the appropriate depth do not continue processing
			if page.Depth < 1 {
				//	fmt.Println("Done")
				continue
			}

			// Process links for page
			for _, link := range page.Links() {
				// Do not process page if we have previously visited
				if _, ok := d.crawledPages[link]; !ok {
					//Do not process if this page has been dispatched
					if _, ok := d.dispatchedPages[link]; !ok {
						d.dispatchedPages[link] = struct{}{}
						d.dispatch(link, page.Depth-1)
					}
				}
			}
		}
	}()

	return d
}

func (d Dispatcher) printStats() {
	fmt.Printf("Processed: %d Dispatched: %d\n", len(d.crawledPages), len(d.dispatchedPages))
}

func (d Dispatcher) dispatch(url string, depth int) {
	fmt.Printf("Dispatch[%d]: %s\n", depth, url)
	pp := PageProcessor{
		Source: url,
		Output: d.output,
		depth:  depth,
	}
	d.pool.Do(pp)
}

func (d Dispatcher) Start(url string) {
	d.dispatch(url, d.depth)
}
