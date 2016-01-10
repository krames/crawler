package main

import (
	"sync"
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
