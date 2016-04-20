package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

func process2 (msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s with sequence #%d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := process2("continuous event stream")
	interrupt := time.After(time.Duration( 2 * time.Second))
	timeout := time.After(time.Duration(5 * time.Second))
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case  <-interrupt:
			fmt.Println("Process is interrupted")
			os.Exit(2)
		case  <-timeout:
			fmt.Println("Process has timed out")
			os.Exit(1)
		}
	}
}
