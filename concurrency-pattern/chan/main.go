package main

import (
	"fmt"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// init out channel: communicate main <->boring goroutine.
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 10; i++ {
		// <-c read the value from boring function
		// <-c wait for a value to be sent
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring. I'am leaving")
}
