package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string { // <-chan string means receives-only channel of string.
	// return channel
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return c
}

func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs {
		go func(cv <-chan string) {
			for {
				c <- <-cv
			}
		}(ci)
	}
	return c
}

func main() {
	c := fanInSimple(boring("Joe"), boring("Ahn"), boring("Wontak"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring. I'm leaving.")
}
