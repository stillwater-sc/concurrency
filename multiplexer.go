package main

import (
    "fmt"
    "math/rand"
    "time"
)
func fanin1(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

func fanin2(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <-input2: c <- s
			}
		}
	}()
	return c
}

func fanin3(input1, input2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <-input2: c <- s
			}
		}
	}()
	return c
}

func process_A() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Process A produced value #%d", i)
            time.Sleep(time.Duration(rand.Intn(5e2)) * time.Millisecond)
		}
	}()
	return c
}
func process_B() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Process B produced value #%d", i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func process_Aprime() chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Process A produced value #%d", i)
			time.Sleep(time.Duration(rand.Intn(5e2)) * time.Millisecond)
		}
	}()
	return c
}
func process_Bprime() chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Process B produced value #%d", i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// all fanin[123] are functionally equivalent
// however, the type 'receiver channel' === <-chan string is more accurate
// with respect to the use of the channels returned by fanin and Process_[AB]
// You are not using the channel to push values into, only pull data out
// If you try to push on a <-chan you'll get the following error:
// c <- "hello"
// invalid operation: c <- 1 (send to a receive-only type <-chan string
// Otherwise stated, fanin2 with Process_A and Process_B are preferred
// over fanin3 with Process_Aprime and Process_Bprime
func main() {
//	c := fanin1(process_A(), process_B())
//	c := fanin2(process_A(), process_B())
	c := fanin3(process_Aprime(), process_Bprime())
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Program is done")
}
