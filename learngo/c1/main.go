package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	test3()
}

func test3() {
	fmt.Println(strings.Join(os.Args[1:], ","))
}

func test1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func test2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}