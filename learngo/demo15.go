package main

import "fmt"

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:8]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)

	//fmt.Printf("s3:%d, s4:%d\n", &s3, &s4)
}
