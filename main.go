package main

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/mylakehead/sync/events"
	"github.com/mylakehead/sync/lib"
	"github.com/mylakehead/sync/runtime"
)

func main() {
	rt, err := runtime.New()
	if err != nil {
		log.Fatal(err)
	}

	abs, err := filepath.Abs(rt.Config.Contract.IndexPath)
	if err != nil {
		log.Fatal(err)
	}

	err = lib.CreatePath(abs)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	marketEvent, err := events.NewMarketEvent(&wg, rt, abs)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go marketEvent.ListenEvents()
	wg.Wait()
}
