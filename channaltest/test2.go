package main

import (
	"fmt"
	"time"
)

func send(ch chan string)  {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("%v", i)
	}
}

func receive(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
	fmt.Printf("received done!\n")
	close(ch)
}

func main()  {
	ch := make(chan string, 10)
	go send(ch)
	go receive(ch)
	time.Sleep(1 * time.Second)
	fmt.Printf("main done")
}