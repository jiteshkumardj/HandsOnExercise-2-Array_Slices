package main

import "fmt"

func main() {
	// Createing Array
	ai := [5]int{}

	//assign value to each index position
	ai[0] = 1
	ai[1] = 2
	ai[2] = 3
	ai[3] = 4
	ai[4] = 5

	// printing out
	for _, v := range ai {
		fmt.Printf("%v and the type is %T", v, v)
	}

}
