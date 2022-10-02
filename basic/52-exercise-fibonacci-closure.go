package main

import "fmt"

func fibonacci() func(int) int {
	t1 := 0
	t2 := 0
	return func(n int) int {
		res := t1 + t2
		t1 = t2
		t2 = n
		return res

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i ++ {
		fmt.Println(i, f(i))
	}
}
