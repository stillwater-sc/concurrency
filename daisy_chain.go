package main

import (
    "fmt"
)
func stage(left, right chan int) {
	right <- (1 + <-left)
}
// building a pipeline pushing 0 in and getting the number of stages out
func main() {
	const n = 100000
	start := make(chan int)
	left  := start
	right := left
	for i := 0; i < n; i++ {
		right = make(chan int)
		go stage(left, right)
		left = right
	}
    // make the last one a sink channel
	sink := left
	// inject the starting value into the daisy chain
	start <- 0
	fmt.Println(<-sink)
}
