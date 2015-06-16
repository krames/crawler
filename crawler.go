package main

import (
	"os"
	"os/signal"

	"github.com/krames/crawler/domain"
)

func main() {
	pool := NewWorkerPool(4)
	output := make(chan domain.Page)

	dispatcher := NewDispatcher(pool, output)
	dispatcher.Start("http://bbq.kylerames.com")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
	pool.Shutdown()
}
