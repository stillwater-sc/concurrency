package main
import (
	"fmt"
	"math/rand"
	"time"
)

func timed_process(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0;; i++ {
			delay := rand.Intn(1.1e3)
			c <- fmt.Sprintf("%s produced value #%d, delaying %dmsec", msg, i, delay)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := timed_process("Message Generator")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Generator was too slow")
			return
		}
	}
}
