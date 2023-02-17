package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		fmt.Println(<-ch)
	}(ch)
	ch <- 10
}
