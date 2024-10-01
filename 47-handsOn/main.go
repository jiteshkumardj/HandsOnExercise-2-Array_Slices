package main

import (
	"fmt"
)

func main() {
	x1 := make([]int, 50)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	x2 := make([]int, 0, 50)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	x1 = append(x1, 98)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	x2 = append(x2, 99)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))
}
