package thread

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutex = Only one goroutine can access the critical section of code at a time.
If one goroutine calls mu.Lock(), other goroutines that call mu.Lock() will block until
the first goroutine calls mu.Unlock().

Condition
*/
func DemoSignals() {
	value := 0
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu) // Conditions wait needs to release lock to allow for other goroutines to run, so it needs a mutex to manage access to the shared resource
	ready := false

	go uploadSignal(&value, mu, cond, &ready) // &value - takes the address

	mu.Lock() // If grabs first, the cond.Wait will release lock and wait for signal
	for !ready {
		cond.Wait() // Wait for the signal that the value is ready, once ready receives lock again to continue locked code
	}

	fmt.Println(value)
	mu.Unlock()
}
func uploadSignal(value *int, mu *sync.Mutex, cond *sync.Cond, ready *bool) {
	mu.Lock()
	// simulate upload
	time.Sleep(6000 * time.Second)
	*value = 10 // *val - dereference the pointer (accesses/sets the underlying value)
	*ready = true
	cond.Signal() // Signal the waiting goroutine that the value is ready
	mu.Unlock()
}
