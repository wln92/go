package main

import (
	"fmt"
	"time"
)

func main()  {
	stopch := make(chan struct{})

	go func() {
		for {
			fmt.Printf("Hello World!\n")
			time.Sleep(1 * time.Second)
		}
	}()

	<- stopch
}
