package thread

import (
	"fmt"
	"time"
)

func DemoChannel() {
	ch := make(chan int)
	// Run the goroutine and pass the channel
	go uploadChannel(ch)

	// Blocks the main go routine until it receives a value from the channel (sleep)
	val := <-ch
	fmt.Println(val)
}

func uploadChannel(desiredChannel chan int) {
	// simulate upload
	time.Sleep(2 * time.Second)
	desiredChannel <- 10
}
