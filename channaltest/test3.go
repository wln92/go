package main

import (
	"fmt"
	"time"
)

func testchan1() {
	ch := make(chan int)
	ch = nil
	go func() {
		//<- ch
		ch <- 1
		fmt.Printf("done!\n")
	}()
	time.Sleep(3 * time.Second)
	fmt.Printf("channel done!\n")
}

func main() {
	testchan1()
}
