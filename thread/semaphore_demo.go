package thread

import (
	"fmt"
	"sync"
	"time"
)

type semaphore struct {
	available int
	mu        *sync.Mutex
	cond      *sync.Cond
}

func DemoSemaphore() {
	mu := &sync.Mutex{}
	sem := &semaphore{available: 3, mu: mu, cond: sync.NewCond(mu)}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			acquire(sem)
			// Simulate work with the acquired resource
			time.Sleep(4 * time.Second)
			fmt.Println("Goroutine", id, "finished processing")
			release(sem)
		}(i)
	}
	wg.Wait()
}

func acquire(sem *semaphore) {
	sem.mu.Lock()
	for sem.available <= 0 {
		sem.cond.Wait() // Wait for the signal that a resource is available, once signaled receives lock again to continue locked code
	}
	sem.available--
	sem.mu.Unlock()
}

func release(sem *semaphore) {
	sem.mu.Lock()
	sem.available++
	sem.cond.Signal() // Signal one waiting goroutine that a resource is available (if there are none waiting, the signal is lost)
	sem.mu.Unlock()
}
