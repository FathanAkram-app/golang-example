package main

import (
	"fmt"
	"time"
)

func concurrencyTest() {
	fmt.Println("")
	// uncomment any function below to try a function

	// goroutineUnderstanding()

	// channelsUnderstanding()

	// selectChannelUnderstanding()

}

func goroutineUnderstanding() {
	c := make(chan string)

	// "go" = goroutine
	// go anyFunction() <-- format
	// it runs in the background, so the next line of code can be excecuted at the same time
	go count("sheep", 4, c)
	fmt.Println("it started already, sorry im late to tell you")
	// recieve a message from channel
	for msg := range c {
		fmt.Println(msg)
	}
	fmt.Println("")
}
func count(thing string, x int, c chan string) {
	for i := 1; i <= x; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
		// wait for recieve "thing"
		c <- thing
	}
	close(c)
}

func channelsUnderstanding() {
	// make a new instance of channel
	c := make(chan string, 2)

	// send something to channel c
	c <- "hey"
	c <- "fathan"

	// recieve from channel c
	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

}

func selectChannelUnderstanding() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 Seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// select is basically a switch case statement
		// if some channel is ready, the code will be executed
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

}
