package main

import "fmt"

var block = "package"
//
func main()  {
	var block = 123
	{
		var block = "inner"
		fmt.Printf("The inner block is %s.\n", block)
	}
	//block = "test"
	fmt.Printf("The outer block is %v.\n", block)
}
