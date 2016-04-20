package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func continuous_process(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0;; i++ {
			c <- fmt.Sprintf("%s #%d at %v", msg, i, time.Now())
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := continuous_process("continuous event")
	timeout := time.After(time.Duration(5 * time.Second))
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Time out reached")
			os.Exit(0)
		}
	}
}
