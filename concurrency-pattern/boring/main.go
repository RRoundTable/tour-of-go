package main

import (
	"fmt"
	"time"
)

func boring(msg string) {
	// nerver stop goroutine
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// main goroutine calls boring gorutine.
	// But main gorutine doesn't wait for boring finished.
	go boring("Boring")

	fmt.Println("I'am listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring. I'm leaving")
	// When main goroutine is finished, boring goroutine exisited even though boring goroutine is not finished.
}
