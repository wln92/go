package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
}

func write(ch chan string) {
	ch <- "Hello World"
	fmt.Printf("write done!\n")
}

func read(ch chan string) {
	value, ok := <- ch
	fmt.Printf("received value:%v, ok:%v\n", value, ok)
}

func close(ch chan string) {
	close(ch)
	fmt.Printf("close channel done\n")
}

func test1() {
	ch := make(chan string)
	go write(ch)
	time.Sleep(1 * time.Second)
	go close(ch)
	time.Sleep(1 * time.Second)
	go read(ch)
	time.Sleep(2 * time.Second)
}
