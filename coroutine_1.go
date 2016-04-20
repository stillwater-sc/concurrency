package main
import (
	"fmt"
	"math/rand"
	"time"
)

func rnd_time_loop(msg string) <-chan string {
	c := make(chan string)
	go func() { // launch an anonymous goroutine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c  // return the channel to the caller
}
func main() {
	c := rnd_time_loop("timed ")
	for i := 0; i < 5; i++ {
		fmt.Printf("Msg: %q\n", <-c)
	}
	fmt.Println("Program is done")
}
