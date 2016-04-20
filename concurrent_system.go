package main
import (
	"time"
	"fmt"
    "math/rand"
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

var (
	Web    = fakeSearch("web")
	Web1   = fakeSearch("web-1")
	Web2   = fakeSearch("web-2")
	Image  = fakeSearch("image")
   	Image1 = fakeSearch("image-1")
   	Image2 = fakeSearch("image-2")
	Image3 = fakeSearch("image-3")
	Video  = fakeSearch("video")
	Video1 = fakeSearch("video-1")
	Video2 = fakeSearch("video-2")
	Video3 = fakeSearch("video-3")
)

func SequentialSearch(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func ConcurrentSearch(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func ConcurrentSearchWithCutOff(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {  // for each goroutine that is ready pick up results
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func ReplicaSearch(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2, Image3) }()
	go func() { c <- First(query, Video1, Video2, Video3) }()

    	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("seach timed out")
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

//	results := SequentialSearch("golang")
//	results := ConcurrentSearch("golang")
//	results := ConcurrentSearchWithCutOff("golang")
//	results := First("golang", fakeSearch("replica-1"), fakeSearch("replica-2"))
	results := ReplicaSearch("golang")

	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
