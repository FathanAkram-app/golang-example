package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrencyTest() {
	var wg sync.WaitGroup
	wg.Add(1)

	// "go" = goroutine
	// go anyFunction() <-- format
	// it runs in the background, so the next line of code can be excecuted at the same time
	go func() {
		count("sheep", 4)
		wg.Done()
	}()
	count("fish", 3)

	wg.Wait()

}

func count(thing string, x int) {
	for i := 1; i <= x; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}
