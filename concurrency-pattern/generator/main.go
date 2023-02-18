package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {
	// return channel
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(100 * time.Millisecond)
		}
		// Close channel after goroutine is finished.
		close(c)
	}()
	return c
}

func main() {
	joe := boring("Joe")
	ahn := boring("Ahn")

	for i := 0; i < 20; i++ {
		fmt.Println(i, "th")
		fmt.Println(<-joe)
		fmt.Println(<-ahn)
	}

	fmt.Println("You are both boring. I'm leaving.")
}
