package main

import (
	"os"
	"os/signal"
)

func main() {
	pool := NewWorkerPool(10)
	output := make(chan Page, 1024)

	dispatcher := NewDispatcher(pool, output, 2)
	dispatcher.Start("http://bbq.kylerames.com")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
	pool.Shutdown()
}
