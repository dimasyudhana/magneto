package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func do(arg map[string]any, msg chan map[string]any) {
	msg <- arg
}

func main() {
	// Use the number of available CPU cores for concurrency
	maxWorkers := runtime.GOMAXPROCS(0)
	sem := semaphore.NewWeighted(int64(maxWorkers))
	ctx := context.Background()

	start := time.Now()

	args := map[string]any{
		"field1":  100.0,
		"field2":  "99.0",
		"field3":  true,
		"field4":  100.0,
		"field5":  "99.0",
		"field6":  true,
		"field7":  100.0,
		"field8":  "99.0",
		"field9":  true,
		"field10": 100.0,
	}

	msg := make(chan map[string]any, len(args))
	var wg sync.WaitGroup

	for key, val := range args {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		wg.Add(1)
		go func(key string, val any) {
			defer sem.Release(1)
			defer wg.Done()

			arg := map[string]any{key: val}
			do(arg, msg)
		}(key, val)
	}

	wg.Wait()
	close(msg)

	for data := range msg {
		for key, val := range data {
			log.Printf("%s : %v", key, val)
		}
	}

	log.Printf("Concurrent semaphore execution time: %v", time.Since(start))
}
