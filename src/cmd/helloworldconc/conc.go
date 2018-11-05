package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i, ch)
	}

	for {
		fmt.Println(<-ch)
	}

	//time.Sleep(10 * time.Second)

}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello World from goroutine %d!\n", i)
	}
}
