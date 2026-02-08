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

Condition = A wrapper around a mutex that allows threads to sleep/signal other go routines
*/
func DemoMutex() {
	value := 0
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu) // Conditions wait needs to release lock to allow for other goroutines to run, so it needs a mutex to manage access to the shared resource
	ready := false

	go uploadSignal(&value, &mu, cond, &ready) // &value - takes the address

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
	time.Sleep(2 * time.Second)
	*value = 10 // *val - dereference the pointer (accesses/sets the underlying value)
	*ready = true
	cond.Signal() // Signal one waiting goroutine that the value is ready (if there are none waiting, the signal is lost)
	mu.Unlock()
}
